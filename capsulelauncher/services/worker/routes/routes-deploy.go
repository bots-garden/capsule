package routes

import (
    "fmt"
    "github.com/bots-garden/capsule/capsulelauncher/services/worker/helpers"
    "github.com/bots-garden/capsule/capsulelauncher/services/worker/models"
    "github.com/gin-gonic/gin"
    "github.com/go-resty/resty/v2"
    "github.com/google/uuid"
    "net/http"
    "os"
    "os/exec"
    "strconv"
)

var httpPortCounter = 10000 //TODO: the starting number could be a parameter
var pidCounter = 0

// AddRunningWasmModuleToRevision updates the functions map
func AddRunningWasmModuleToRevision(functionName, revisionName string, wasmModule models.RunningWasmModule, functions map[string]models.Function) {
    fmt.Println("🟡[scale] Add a new wasm module to an existing revision of an existing function:", revisionName)
    functions[functionName].Revisions[revisionName].WasmModules[wasmModule.Pid] = wasmModule
}

// AddRevisionWithWasmModuleToFunction updates the functions map
func AddRevisionWithWasmModuleToFunction(functionName, revisionName, wasmModuleUrl string, wasmModule models.RunningWasmModule, functions map[string]models.Function) {
    fmt.Println("🟠️ Creation of the revision to an existing function:", revisionName)
    functions[functionName].Revisions[revisionName] = models.Revision{
        Name:            revisionName,
        WasmRegistryUrl: wasmModuleUrl,
        WasmModules: map[int]models.RunningWasmModule{
            wasmModule.Pid: wasmModule,
        },
    }
}

// AddFunctionWithRevisionWithWasmModule updates the functions map
func AddFunctionWithRevisionWithWasmModule(functionName, revisionName, wasmModuleUrl string, wasmModule models.RunningWasmModule, functions map[string]models.Function) {
    fmt.Println("🟣️ First deployment of the function:", functionName)
    fmt.Println("🟣️ Creation of the revision:", revisionName)

    functions[functionName] = models.Function{
        Name: functionName,
        Revisions: map[string]models.Revision{
            revisionName: models.Revision{
                Name:            revisionName,
                WasmRegistryUrl: wasmModuleUrl,
                WasmModules: map[int]models.RunningWasmModule{
                    wasmModule.Pid: wasmModule,
                },
            },
        },
    }
}

func DefineDeployRoute(router *gin.Engine, functions map[string]models.Function, workerDomain, reverseProxy, backend string) {

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
            wasmModuleUrl := jsonMap["downloadUrl"].(string) // the downloadUrl to download the module from the registry
            envVariables := jsonMap["envVariables"]

            //fmt.Println("🚀[envVariables]:", envVariables)
            //fmt.Println("🚀🚀🚀", envVariables == nil)
            //fmt.Println("🚀🚀🚀", envVariables)

            var wasmEnvVariables = make(map[string]string)
            if envVariables == nil {
                wasmEnvVariables = map[string]string{}
            } else {

                for key, value := range envVariables.(map[string]interface{}) {
                    strKey := fmt.Sprintf("%v", key)
                    strValue := fmt.Sprintf("%v", value)
                    wasmEnvVariables[strKey] = strValue
                }

                //wasmEnvVariables = envVariables.(map[string]string)
            }
            fmt.Println("🚀[envVariables]:", wasmEnvVariables)

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
                    // The revision already exists
                    // Then we will add a new running wasm module (scale)
                    wasmModule := models.RunningWasmModule{
                        Pid:          pidCounter, // at the end pid counter will be the id of the process
                        Status:       "🚧wip-not used",
                        LocalUrl:     moduleServerUrl,
                        RemoteUrl:    reverseProxy + "/functions/" + functionName + "/" + revisionName,
                        EnvVariables: wasmEnvVariables,
                    }
                    AddRunningWasmModuleToRevision(functionName, revisionName, wasmModule, functions)

                } else {
                    // The revision does not exist
                    // Then Create a new revision for the function
                    // With a running wasm module
                    wasmModule := models.RunningWasmModule{
                        Pid:          pidCounter, // at the end pid counter will be the id of the process
                        Status:       "🚧wip-not used",
                        LocalUrl:     moduleServerUrl,
                        RemoteUrl:    reverseProxy + "/functions/" + functionName + "/" + revisionName,
                        EnvVariables: wasmEnvVariables,
                    }
                    AddRevisionWithWasmModuleToFunction(functionName, revisionName, wasmModuleUrl, wasmModule, functions)

                }
            } else {

                /*
                   cmd := &exec.Cmd{
                       Path:   "./capsule",
                       Args:   []string{"-url=" + wasmModuleUrl, "-mode=http", "-httpPort=" + strconv.Itoa(httpPortCounter), "-wasm=./tmp/" + uuid.New().String() + ".wasm"},
                       Stdout: os.Stdout,
                       Stderr: os.Stdout,
                   }
                */
                /* =========================================
                   Start a function
                  ========================================= */
                tmpFileName := uuid.New().String() + ".wasm"
                cmd := exec.Command(
                    "./capsule",
                    "-url="+wasmModuleUrl,
                    "-mode=http",
                    "-httpPort="+strconv.Itoa(httpPortCounter),
                    "-wasm=./tmp/"+tmpFileName) //TODO: record this in the list of modules to clean when undeploy
                cmd.Env = os.Environ()
                err := cmd.Start()

                var processStatus string
                if err != nil {
                    fmt.Println("😡 when starting the capsule process", err.Error())
                    processStatus = "NOT_STARTED"
                } else {
                    processStatus = "STARTED"
                }

                fmt.Println("🚀 service started, process Id:", cmd.Process.Pid)

                // TODO: environment variables
                //cmd.Env = append(cmd.Env, "MY_VAR=some_value")

                /* =========================================
                   Register a function to the reverse proxy
                  ========================================= */

                /*

                   # Register the function with the 0.0.0 revision
                   curl -v -X POST \
                     http://localhost:8888/memory/functions/registration \
                     -H 'content-type: application/json; charset=utf-8' \
                     -d '{"function": "hola", "revision": "0.0.0", "url": "http://localhost:7070"}'

                   # Add the default revision
                   curl -v -X POST \
                     http://localhost:8888/memory/functions/hola/revision \
                     -H 'content-type: application/json; charset=utf-8' \
                     -d '{"function": "hola", "revision": "default", "url": "http://localhost:7070"}'


                   # Add the 0.0.1 revision
                   curl -v -X POST \
                     http://localhost:8888/memory/functions/hola/revision \
                     -H 'content-type: application/json; charset=utf-8' \
                     -d '{"function": "hola", "revision": "0.0.1", "url": "http://localhost:7071"}'

                */

                client := resty.New()
                bodyStr := `{"function":"` + functionName + `", "revision":"` + revisionName + `", "url":"` + moduleServerUrl + `"}`
                resp, err := client.
                    R().
                    EnableTrace().
                    SetHeader("Content-Type", "application/json; charset=utf-8").
                    SetBody(bodyStr).
                    Post(reverseProxy + "/" + backend + "/functions/registration")

                if err != nil {
                    fmt.Println("😡 when registering the url to the reverse proxy", err.Error())
                    //fmt.Println(bodyStr)
                    processStatus += "[NOT_REGISTERED]"
                } else {
                    fmt.Println(resp)
                }

                // The function does not exist: this is the first deployment of the function
                // The revision does not exist
                // Then, create the function and the revision
                wasmModule := models.RunningWasmModule{
                    Pid:          cmd.Process.Pid, // at the end pid counter will be the id of the process
                    Status:       processStatus,
                    LocalUrl:     moduleServerUrl,
                    RemoteUrl:    reverseProxy + "/functions/" + functionName + "/" + revisionName,
                    EnvVariables: wasmEnvVariables,
                    TmpFileName:  tmpFileName,
                }
                AddFunctionWithRevisionWithWasmModule(functionName, revisionName, wasmModuleUrl, wasmModule, functions)

            }

            //TODO: do the job for real

            c.JSON(http.StatusAccepted, gin.H{
                "code":     "OK",
                "message":  "Function deployed",
                "function": functionName,
                "revision": revisionName})

        }
    })
}
