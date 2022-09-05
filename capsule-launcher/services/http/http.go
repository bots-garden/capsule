package capsulehttp

import (
	"encoding/json"
	"fmt"
	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions"
	"github.com/bots-garden/capsule/capsule-launcher/services/wasmrt"
	"github.com/bots-garden/capsule/commons"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/mem"
	"net/http"
)

func Serve(httpPort string, wasmFile []byte, crt, key string) {
	//fmt.Println("👋[checking]capsulehttp.Serve")

	hostfunctions.HostInformation = `{"httpPort":` + httpPort + `}`

	v, _ := mem.VirtualMemory()

	if commons.GetEnv("DEBUG", "false") == "false" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// === call the onLpad Method of the wasm module ===

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

	if crt != "" {
		// certs/procyon-registry.local.crt
		// certs/procyon-registry.local.key
		fmt.Println("💊 Capsule (", commons.CapsuleVersion(), ") http server is listening on:", httpPort, "🔐🌍")

		router.RunTLS(":"+httpPort, crt, key)
	} else {
		fmt.Println("💊 Capsule (", commons.CapsuleVersion(), ") http server is listening on:", httpPort, "🌍")
		router.Run(":" + httpPort)
	}

}
