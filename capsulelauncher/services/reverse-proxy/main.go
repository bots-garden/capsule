package reverse_proxy

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// 👀 See https://github.com/bots-garden/procyon/blob/main/procyon-reverse-proxy/main.go
func proxy(c *gin.Context) {
	var functionUrl string
	functionName := c.Param("function_name")

	//functionUrl := procyonDomain + ":" + strconv.Itoa(defaultRevisionsMap[c.Param("function_name")].WasmFunctionHttpPort)

	if functionName == "hello" {
		functionUrl = "http://localhost:9091"
	}

	if functionName == "hey" {
		functionUrl = "http://localhost:9092"
	}

	if functionName == "hola" {
		functionUrl = "http://localhost:9093"
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

func proxyRevision(c *gin.Context) {
	// Foo
}

func Serve(httpPort string) {

	//go getFunctionsList()
	//go getDefaultRevisionsList()

	r := gin.Default()

	//Create catchall routes
	r.Any("/functions/:function_name", proxy)
	// 🚧 work in progress
	r.Any("/functions/:function_name/:function_revision", proxyRevision)

	// TODO: handle it with a flag
	if getEnv("PROXY_CRT", "") != "" {
		r.RunTLS(":"+getEnv("PROXY_HTTPS", "4443"), getEnv("PROXY_CRT", "certs/procyon-registry.local.crt"), getEnv("PROXY_KEY", "certs/procyon-registry.local.key"))
	} else {
		r.Run(":" + httpPort)
	}

}
