package reverse_proxy

import (
	"fmt"
	"github.com/bots-garden/capsule/capsulelauncher/commons"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"math/rand"
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
		c.JSON(http.StatusInternalServerError, gin.H{"code": "ERROR", "message": "😢 Houston? We have a problem 🥵"})
	}

}

func proxyRevision(c *gin.Context) {

	functionName := c.Param("function_name")
	functionRevision := c.Param("function_revision")
	functionUrls := functions[functionName][functionRevision]

	if functionUrls != nil {
		redirect(functionUrls.([]interface{}), c)
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

	// Get the list of the registered functions
	router.GET("/memory/functions/list", func(c *gin.Context) {
		// transform functions to json (I'm pretty sure that there is something simpler)

		registeredFunctions := make(map[string]map[string]interface{})

		for functionName, revisions := range functions {

			functionsRevisions := make(map[string]interface{})
			for revisionName, element := range revisions {
				functionsRevisions[revisionName.(string)] = element
			}
			registeredFunctions[functionName.(string)] = functionsRevisions
		}
		c.JSON(http.StatusOK, registeredFunctions)
	})

	/*
	   ==============================================================
	   Register a function with a unique revision (and only one url)
	   ==============================================================
	      Payload
	         {
	           "function": "morgen", "revision": "default", "url": "http://localhost:5050"
	         }
	      Curl Query
	       curl -v -X POST \
	         http://localhost:8888/memory/functions/registration \
	         -H 'content-type: application/json; charset=utf-8' \
	         -d '{"function": "morgen", "revision": "default", "url": "http://localhost:5050"}'
	       echo ""

	*/
	router.POST("memory/functions/registration", func(c *gin.Context) {
		//TODO: add an authentication token
		jsonMap := make(map[string]interface{})
		if err := c.Bind(&jsonMap); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    "ERROR",
				"message": err.Error()})
		}

		//TODO: check if the values are empty or not
		functionName := jsonMap["function"].(string)
		revisionName := jsonMap["revision"].(string)
		firstUrl := jsonMap["url"].(string)

		if functions[functionName] != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    "ERROR_FUNCTION_ALREADY_EXISTS",
				"message": functionName + " already exists"})
		} else {
			functions[functionName] = map[interface{}]interface{}{revisionName: []string{firstUrl}}

			//add a revision to a function
			//functions["sandbox"]["green"] = []string{"http://localhost:5052"}

			c.JSON(http.StatusAccepted, gin.H{
				"code":     "OK",
				"message":  "Function created",
				"function": functionName,
				"revision": revisionName,
				"url":      firstUrl})
		}

	})
	//TODO: remove a function
	router.DELETE("memory/functions/registration", func(c *gin.Context) {

	})

	/*
	   ==========================================================
	   Add a revision (and only one url) to an existing function
	   ==========================================================
	      Payload
	         {
	           "revision": "default", "url": "http://localhost:5050"
	         }
	      Curl Query
	       curl -v -X POST \
	         http://localhost:8888/memory/functions/morgen/revision \
	         -H 'content-type: application/json; charset=utf-8' \
	         -d '{"revision": "blue", "url": "http://localhost:5051"}'
	       echo ""
	*/
	router.POST("memory/functions/:function_name/revision", func(c *gin.Context) {
		functionName := c.Param("function_name")

		//TODO: add an authentication token
		jsonMap := make(map[string]interface{})
		if err := c.Bind(&jsonMap); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    "ERROR_JSON_BINDING",
				"message": err.Error()})
		}

		//TODO: check if the values are empty or not
		revisionName := jsonMap["revision"].(string)
		firstUrl := jsonMap["url"].(string)

		// if the function does not exist, we cannot add a revision
		if functions[functionName] == nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    "ERROR_FUNCTION_NOT_FOUND",
				"message": functionName + " does not exist"})
		} else {
			// if the revision already exists we cannot add it
			if functions[functionName][revisionName] != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    "ERROR_REVISION_ALREADY_EXISTS",
					"message": revisionName + " already exists"})
			} else {
				// add the revision to the function
				functions[functionName][revisionName] = []string{firstUrl}

				c.JSON(http.StatusAccepted, gin.H{
					"code":     "OK",
					"message":  "Revision added to the function",
					"function": functionName,
					"revision": revisionName,
					"url":      firstUrl})
			}
		}
	})
	//TODO: remove a revision from a function
	router.DELETE("memory/functions/:function_name/revision", func(c *gin.Context) {

	})

	/*
	   ======================================
	   Add a URL to an existing revision
	   ======================================
	      Payload
	         {
	           "url": "http://localhost:5053"
	         }
	      Curl Query
	       curl -v -X POST \
	         http://localhost:8888/memory/functions/morgen/blue/url \
	         -H 'content-type: application/json; charset=utf-8' \
	         -d '{"url": "http://localhost:5053"}'
	       echo ""
	       Remark: it's like a scale
	*/
	router.POST("memory/functions/:function_name/:function_revision/url", func(c *gin.Context) {
		functionName := c.Param("function_name")
		revisionName := c.Param("function_revision")

		//TODO: add an authentication token
		jsonMap := make(map[string]interface{})
		if err := c.Bind(&jsonMap); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    "ERROR_JSON_BINDING",
				"message": err.Error()})
		}

		//TODO: check if the values are empty or not
		urlToAdd := jsonMap["url"].(string)

		// if the function does not exist, we cannot add an url to a revision
		if functions[functionName] == nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    "ERROR_FUNCTION_NOT_FOUND",
				"message": functionName + " does not exist"})
		} else {
			// if the revision does not exist we cannot add an url to it
			if functions[functionName][revisionName] == nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    "ERROR_REVISION_NOT_FOUND",
					"message": revisionName + " does not exist"})
			} else {
				// add the url to the revision
				currentUrlsList := functions[functionName][revisionName].([]string)
				currentUrlsList = append(currentUrlsList, urlToAdd)
				functions[functionName][revisionName] = currentUrlsList

				c.JSON(http.StatusAccepted, gin.H{
					"code":     "OK",
					"message":  "URL added to the revision",
					"function": functionName,
					"revision": revisionName,
					"url":      urlToAdd})
			}
		}

	})

	/*
	   =======================================
	   Remove a URL from an existing revision
	   =======================================
	      Payload
	         {
	           "url": "http://localhost:5053"
	         }
	      Curl Query
	       curl -v -X DELETE \
	         http://localhost:8888/memory/functions/morgen/blue/url \
	         -H 'content-type: application/json; charset=utf-8' \
	         -d '{"url": "http://localhost:5053"}'
	       echo ""
	       Remark: it's like a scale
	*/
	router.DELETE("memory/functions/:function_name/:function_revision/url", func(c *gin.Context) {

		functionName := c.Param("function_name")
		revisionName := c.Param("function_revision")

		//TODO: add an authentication token
		jsonMap := make(map[string]interface{})
		if err := c.Bind(&jsonMap); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    "ERROR_JSON_BINDING",
				"message": err.Error()})
		}

		//TODO: check if the values are empty or not
		urlToRemove := jsonMap["url"].(string)

		// if the function does not exist, we cannot remove an url from a revision
		if functions[functionName] == nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    "ERROR_FUNCTION_NOT_FOUND",
				"message": functionName + " does not exist"})
		} else {
			// if the revision does not exist we cannot remove an url from it
			if functions[functionName][revisionName] == nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    "ERROR_REVISION_NOT_FOUND",
					"message": revisionName + " does not exist"})
			} else {
				// Remove the url from the revision
				currentUrlsList := functions[functionName][revisionName].([]string)

				for index, url := range currentUrlsList {
					if url == urlToRemove {
						currentUrlsList = append(currentUrlsList[:index], currentUrlsList[index+1:]...)
						break
					}
				}

				functions[functionName][revisionName] = currentUrlsList

				c.JSON(http.StatusAccepted, gin.H{
					"code":     "OK",
					"message":  "URL removed from the revision",
					"function": functionName,
					"revision": revisionName,
					"url":      urlToRemove})
			}
		}

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
