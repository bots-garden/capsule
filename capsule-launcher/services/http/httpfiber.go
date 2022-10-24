package capsulehttp

import (
    "context"
    "fmt"
    "github.com/bots-garden/capsule/capsule-launcher/hostfunctions"
    capsule "github.com/bots-garden/capsule/capsule-launcher/services/wasmrt"
    "github.com/bots-garden/capsule/commons"
    "github.com/gofiber/fiber/v2"
    "net/http"
    "os/signal"
    "syscall"
    "time"
)

type RemoteWasmModule struct {
    Url  string `json:"url"`
    Path string `json:"path"`
}

func FiberServe(httpPort string, wasmFileModule []byte, crt, key string) {

    // to help to hot reload a wasm module
    wasmFile := wasmFileModule

    // Create context that listens for the interrupt signal from the OS.
    ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
    defer stop()

    // Store some information available for the wasm modules
    hostfunctions.HostInformation = `{"httpPort":` + httpPort + `,"capsuleVersion":"` + commons.CapsuleVersion() + `"}`

    // This will call the OnLoad function of the wasm module if it exists
    /*
       //export OnLoad
       func OnLoad() {
           hf.Log("👋 from the OnLoad function")
       }
    */
    capsule.CallExportedOnLoad(wasmFile)

    app := fiber.New(fiber.Config{
        DisableStartupMessage: true,
        //DisableKeepalive:      true,
        //Concurrency:           100000,
    })

    /* TODO: ... one day
       v, _ := mem.VirtualMemory()

       app.Get("/host-metrics", func(c *fiber.Ctx) error {
           jsonMap := make(map[string]interface{})
           json.Unmarshal([]byte(v.String()), &jsonMap)
           c.Status(http.StatusOK)
           return c.JSON(jsonMap)
       })

       app.Get("/health", func(c *fiber.Ctx) error {
           c.Status(http.StatusOK)
           return c.SendString("OK")
       })
    */

    app.All("/", func(c *fiber.Ctx) error {
        reqId := hostfunctions.StoreRequestParams(c)

        wasmModule, wasmFunction, wasmCtx := capsule.GetModuleFunctionForHttp(wasmFile)

        bytes, err := capsule.ExecHandleFunctionForHttp(wasmFunction, wasmModule, wasmCtx, uint64(reqId))
        if err != nil {
            c.Status(500)
            return c.SendString("out of range of memory size")
        }

        bodyStr, headers := GetBodyAndHeaders(bytes, c)

        hostfunctions.DeleteRequestParams(reqId)

        // check the return value
        if commons.IsErrorString(bodyStr) {
            return SendErrorMessage(bodyStr, headers, c)
        } else if IsBodyString(bodyStr) {
            return SendJsonMessage(bodyStr, headers, c)
        } else {
            c.Status(http.StatusOK)
            return c.SendString(bodyStr)
        }

    })

    // 🖐 use this at your own risk
    // 🖐 this feature is subject to change

    /*
       CAPSULE_RELOAD_TOKEN

       curl -v -X POST \
         http://localhost:7070/load-wasm-module \
         -H 'content-type: application/json; charset=utf-8' \
         -d '{"url": "http://localhost:9090/hello.wasm", "path": "./tmp/hello.wasm"}'
         echo ""
    */

    app.Post("/load-wasm-module", func(c *fiber.Ctx) error {

        reloadWasmFile := func() error {

            wm := new(RemoteWasmModule)

            if err := c.BodyParser(wm); err != nil {
                c.Status(500)
                return c.SendString("😡[/load-wasm-module] " + err.Error())
            }

            var errWasmFile error
            wasmFile, errWasmFile = capsule.GetWasmFileFromUrl(wm.Url, wm.Path)

            if errWasmFile != nil {
                c.Status(500)
                return c.SendString("😡[/load-wasm-module] " + errWasmFile.Error())
            }

            c.Status(http.StatusOK)
            return c.SendString("🙂 " + wm.Url + " loaded")
        }

        headerTokenReload := GetReloadTokenFromHeadersRequest(c)
        envVarTokenReload := commons.GetEnv("CAPSULE_RELOAD_TOKEN", "")

        if envVarTokenReload != "" { // you need to add a token to the header request
            if headerTokenReload == envVarTokenReload {
                return reloadWasmFile()
            } else {
                // not authorized: 401 Unauthorized
                c.Status(401)
                return c.SendString("😡[/load-wasm-module] Unauthorized")
            }
        } else { // you don't need a token
            return reloadWasmFile()
        }

    })

    go func() {
        if crt != "" {
            // certs/procyon-registry.local.crt
            // certs/procyon-registry.local.key
            fmt.Println("💊 Capsule (", commons.CapsuleVersion(), ") http server is listening on:", httpPort, "🔐🌍")
            app.ListenTLS(":"+httpPort, crt, key)

        } else {
            fmt.Println("💊 Capsule (", commons.CapsuleVersion(), ") http server is listening on:", httpPort, "🌍")
            app.Listen(":" + httpPort)
        }
    }()

    //+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
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
