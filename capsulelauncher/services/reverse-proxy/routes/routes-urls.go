package reverse_proxy_memory_routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefineUrlsRoutes(router *gin.Engine, functions map[interface{}]map[interface{}]interface{}) {

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

		// TODO: add an authentication token
		jsonMap := make(map[string]interface{})
		if err := c.Bind(&jsonMap); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    "ERROR_JSON_BINDING",
				"message": err.Error()})
		}

		// TODO: check if the values are empty or not
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
				// TODO: check if the url already exists
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
	       Remark: it's like a downscale
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
}
