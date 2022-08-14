package reverse_proxy

import (
	"fmt"
	"github.com/bots-garden/capsule/capsulelauncher/commons"
	reverse_proxy_memory_routes "github.com/bots-garden/capsule/capsulelauncher/services/reverse-proxy/routes"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

/*
func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}
*/

var lastUrlIndex = 0

func redirect(functionUrls []string, c *gin.Context) {
	//fmt.Println("🟢🖐", functionUrls)
	var functionUrl = ""

	if len(functionUrls) == 1 {
		functionUrl = functionUrls[0]
	} else {
		//TODO better repartition (load balancing) handling
		lastUrlIndex += 1
		if lastUrlIndex == len(functionUrls) {
			lastUrlIndex = 0
		}

		functionUrl = functionUrls[lastUrlIndex]

		//fmt.Println("🛑 kind of load balancing", lastUrlIndex)
		//fmt.Println("🌍", functionUrl)
	}

	remote, err := url.Parse(functionUrl)

	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")

		//fmt.Println("🔴", c.Request.Header)
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

// 👀 See https://github.com/bots-garden/procyon/blob/main/procyon-reverse-proxy/main.go
func proxy(c *gin.Context) {

	functionName := c.Param("function_name")
	functionUrls := functions[functionName]["default"]

	//fmt.Println("👋functionName", functionName)
	//fmt.Println("👋functionUrls", functionUrls)

	if functionUrls != nil {
		/*
		   2022/08/09 14:11:44 http: panic serving 127.0.0.1:65399:
		   interface conversion: interface {} is []string, not []interface {}

		*/

		//redirect(functionUrls.([]interface{}), c)
		redirect(functionUrls.([]string), c)

	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "ERROR", "message": "😢 Houston? We have a problem 🥵"})
	}

}

func proxyRevision(c *gin.Context) {

	functionName := c.Param("function_name")
	functionRevision := c.Param("function_revision")
	functionUrls := functions[functionName][functionRevision]

	if functionUrls != nil {
		redirect(functionUrls.([]string), c)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "ERROR", "message": "😢 Houston? We have a problem 🥵"})
	}
}

var yamlConfig = make(map[interface{}]map[interface{}]map[interface{}]interface{})
var functions = make(map[interface{}]map[interface{}]interface{})
var filters = make(map[interface{}]map[interface{}]interface{})

func Serve(httpPort, config, backend, crt, key string) {

	if config != "" {
		fmt.Println("📝 config file:", config)
		yamlFile, errFile := ioutil.ReadFile(config)
		if errFile != nil {
			log.Fatal(errFile)
		}

		errYaml := yaml.Unmarshal(yamlFile, &yamlConfig)

		if errYaml != nil {
			log.Fatal(errYaml)
		}

		functions = yamlConfig["functions"]
		filters = yamlConfig["filters"]

		// The code below is for testing
		/*
		   functions["sandbox"] = map[interface{}]interface{}{"default": []string{"http://localhost:5050"}, "blue": []string{"http://localhost:5051"}}

		   functions["sandbox"]["green"] = []string{"http://localhost:5052"}

		   urlsList := functions["sandbox"]["green"].([]string)
		   urlsList = append(urlsList, "http://localhost:5053")

		   functions["sandbox"]["green"] = urlsList
		*/
	}

	switch backend {
	case "yaml":
		fmt.Println("👋 routes are defined in", config)
	case "memory":
		// it's possible to mix memory and yaml for the functions
		fmt.Println("👋 routes are defined in memory")
	case "redis":
	default:
		fmt.Println("👋 routes are defined in a Redis backend")
		//TODO check the environment variables
		fmt.Println("👋 use environment variable for the Redis configuration ")
	}

	if commons.GetEnv("DEBUG", "false") == "false" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	//router := gin.Default()
	router := gin.New()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "😢 Page not found 🥵"})
	})

	/*
	   You need to use a header with this key: CAPSULE_REVERSE_PROXY_ADMIN_TOKEN
	*/
	reverseProxyAdminToken := commons.GetEnv("CAPSULE_REVERSE_PROXY_ADMIN_TOKEN", "")

	reverse_proxy_memory_routes.DefineFunctionsRoutes(router, functions, reverseProxyAdminToken)
	reverse_proxy_memory_routes.DefineRevisionsRoutes(router, functions, reverseProxyAdminToken)
	reverse_proxy_memory_routes.DefineUrlsRoutes(router, functions, reverseProxyAdminToken)

	// Call the functions
	router.Any("/functions/:function_name", proxy)
	router.Any("/functions/:function_name/:function_revision", proxyRevision)

	if crt != "" {
		// certs/procyon-registry.local.crt
		// certs/procyon-registry.local.key
		fmt.Println("💊 Capsule (", commons.CapsuleVersion(), ") Reverse-Proxy is listening on:", httpPort, "🔐🌍")

		router.RunTLS(":"+httpPort, crt, key)
	} else {
		fmt.Println("💊 Capsule (", commons.CapsuleVersion(), ") Reverse-Proxy is listening on:", httpPort, "🌍")
		router.Run(":" + httpPort)
	}

}
