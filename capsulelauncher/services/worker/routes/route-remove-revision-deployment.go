package routes

import (
	"fmt"
	"github.com/bots-garden/capsule/capsulelauncher/services/worker/models"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
	"syscall"
)

// delete(functions[functionName].Revisions, revisionName)

func RemoveRevisionFromReverseProxy(functionName, revisionName, reverseProxy, backend string) (status string) {
	/* add revision to a function
	   curl -v -X DELETE \
	     http://localhost:8888/memory/functions/morgen/revision \
	     -H 'content-type: application/json; charset=utf-8' \
	     -d '{"revision": "blue"}'
	   echo ""
	*/

	client := resty.New()
	bodyStr := `{"revision":"` + revisionName + `"}`
	resp, err := client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json; charset=utf-8").
		SetBody(bodyStr).
		Delete(reverseProxy + "/" + backend + "/functions/" + functionName + "/revision")

	if err != nil {
		fmt.Println("😡 when removing the revision from the reverse proxy", err.Error())
		//fmt.Println(bodyStr)
		status = "REVISION_NOT_REMOVED"
	} else {
		fmt.Println("🖐[RemoveRevisionFromReverseProxy]", resp)
		status = "REVISION_REMOVED"
	}
	return status
}

func DefineRemoveRevisionDeploymentRoute(router *gin.Engine, functions map[string]models.Function, capsulePath string, httpPortCounter int, workerDomain, reverseProxy, backend string) {

	// remove a deployment of a revision
	// and kill the associated processes
	router.DELETE("functions/revisions/deployments", func(c *gin.Context) {

		//TODO: add an authentication token

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

			// before removing, kill the associated processes
			for pid, wasmModule := range functions[functionName].Revisions[revisionName].WasmModules {
				fmt.Println("😢 killing :", wasmModule.Pid, wasmModule.LocalUrl)
				err := syscall.Kill(pid, syscall.SIGINT)
				if err != nil {
					fmt.Println("😡 error when stopping the wasm module", err.Error())
				}
			}

			// remove the revision from the proxy
			RemoveRevisionFromReverseProxy(functionName, revisionName, reverseProxy, backend)

			// remove the revision from the memory
			delete(functions[functionName].Revisions, revisionName)

			//TODO: better error handling
			c.JSON(http.StatusAccepted, gin.H{
				"code":     "REVISION_DEPLOYMENT_REMOVED",
				"message":  "revision deployment removed (all processes killed)",
				"revision": revisionName,
				"function": functionName})

		}

	})
}
