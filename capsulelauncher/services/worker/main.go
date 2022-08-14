package worker

import (
	"fmt"
	"github.com/bots-garden/capsule/capsulelauncher/commons"
	"github.com/bots-garden/capsule/capsulelauncher/services/worker/models"
	"github.com/bots-garden/capsule/capsulelauncher/services/worker/routes"
	"github.com/gin-gonic/gin"
)

var functions = make(map[string]models.Function)
var httpPortCounter int

//TODO: implement reverse proxy authentication token

func Serve(httpPort, capsulePath string, httpPortCounterStart int, reverseProxy, workerDomain, backend, crt, key string) {

	httpPortCounter = httpPortCounterStart

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

	routes.DefineDeployRoute(router, functions, capsulePath, httpPortCounter, workerDomain, reverseProxy, backend)

	routes.DefineSwitchRoutes(router, functions, capsulePath, httpPortCounter, workerDomain, reverseProxy, backend)

	routes.DefineDeploymentsListRoute(router, functions, reverseProxy, backend)

	routes.DefineRemoveRevisionDeploymentRoute(router, functions, capsulePath, httpPortCounter, workerDomain, reverseProxy, backend)

	routes.DefineDownscaleRevisionDeploymentRoute(router, functions, capsulePath, httpPortCounter, workerDomain, reverseProxy, backend)

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
