package capsulehttp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions"
	capsule "github.com/bots-garden/capsule/capsule-launcher/services/wasmrt"
	"github.com/bots-garden/capsule/commons"
	"github.com/gofiber/fiber/v2"
	"github.com/shirou/gopsutil/v3/mem"
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

	hostfunctions.HostInformation = `{"httpPort":` + httpPort + `,"capsuleVersion":"` + commons.CapsuleVersion() + `"}`
	v, _ := mem.VirtualMemory()

	// OnLoad
	capsule.CallExportedOnLoad(wasmFile)

	//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	// host-metrics
	app.Get("/host-metrics", func(c *fiber.Ctx) error {
		jsonMap := make(map[string]interface{})
		json.Unmarshal([]byte(v.String()), &jsonMap)
		c.Status(http.StatusOK)
		return c.JSON(jsonMap)
	})

	// health
	app.Get("/health", func(c *fiber.Ctx) error {
		c.Status(http.StatusOK)
		return c.SendString("OK")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		jsonStr := string(c.Body())

		headersStr := GetHeadersStringFromHeadersRequest(c)
		uri := c.Request().URI().String()
		method := c.Method()

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
			c.Status(500)
			return c.SendString("out of range of memory size")
		}
		bodyStr, headers := GetBodyAndHeaders(bytes, c)

		// check the return value
		if commons.IsErrorString(bodyStr) {
			return SendErrorMessage(bodyStr, headers, c)
		} else if IsBodyString(bodyStr) {
			return SendBodyMessage(bodyStr, headers, c)
		} else {
			c.Status(http.StatusOK)
			return c.SendString(bodyStr)
		}
	})

	app.Post("/", func(c *fiber.Ctx) error {
		jsonStr := string(c.Body())

		headersStr := GetHeadersStringFromHeadersRequest(c)
		uri := c.Request().URI().String()
		method := c.Method()

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
			c.Status(500)
			return c.SendString("out of range of memory size")
		}
		bodyStr, headers := GetBodyAndHeaders(bytes, c)

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

	// https://docs.gofiber.io/api/ctx#bodyparser
	//TODO: protect with token

	// 🖐 use this at your own risk
	// 🖐 this feature is subject to change

	/*
	   curl -v -X POST \
	     http://localhost:7070/load-wasm-module \
	     -H 'content-type: application/json; charset=utf-8' \
	     -d '{"url": "http://localhost:9090/hello.wasm", "path": "./tmp/hello.wasm"}'
	     echo ""
	*/

	app.Post("/load-wasm-module", func(c *fiber.Ctx) error {
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
	})

	go func() {
		if crt != "" {
			//TODO: cert & key

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
