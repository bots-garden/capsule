package reverse_proxy_memory_routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefineRevisionsRoutes(router *gin.Engine, functions map[interface{}]map[interface{}]interface{}, reverseProxyAdminToken string) {

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
		//TODO: check if there is a better practice to handle authentication token
		if len(reverseProxyAdminToken) == 0 || CheckReverseProxyAdminToken(c, reverseProxyAdminToken) {

			functionName := c.Param("function_name")

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
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    "KO",
				"from":    "reverse-proxy",
				"message": "Forbidden"})
		}

	})

	/*
	   ==========================================================
	   Remove a revision from an existing function
	   ==========================================================
	      Payload
	         {
	           "revision": "blue"
	         }
	      Curl Query
	       curl -v -X POST \
	         http://localhost:8888/memory/functions/morgen/revision \
	         -H 'content-type: application/json; charset=utf-8' \
	         -d '{"revision": "blue"}'
	       echo ""
	*/
	router.DELETE("memory/functions/:function_name/revision", func(c *gin.Context) {
		//TODO: check if there is a better practice to handle authentication token
		if len(reverseProxyAdminToken) == 0 || CheckReverseProxyAdminToken(c, reverseProxyAdminToken) {

			functionName := c.Param("function_name")

			jsonMap := make(map[string]interface{})
			if err := c.Bind(&jsonMap); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    "ERROR_JSON_BINDING",
					"message": err.Error()})
			}

			//TODO: check if the values are empty or not
			revisionName := jsonMap["revision"].(string)
			//firstUrl := jsonMap["url"].(string)

			// if the function does not exist, we cannot add a revision
			if functions[functionName] == nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    "ERROR_FUNCTION_NOT_FOUND",
					"message": functionName + " does not exist"})
			} else {
				// if the revision does not exist we cannot remove it
				if functions[functionName][revisionName] == nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code":    "ERROR_REVISION_NOT_FOUND",
						"message": revisionName + " does not exist"})
				} else {
					// add the revision to the function
					//functions[functionName][revisionName] = []string{firstUrl}
					delete(functions[functionName], revisionName)

					c.JSON(http.StatusAccepted, gin.H{
						"code":     "OK",
						"message":  "Revision removed from the function",
						"function": functionName,
						"revision": revisionName})
				}
			}
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    "KO",
				"from":    "reverse-proxy",
				"message": "Forbidden"})
		}

	})
}
