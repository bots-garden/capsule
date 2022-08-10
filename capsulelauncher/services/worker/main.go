package worker

import (
	"encoding/json"
	"fmt"
	"github.com/bots-garden/capsule/capsulelauncher/commons"
	"github.com/bots-garden/capsule/capsulelauncher/services/worker/helpers"
	"github.com/bots-garden/capsule/capsulelauncher/services/worker/models"
	"github.com/bots-garden/capsule/capsulelauncher/services/worker/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

var functions = make(map[string]models.Function)

func Serve(httpPort, reverseProxy, workerDomain, backend, crt, key string) {

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

	routes.DefineDeployRoute(router, functions, workerDomain, reverseProxy, backend)

	//TODO: 🚧 WIP cf JsonFuncList
	router.GET("functions/list", func(c *gin.Context) {

		// Declared an empty map interface
		var result map[string]interface{}

		// Unmarshal or Decode the JSON to the interface.
		err := json.Unmarshal([]byte(helpers.JsonFuncList(functions)), &result)
		if err != nil {
			//fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    "JSON_PARSE_ERROR",
				"message": err.Error()})
		} else {
			c.JSON(http.StatusAccepted, result)
		}

	})

	if crt != "" {
		// certs/procyon-registry.local.crt
		// certs/procyon-registry.local.key
		fmt.Println("🚙 Reverse-proxy:", reverseProxy)
		fmt.Println("💊 Capsule (", commons.CapsuleVersion(), ") 🚧 Worker is listening on:", httpPort, "🔐🌍")

		router.RunTLS(":"+httpPort, crt, key)
	} else {
		fmt.Println("🚙 Reverse-proxy:", reverseProxy)
		fmt.Println("💊 Capsule (", commons.CapsuleVersion(), ") 🚧 Worker is listening on:", httpPort, "🌍")
		router.Run(":" + httpPort)
	}

}
