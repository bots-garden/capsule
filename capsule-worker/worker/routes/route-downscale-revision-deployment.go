package routes

import (
	"fmt"
	"github.com/bots-garden/capsule/capsule-worker/worker/models"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
	"syscall"
)

func RemoveUrlRevisionFromReverseProxy(urlToRemove, functionName, revisionName, reverseProxy, backend, reverseProxyAdminToken string) (status string) {
	/* remove a URL from an existing revision
	   Curl Query
	    curl -v -X DELETE \
	      http://localhost:8888/memory/functions/morgen/blue/url \
	      -H 'content-type: application/json; charset=utf-8' \
	      -d '{"url": "http://localhost:5053"}'
	    echo ""
	    Remark: it's like a downscale
	*/

	client := resty.New()
	bodyStr := `{"url":"` + urlToRemove + `"}`
	resp, err := client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json; charset=utf-8").
		SetHeader("CAPSULE_REVERSE_PROXY_ADMIN_TOKEN", reverseProxyAdminToken).
		SetBody(bodyStr).
		Delete(reverseProxy + "/" + backend + "/functions/" + functionName + "/" + revisionName + "/url")

	if err != nil {
		fmt.Println("😡 when removing the url of the revision from the reverse proxy", err.Error())
		//fmt.Println(bodyStr)
		status = "REVISION_URL_NOT_REMOVED"
	} else {
		fmt.Println("🖐[RemoveUrlRevisionFromReverseProxy]", resp)
		status = "REVISION_URL_REMOVED"
	}
	return status
}

func DefineDownscaleRevisionDeploymentRoute(router *gin.Engine, functions map[string]models.Function, capsulePath string, httpPortCounter int, workerDomain, reverseProxy, backend, reverseProxyAdminToken, workerAdminToken string) {

	// remove a deployment of a revision
	// and kill the associated processes
	router.DELETE("functions/revisions/downscale", func(c *gin.Context) {
		//TODO: check if there is a better practice to handle authentication token
		if len(workerAdminToken) == 0 || c.GetHeader("CAPSULE_WORKER_ADMIN_TOKEN") == workerAdminToken {

			// check json payload parameters
			jsonMap := make(map[string]interface{})
			if err := c.Bind(&jsonMap); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    "JSON_PARSE_ERROR",
					"message": err.Error()})
			} else {

				//TODO: check if the values are empty or not
				functionName := jsonMap["function"].(string)
				revisionName := jsonMap["revision"].(string)

				var selectedPid int
				var message string
				var code string

				// you cannot remove all the process (running modules)
				if len(functions[functionName].Revisions[revisionName].WasmModules) >= 2 {
					selectedPid = 0
					// before removing, kill the associated processes
					for pid, wasmModule := range functions[functionName].Revisions[revisionName].WasmModules {
						fmt.Println("😢 killing :", wasmModule.Pid, wasmModule.LocalUrl)
						err := syscall.Kill(pid, syscall.SIGINT)
						if err != nil {
							fmt.Println("😡 error when stopping the wasm module", err.Error())
						} else {
							selectedPid = pid
							code = "WASM_MODULE_DEPLOYMENT_REMOVED"
							message = "wasm module deployment removed (its process is killed)"

							urlToRemove := wasmModule.LocalUrl

							// remove the url of the revision from the proxy
							RemoveUrlRevisionFromReverseProxy(urlToRemove, functionName, revisionName, reverseProxy, backend, reverseProxyAdminToken)
							// trying to remove url from default revision of the function if it exists
							RemoveUrlRevisionFromReverseProxy(urlToRemove, functionName, "default", reverseProxy, backend, reverseProxyAdminToken)

							// remove the running wasm module from the memory
							delete(functions[functionName].Revisions[revisionName].WasmModules, pid)
						}
						break // we are killing only one running wasm module
					}

				} else {
					selectedPid = 0
					code = "WASM_MODULE_DEPLOYMENT_NOT_REMOVED"
					message = "wasm module deployment not removed (the revision needs at least one running wasm module)"
				}

				//delete(functions[functionName].Revisions, revisionName)

				//TODO: better error handling => add error handling
				c.JSON(http.StatusAccepted, gin.H{
					"code":     code,
					"message":  message,
					"revision": revisionName,
					"function": functionName,
					"pid":      selectedPid})

			}

		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    "KO",
				"from":    "worker",
				"message": "Forbidden"})
		}

	})
}
