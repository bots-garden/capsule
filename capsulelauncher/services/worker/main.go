package worker

import (
    "encoding/json"
    "fmt"
    "github.com/bots-garden/capsule/capsulelauncher/commons"
    "github.com/gin-gonic/gin"
    "log"
    "net"
    "net/http"
    "os"
    "reflect"
    "strconv"
    "strings"
)

// GetOutboundIP Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}

type runningWasmModule struct {
    pid       int
    status    string // for future feature(s)
    localUrl  string
    remoteUrl string
}

type revision struct {
    name string // optional

    wasmRegistryUrl string
    wasmModules     map[int]runningWasmModule
}

type function struct {
    name      string
    revisions map[string]revision
}

/*
	birdData := map[string]interface{}{
		"birdSounds": map[string]string{
			"pigeon": "coo",
			"eagle":  "squak",
		},
		"total birds": 2,
	}

*/

func JsonFuncList(funcList map[string]function) string {

    jsonString := `{`

    for functionName, functionElement := range funcList {

        jsonString += `"` + functionName + `":{`

        fmt.Println("Function:", functionElement.name, "key:", functionName)
        fmt.Println("  Revisions:")

        for revisionName, revisionElement := range functionElement.revisions {

            jsonString += `"` + revisionName + `":{`
            jsonString += `"wasmRegistryUrl":"` + revisionElement.wasmRegistryUrl + `",`

            fmt.Println("    ->", revisionElement.name, "key:", revisionName)
            fmt.Println("      - wasmRegistryUrl:", revisionElement.wasmRegistryUrl)
            fmt.Println("      - (running)wasmModules:")

            for idOfProcess, wasmModuleElement := range revisionElement.wasmModules {

                jsonString += `"` + strconv.Itoa(wasmModuleElement.pid) + `":{`
                jsonString += `"localUrl":"` + wasmModuleElement.localUrl + `",`
                jsonString += `"remoteUrl":"` + wasmModuleElement.remoteUrl + `"`

                fmt.Println("        ->", wasmModuleElement.pid, "key:", idOfProcess)
                fmt.Println("          - localUrl", wasmModuleElement.localUrl)
                fmt.Println("          - remoteUrl", wasmModuleElement.remoteUrl)

                jsonString += `},` // end of running module
            }
            // remove the last ","
            jsonString = strings.TrimSuffix(jsonString, ",")

            jsonString += `},` // end of revision
        }
        // remove the last ","
        jsonString = strings.TrimSuffix(jsonString, ",")

        jsonString += `},` // end of function
    }
    // remove the last ","
    jsonString = strings.TrimSuffix(jsonString, ",")

    jsonString += `}` // end of json string

    return jsonString
}

var functions = make(map[string]function)

//var functions = make(map[interface{}]map[interface{}]interface{})
var httpPortCounter = 10000 //TODO: the starting number could be a parameter
var pidCounter = 0

func Serve(httpPort, reverseProxy, workerDomain, crt, key string) {

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
        var moduleServerUrl string
        moduleServerPort := "http" //TODO: handle the case of https

        if len(workerDomain) == 0 {
            // 🔎 discovering automatically the domain (or IP address) of the worker
            nodeName, err := os.Hostname()
            if err != nil {
                moduleServerUrl = moduleServerPort + "://" + GetOutboundIP().String() + ":" + strconv.Itoa(httpPortCounter)

            } else {
                moduleServerUrl = moduleServerPort + "://" + nodeName + ":" + strconv.Itoa(httpPortCounter)
            }
        } else {
            moduleServerUrl = moduleServerPort + "://" + workerDomain + ":" + strconv.Itoa(httpPortCounter)
        }
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

        if reflect.DeepEqual(functions[functionName], function{}) == false {
            // the function already exists
            if reflect.DeepEqual(functions[functionName].revisions[revisionName], revision{}) == false {
                // the revision already exists
                // then we will add a new running wasm module == scale
                fmt.Println("🟡[scale] Add a new wasm module to an existing revision of an existing function:", revisionName)

                functions[functionName].revisions[revisionName].wasmModules[pidCounter] = runningWasmModule{
                    pid:       pidCounter, // at the end pid counter will be the id of the process
                    status:    "🚧wip-not used",
                    localUrl:  moduleServerUrl,
                    remoteUrl: reverseProxy + "/functions/" + functionName + "/" + revisionName,
                }

            } else {
                // the revision does not exist
                // create a new revision for the function
                fmt.Println("🟠️ Creation of the revision to an existing function:", revisionName)
                functions[functionName].revisions[revisionName] = revision{
                    name:            revisionName,
                    wasmRegistryUrl: wasmModuleUrl,
                    wasmModules: map[int]runningWasmModule{
                        pidCounter: { // at the end pid counter will be the id of the process
                            pid:       pidCounter,
                            status:    "🚧wip-not used",
                            localUrl:  moduleServerUrl,
                            remoteUrl: reverseProxy + "/functions/" + functionName + "/" + revisionName,
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
            functions[functionName] = function{
                name: functionName,
                revisions: map[string]revision{
                    revisionName: revision{
                        name:            revisionName,
                        wasmRegistryUrl: wasmModuleUrl,
                        wasmModules: map[int]runningWasmModule{
                            pidCounter: { // at the end pid counter will be the id of the process
                                pid:       pidCounter,
                                status:    "🚧wip-not used",
                                localUrl:  moduleServerUrl,
                                remoteUrl: reverseProxy + "/functions/" + functionName + "/" + revisionName,
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

    //TODO: 🚧 WIP cf JsonFuncList
    router.GET("functions/list", func(c *gin.Context) {

        // Declared an empty map interface
        var result map[string]interface{}

        // Unmarshal or Decode the JSON to the interface.
        err := json.Unmarshal([]byte(JsonFuncList(functions)), &result)
        fmt.Println(err)

        c.JSON(http.StatusAccepted, result)
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
