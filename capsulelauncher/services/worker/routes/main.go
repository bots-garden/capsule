package routes

import (
	"fmt"
	"github.com/bots-garden/capsule/capsulelauncher/services/worker/helpers"
	"github.com/bots-garden/capsule/capsulelauncher/services/worker/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var httpPortCounter = 10000 //TODO: the starting number could be a parameter
var pidCounter = 0

func DefineDeployRoute(router *gin.Engine, functions map[string]models.Function, workerDomain, reverseProxy string) {

	/*
	    ==============================================================
	    Deploy a new function:
	      - download (from the registry) and start the wasm module
	      - register to the reverse proxy
	    🖐🖐🖐 Revisions and Tags are not necessarily the same thing
	    ==============================================================
	   Curl Query:
	      curl -v -X POST \
	      http://localhost:9999/functions/deploy \
	      -H 'content-type: application/json; charset=utf-8' \
	      -d '{"function": "hello", "revision": "default", "downloadUrl": "http://localhost:4999/k33g/hello/0.0.1/hello.wasm"}'
	      echo ""

	   How to pass the environment variables???
	   If I call it 2 times, it scales
	*/
	router.POST("functions/deploy", func(c *gin.Context) {
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
		wasmModuleUrl := jsonMap["downloadUrl"].(string) // the downloadUrl to download the module from the registry

		fmt.Println("⏳ downloading from:", wasmModuleUrl)
		fmt.Println("🚀 starting on http port:", httpPortCounter)
		httpPortCounter += 1
		pidCounter += 1 //TODO: this is temporary

		//🖐 we need the IP address of the worker (for the registration with the reverse proxy)
		//or domain name
		//because the worker and the module are on the same machine
		//but not necessarily the reverse proxy
		//🤔 how to start a module with https??? (or only the reverse proxy???)

		moduleServerPort := "http" //TODO: handle the case of https
		moduleServerUrl := helpers.GetModuleServerUrl(workerDomain, moduleServerPort, httpPortCounter)

		fmt.Println("🌍 the 💊 Capsule module is running at:", moduleServerUrl)

		// TODO: 🖐🖐🖐🖐🖐 we need to be able to pass environment variables to the module
		// the environment variables are used per runningWasmModule
		/*
		    Start in http mode:
		    MESSAGE="💊 Capsule Rocks 🚀" go run main.go \
		   -wasm=../wasm_modules/capsule-hello-post/hello.wasm \
		   -mode=http \
		   -httpPort=7070

		*/
		fmt.Println("📝 registering to the reverse proxy:", reverseProxy)
		fmt.Println("🎉 you can call the function at:", reverseProxy+"/functions/"+functionName+"/"+revisionName)

		fmt.Println("👨🏻‍💻 updating the list of the functions")

		if models.IsFunctionExist(functionName, functions) == true {
			// the function already exists
			if models.IsRevisionExist(functionName, revisionName, functions) == true {
				// the revision already exists
				// then we will add a new running wasm module == scale
				fmt.Println("🟡[scale] Add a new wasm module to an existing revision of an existing function:", revisionName)

				functions[functionName].Revisions[revisionName].WasmModules[pidCounter] = models.RunningWasmModule{
					Pid:       pidCounter, // at the end pid counter will be the id of the process
					Status:    "🚧wip-not used",
					LocalUrl:  moduleServerUrl,
					RemoteUrl: reverseProxy + "/functions/" + functionName + "/" + revisionName,
				}

			} else {
				// the revision does not exist
				// create a new revision for the function
				fmt.Println("🟠️ Creation of the revision to an existing function:", revisionName)
				functions[functionName].Revisions[revisionName] = models.Revision{
					Name:            revisionName,
					WasmRegistryUrl: wasmModuleUrl,
					WasmModules: map[int]models.RunningWasmModule{
						pidCounter: { // at the end pid counter will be the id of the process
							Pid:       pidCounter,
							Status:    "🚧wip-not used",
							LocalUrl:  moduleServerUrl,
							RemoteUrl: reverseProxy + "/functions/" + functionName + "/" + revisionName,
						},
					},
				}
			}
		} else {
			// the function does not exist
			// this is the first deployment of the function
			fmt.Println("Ⓜ️️ First deployment of the function:", functionName)
			fmt.Println("Ⓜ️ Creation of the revision:", revisionName)
			// the revision does not exist
			// create the function and the revision
			functions[functionName] = models.Function{
				Name: functionName,
				Revisions: map[string]models.Revision{
					revisionName: models.Revision{
						Name:            revisionName,
						WasmRegistryUrl: wasmModuleUrl,
						WasmModules: map[int]models.RunningWasmModule{
							pidCounter: { // at the end pid counter will be the id of the process
								Pid:       pidCounter,
								Status:    "🚧wip-not used",
								LocalUrl:  moduleServerUrl,
								RemoteUrl: reverseProxy + "/functions/" + functionName + "/" + revisionName,
							},
						},
					},
				},
			}

		}

		//fmt.Println(functions)

		//TODO: do the job for real

		c.JSON(http.StatusAccepted, gin.H{
			"code":     "OK",
			"message":  "Function deployed",
			"function": functionName,
			"revision": revisionName})

	})
}
