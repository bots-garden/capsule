package reverse_proxy

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	_ "gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func redirect(functionUrls []interface{}, c *gin.Context) {
	var functionUrl = ""
	if len(functionUrls) == 1 {
		functionUrl = functionUrls[0].(string)
	} else {
		//TODO better repartition handling
		min := 0
		max := len(functionUrls) - 1
		functionUrl = functionUrls[rand.Intn(max-min)+min].(string)
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
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

// 👀 See https://github.com/bots-garden/procyon/blob/main/procyon-reverse-proxy/main.go
func proxy(c *gin.Context) {

	functionName := c.Param("function_name")
	functionUrls := functions[functionName]["default"]

	if functionUrls != nil {
		redirect(functionUrls.([]interface{}), c)
	} else {
		c.Next()
		//c.JSON(http.StatusBadGateway, gin.H{"code": "FUNCTION_NOT_FOUND", "message": "😢 Houston? We have a problem 🥵"})
	}

	redirect(functionUrls.([]interface{}), c)
}

func proxyRevision(c *gin.Context) {

	functionName := c.Param("function_name")
	functionRevision := c.Param("function_revision")
	functionUrls := functions[functionName][functionRevision]
	if functionUrls != nil {
		redirect(functionUrls.([]interface{}), c)
	} else {
		c.Next()
		//c.JSON(http.StatusBadGateway, gin.H{"code": "FUNCTION_NOT_FOUND", "message": "😢 Houston? We have a problem 🥵"})
	}

}

var functions = make(map[interface{}]map[interface{}]interface{})

func getKindOfConfig(config string) string {
	if filepath.Ext(config) == ".yaml" {
		return "yaml"
	}
	return ""
}

func Serve(httpPort, config, crt, key string) {

	switch what := getKindOfConfig(config); what {
	case "yaml":
		fmt.Println("📝 routes are defined in:", config)
		yamlFile, errFile := ioutil.ReadFile(config)
		if errFile != nil {
			log.Fatal(errFile)
		}

		errYaml := yaml.Unmarshal(yamlFile, &functions)

		if errYaml != nil {
			log.Fatal(errYaml)
		}

	default:
		fmt.Println("👋 routes are defined in memory")
	}

	ErrorHandler := func(c *gin.Context) {
		c.Next()
		c.JSON(http.StatusInternalServerError, gin.H{"code": "ERROR", "message": "😢 Houston? We have a problem 🥵"})
	}

	if getEnv("DEBUG", "false") == "false" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	//router := gin.Default()
	router := gin.New()

	router.Use(ErrorHandler)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "😢 Page not found 🥵"})
	})

	router.Any("/functions/:function_name", proxy)
	router.Any("/functions/:function_name/:function_revision", proxyRevision)

	if crt != "" {
		// certs/procyon-registry.local.crt
		// certs/procyon-registry.local.key
		fmt.Println("💊 Capsule Reverse-Proxy is listening on:", httpPort, "🔐🌍")

		router.RunTLS(":"+httpPort, crt, key)
	} else {
		fmt.Println("💊 Capsule Reverse-Proxy is listening on:", httpPort, "🌍")
		router.Run(":" + httpPort)
	}

}
