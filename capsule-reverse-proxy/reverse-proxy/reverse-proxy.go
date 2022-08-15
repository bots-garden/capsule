package reverse_proxy

import (
    "fmt"
    "github.com/bots-garden/capsule/capsule-reverse-proxy/reverse-proxy/routes"
    "github.com/bots-garden/capsule/commons"
    "github.com/gin-gonic/gin"
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

func proxy(c *gin.Context) {

    functionName := c.Param("function_name")
    functionUrls := functions[functionName]["default"]

    if functionUrls != nil {
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

//var yamlConfig = make(map[interface{}]map[interface{}]map[interface{}]interface{})
var functions = make(map[interface{}]map[interface{}]interface{})

//var filters = make(map[interface{}]map[interface{}]interface{})

func Serve(httpPort, config, backend, crt, key string) {

    switch backend {
    case "memory":
        // it's possible to mix memory and yaml for the functions
        fmt.Println("👋 routes are defined in memory")
    case "redis":
        fmt.Println("👋 Redis backend (🚧 not implemented)")
        fmt.Println("👋 routes are defined in a Redis backend")
        fmt.Println("👋 use environment variable for the Redis configuration ")
    default:
        fmt.Println("👋 routes are defined in a Redis backend")
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
