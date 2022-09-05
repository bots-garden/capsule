package capsulehttp

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/bots-garden/capsule/capsule-launcher/hostfunctions"
    "github.com/bots-garden/capsule/capsule-launcher/services/wasmrt"
    "github.com/bots-garden/capsule/commons"
    "github.com/gin-gonic/gin"
    "github.com/shirou/gopsutil/v3/mem"
    "net/http"
    "os/signal"
    "syscall"
    "time"
)

func Serve(httpPort string, wasmFile []byte, crt, key string) {
    //fmt.Println("👋[checking]capsulehttp.Serve")

    // Create context that listens for the interrupt signal from the OS.
    ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
    defer stop()

    hostfunctions.HostInformation = `{"httpPort":` + httpPort + `,"capsuleVersion":"` + commons.CapsuleVersion() + `"}`

    v, _ := mem.VirtualMemory()

    if commons.GetEnv("DEBUG", "false") == "false" {
        gin.SetMode(gin.ReleaseMode)
    } else {
        gin.SetMode(gin.DebugMode)
    }

    // === Call the OnLoad function of the wasm module ===
    /*
       It happens only if you add this code to the wasm module
       //export OnLoad
       func OnLoad() {
           hf.Log("👋 from the OnLoad function")
       }
    */
    capsule.CallExportedOnLoad(wasmFile)

    router := gin.New()

    router.GET("/host-metrics", func(c *gin.Context) {
        jsonMap := make(map[string]interface{})
        json.Unmarshal([]byte(v.String()), &jsonMap)
        c.JSON(http.StatusOK, jsonMap)
    })

    router.GET("/health", func(c *gin.Context) {
        c.String(http.StatusOK, "OK")
    })

    router.GET("/", func(c *gin.Context) {

        jsonStr, _ := GetJsonStringFromPayloadRequest(c)
        headersStr := GetHeadersStringFromHeadersRequest(c)
        uri := c.Request.RequestURI
        method := c.Request.Method

        wasmRuntime, wasmModule, wasmFunction, ctx := capsule.GetNewWasmRuntimeForHttp(wasmFile)
        defer wasmRuntime.Close(ctx)

        uriPos, uriLen, free, err := capsule.ReserveMemorySpaceFor(uri, wasmModule, ctx)
        defer free.Call(ctx, uriPos)

        jsonStrPos, jsonStrLen, free, err := capsule.ReserveMemorySpaceFor(jsonStr, wasmModule, ctx)
        defer free.Call(ctx, jsonStrPos)

        headersStrPos, headersStrLen, free, err := capsule.ReserveMemorySpaceFor(headersStr, wasmModule, ctx)
        defer free.Call(ctx, headersStrPos)

        methodPos, methodLen, free, err := capsule.ReserveMemorySpaceFor(method, wasmModule, ctx)
        defer free.Call(ctx, methodPos)

        bytes, err := capsule.ExecHandleFunction(wasmFunction, wasmModule, ctx, jsonStrPos, jsonStrLen, uriPos, uriLen, headersStrPos, headersStrLen, methodPos, methodLen)
        if err != nil {
            c.String(500, "out of range of memory size")
        }
        bodyStr, headers := GetBodyAndHeaders(bytes, c)

        // check the return value
        if commons.IsErrorString(bodyStr) {
            SendErrorMessage(bodyStr, headers, c)
        } else if IsBodyString(bodyStr) {
            SendBodyMessage(bodyStr, headers, c)
        } else {
            c.String(http.StatusOK, bodyStr)
        }

    })

    router.POST("/", func(c *gin.Context) {

        // Parameters "setup"
        jsonStr, _ := GetJsonStringFromPayloadRequest(c)
        headersStr := GetHeadersStringFromHeadersRequest(c)
        uri := c.Request.RequestURI
        method := c.Request.Method

        wasmRuntime, wasmModule, wasmFunction, ctx := capsule.GetNewWasmRuntimeForHttp(wasmFile)
        defer wasmRuntime.Close(ctx)

        uriPos, uriLen, free, err := capsule.ReserveMemorySpaceFor(uri, wasmModule, ctx)
        defer free.Call(ctx, uriPos)

        jsonStrPos, jsonStrLen, free, err := capsule.ReserveMemorySpaceFor(jsonStr, wasmModule, ctx)
        defer free.Call(ctx, jsonStrPos)

        headersStrPos, headersStrLen, free, err := capsule.ReserveMemorySpaceFor(headersStr, wasmModule, ctx)
        defer free.Call(ctx, headersStrPos)

        methodPos, methodLen, free, err := capsule.ReserveMemorySpaceFor(method, wasmModule, ctx)
        defer free.Call(ctx, methodPos)

        bytes, err := capsule.ExecHandleFunction(wasmFunction, wasmModule, ctx, jsonStrPos, jsonStrLen, uriPos, uriLen, headersStrPos, headersStrLen, methodPos, methodLen)

        if err != nil {
            c.String(500, "out of range of memory size")
        }

        bodyStr, headers := GetBodyAndHeaders(bytes, c)

        // check the return value
        if commons.IsErrorString(bodyStr) {
            SendErrorMessage(bodyStr, headers, c)
        } else if IsBodyString(bodyStr) {
            SendJsonMessage(bodyStr, headers, c)
        } else {
            c.String(http.StatusOK, bodyStr)
        }

        //c.String(http.StatusOK, bodyStr)

    })

    go func() {
        if crt != "" {
            // certs/procyon-registry.local.crt
            // certs/procyon-registry.local.key
            fmt.Println("💊 Capsule (", commons.CapsuleVersion(), ") http server is listening on:", httpPort, "🔐🌍")

            router.RunTLS(":"+httpPort, crt, key)

        } else {
            fmt.Println("💊 Capsule (", commons.CapsuleVersion(), ") http server is listening on:", httpPort, "🌍")
            router.Run(":" + httpPort)
        }
    }()

    // Listen for the interrupt signal.
    <-ctx.Done()

    // Restore default behavior on the interrupt signal and notify user of shutdown.
    stop()
    fmt.Println("💊 Capsule shutting down gracefully ...")

    // === Call the OnExit function of the wasm module ===
    /*
       It happens only if you add this code to the wasm module
       //export OnExit
       func OnExit() {
           hf.Log("👋 from the OnExit function")
       }
    */
    capsule.CallExportedOnExit(wasmFile)

    // The context is used to inform the server it has 5 seconds to finish
    // the request it is currently handling
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    fmt.Println("💊 Capsule exiting")

}
