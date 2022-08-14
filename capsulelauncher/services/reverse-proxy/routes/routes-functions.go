package reverse_proxy_memory_routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefineFunctionsRoutes(router *gin.Engine, functions map[interface{}]map[interface{}]interface{}) {

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
		c.IndentedJSON(http.StatusOK, registeredFunctions)
		//c.JSON(http.StatusOK, registeredFunctions)
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

	/*
	   ==============================================================
	   Remove a function (delete the registration)
	   ==============================================================
	      Payload
	         {
	           "function": "morgen"
	         }
	      Curl Query
	       curl -v -X POST \
	         http://localhost:8888/memory/functions/registration \
	         -H 'content-type: application/json; charset=utf-8' \
	         -d '{"function": "morgen"}'
	       echo ""

	*/
	router.DELETE("memory/functions/registration", func(c *gin.Context) {
		//TODO: add an authentication token
		jsonMap := make(map[string]interface{})
		if err := c.Bind(&jsonMap); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    "ERROR",
				"message": err.Error()})
		}

		//TODO: check if the values are empty or not
		functionName := jsonMap["function"].(string)

		if functions[functionName] == nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    "ERROR_FUNCTION_NOT_FOUND",
				"message": functionName + " does not exist"})
		} else {

			delete(functions, functionName)

			c.JSON(http.StatusAccepted, gin.H{
				"code":     "OK",
				"message":  "Function removed",
				"function": functionName})
		}
	})
}
