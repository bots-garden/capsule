package routes

import (
	"encoding/json"
	"github.com/bots-garden/capsule/capsulelauncher/services/worker/helpers"
	"github.com/bots-garden/capsule/capsulelauncher/services/worker/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefineDeploymentsListRoute(router *gin.Engine, functions map[string]models.Function, reverseProxy, backend, reverseProxyAdminToken, workerAdminToken string) {
	router.GET("functions/list", func(c *gin.Context) {

		// Declared an empty map interface
		var result map[string]interface{}

		// Unmarshal or Decode the JSON to the interface.
		err := json.Unmarshal([]byte(helpers.JsonFuncList(functions, reverseProxy, backend)), &result)
		if err != nil {
			//fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    "JSON_PARSE_ERROR",
				"message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusAccepted, result)
			//c.JSON(http.StatusAccepted, result)
		}

	})
}
