package routes

import (
	"fmt"
	"github.com/bots-garden/capsule/capsule-worker/worker/models"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
)

// don't remove the revision
// create a default revision from an existing revision
// 🤔 perhaps better: only change the registration to the reverse proxy

/*
- add a flag default
- register to the reverse proxy
- it's not possible to create "physically" a default revision (prevent this)

*/

/* creation of a function with revision
curl -v -X POST \
http://localhost:9999/functions/deploy \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hola",
    "revision": "blue",
    "downloadUrl": "http://localhost:4999/k33g/hola/0.0.1/hola.wasm",
    "envVariables": {
        "MESSAGE": "🔵 Blue revision of Hola",
        "TOKEN": "this is not a header token"
    }
}
EOF
echo ""
*/

/*
# Add the default revision
curl -v -X POST \
  http://localhost:8888/memory/functions/hola/revision \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"function": "hola", "revision": "default", "url": "http://localhost:7070"}'

*/

/*
- I need the url of a specific revision of a function
- I will create a default revision with this url (only on the reverse-proxy side)
- change the value of the `isDefaultRevision` flag of the revision in the function list
- TODO: check if I need to remove the default revision before change it: we can have only one default revision
*/

/* revision switch
curl -v -X POST \
http://localhost:9999/functions/set_default_revision \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hola",
    "revision": "blue"
}
EOF
echo ""
*/

func ChangeDefaultRevisionFlag() {
	// worker side
}

// CreateDefaultRevisionWithFirstModuleUrl : reverse-proxy side
func CreateDefaultRevisionWithFirstModuleUrl(functionName, revisionName, moduleServerUrl, reverseProxy, backend, reverseProxyAdminToken string) (status string) {
	/* add revision to a function
	   curl -v -X POST \
	     http://localhost:8888/memory/functions/morgen/revision \
	     -H 'content-type: application/json; charset=utf-8' \
	     -d '{"revision": "blue", "url": "http://localhost:5051"}'
	   echo ""
	*/

	client := resty.New()
	bodyStr := `{"function":"` + functionName + `", "revision":"` + revisionName + `", "url":"` + moduleServerUrl + `"}`
	resp, err := client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json; charset=utf-8").
		SetHeader("CAPSULE_REVERSE_PROXY_ADMIN_TOKEN", reverseProxyAdminToken).
		SetBody(bodyStr).
		Post(reverseProxy + "/" + backend + "/functions/" + functionName + "/revision")

	if err != nil {
		fmt.Println("😡 when registering the revision to the reverse proxy", err.Error())
		//fmt.Println(bodyStr)
		status = "DEFAULT_REVISION_NOT_REGISTERED"
	} else {
		fmt.Println("🖐[CreateDefaultRevisionWithFirstModuleUrl]", resp)
		status = "DEFAULT_REVISION_REGISTERED"
	}
	return status
}

// RemoveDefaultRevisionFromReverseProxy : remove the default revision from the reverse proxy
func RemoveDefaultRevisionFromReverseProxy(functionName, reverseProxy, backend, reverseProxyAdminToken string) (status string) {
	/* add revision to a function
	   curl -v -X DELETE \
	     http://localhost:8888/memory/functions/morgen/revision \
	     -H 'content-type: application/json; charset=utf-8' \
	     -d '{"revision": "blue"}'
	   echo ""
	*/
	revisionName := "default"

	client := resty.New()
	bodyStr := `{"revision":"` + revisionName + `"}`
	resp, err := client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json; charset=utf-8").
		SetHeader("CAPSULE_REVERSE_PROXY_ADMIN_TOKEN", reverseProxyAdminToken).
		SetBody(bodyStr).
		Delete(reverseProxy + "/" + backend + "/functions/" + functionName + "/revision")

	if err != nil {
		fmt.Println("😡 when removing the revision from the reverse proxy", err.Error())
		//fmt.Println(bodyStr)
		status = "DEFAULT_REVISION_NOT_REMOVED"
	} else {
		fmt.Println("🖐[RemoveDefaultRevisionFromReverseProxy]", resp)
		status = "DEFAULT_REVISION_REMOVED"
	}
	return status
}

func AddUrlToRevision(functionName, revisionName, moduleServerUrl, reverseProxy, backend, reverseProxyAdminToken string) (status string) {
	/* add url to a revision
	   curl -v -X POST \
	   http://localhost:8888/memory/functions/hola/default/url \
	   -H 'content-type: application/json; charset=utf-8' \
	   -d '{"url": "http://localhost:10003"}'
	   echo ""
	*/
	fmt.Println("🟢", functionName, revisionName, moduleServerUrl)

	client := resty.New()
	bodyStr := `{"url":"` + moduleServerUrl + `"}`
	resp, err := client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json; charset=utf-8").
		SetHeader("CAPSULE_REVERSE_PROXY_ADMIN_TOKEN", reverseProxyAdminToken).
		SetBody(bodyStr).
		Post(reverseProxy + "/" + backend + "/functions/" + functionName + "/" + revisionName + "/url")

	if err != nil {
		fmt.Println("😡 when registering the url to the reverse proxy", err.Error())
		//fmt.Println(bodyStr)
		status = "URL_NOT_REGISTERED"
	} else {
		fmt.Println("🖐[AddUrlToRevision]", resp)
		status = "URL_REGISTERED"
	}
	return status
}

func DefineSwitchRoutes(router *gin.Engine, functions map[string]models.Function, capsulePath string, httpPortCounter int, workerDomain, reverseProxy, backend, reverseProxyAdminToken, workerAdminToken string) {
	//Delete on the reverse-proxy-side
	router.DELETE("functions/remove_default_revision", func(c *gin.Context) {
		//TODO: check if there is a better practice to handle authentication token
		if len(workerAdminToken) == 0 || CheckWorkerAdminToken(c, workerAdminToken) {
			// check json payload parameters
			jsonMap := make(map[string]interface{})
			if err := c.Bind(&jsonMap); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    "JSON_PARSE_ERROR",
					"message": err.Error()})
			} else {

				//TODO: check if the values are empty or not
				functionName := jsonMap["function"].(string)
				RemoveDefaultRevisionFromReverseProxy(functionName, reverseProxy, backend, reverseProxyAdminToken)

				//TODO: better error handling
				c.JSON(http.StatusAccepted, gin.H{
					"code":     "DEFAULT_REVISION_REMOVED",
					"message":  "Default revision removed",
					"function": functionName})

			}
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    "KO",
				"from":    "worker",
				"message": "Forbidden"})
		}

	})

	router.POST("functions/set_default_revision", func(c *gin.Context) {
		//TODO: check if there is a better practice to handle authentication token
		if len(workerAdminToken) == 0 || CheckWorkerAdminToken(c, workerAdminToken) {

			// check json payload parameters
			jsonMap := make(map[string]interface{})
			if err := c.Bind(&jsonMap); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    "JSON_PARSE_ERROR",
					"message": err.Error()})
			} else {

				//TODO: check if the values are empty or not
				functionName := jsonMap["function"].(string)
				currentRevisionName := jsonMap["revision"].(string)

				//🖐 a revision can have several modules
				wasmModulesOfTheRevision := functions[functionName].Revisions[currentRevisionName].WasmModules

				var urlList []string
				for _, wasmModule := range wasmModulesOfTheRevision {
					urlList = append(urlList, wasmModule.LocalUrl)
				}
				moduleServerUrl := urlList[0]

				//TODO 🖐 remove the flag isDefault revision from the Revision structure and retrieve the information from the reverse-proxy

				status := CreateDefaultRevisionWithFirstModuleUrl(functionName, "default", moduleServerUrl, reverseProxy, backend, reverseProxyAdminToken)
				fmt.Println(status)

				// Add the other module urls to the revision

				if len(urlList) > 1 {
					for idx, url := range urlList {
						if idx != 0 {
							AddUrlToRevision(functionName, "default", url, reverseProxy, backend, reverseProxyAdminToken)
						}
					}
				}
				//TODO: better error handling (get status)
				c.JSON(http.StatusAccepted, gin.H{
					"code":     "DEFAULT_REVISION_REGISTERED",
					"message":  "Default revision registered",
					"revision": currentRevisionName,
					"function": functionName,
					"url":      reverseProxy + "/functions/" + functionName})
			}
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    "KO",
				"from":    "worker",
				"message": "Forbidden"})
		}

	})

}
