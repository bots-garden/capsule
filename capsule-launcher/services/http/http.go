package capsulehttp

import (
	"encoding/json"
	"fmt"
	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions"
	capsule "github.com/bots-garden/capsule/capsule-launcher/services/wasmrt"
	"github.com/bots-garden/capsule/commons"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/mem"
	"net/http"
)

func Serve(httpPort string, wasmFile []byte, crt, key string) {

	hostfunctions.HostInformation = `{"httpPort":` + httpPort + `}`

	v, _ := mem.VirtualMemory()

	if commons.GetEnv("DEBUG", "false") == "false" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()

	router.GET("/host-metrics", func(c *gin.Context) {
		jsonMap := make(map[string]interface{})
		json.Unmarshal([]byte(v.String()), &jsonMap)
		c.JSON(http.StatusOK, jsonMap)
	})

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	//TODO: be able to get the query string from the wasm module
	// we need to be able to return json, html, txt
	router.GET("/", func(c *gin.Context) {
		wasmRuntime, wasmModule, wasmFunction, ctx := capsule.GetNewWasmRuntimeForHttp(wasmFile)
		defer wasmRuntime.Close(ctx)

		query := "empty" //🚧
		queryPos, queryLen, free, err := capsule.ReserveMemorySpaceFor(query, wasmModule, ctx)
		defer free.Call(ctx, queryPos)

		headersStr := GetHeadersStringFromHeadersRequest(c)
		headersStrPos, headersStrLen, free, err := capsule.ReserveMemorySpaceFor(headersStr, wasmModule, ctx)
		defer free.Call(ctx, headersStrPos)

		bytes, err := capsule.ExecHandleFunction(wasmFunction, wasmModule, ctx, queryPos, queryLen, headersStrPos, headersStrLen)
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

		//time.Sleep(500 * time.Millisecond)

		wasmRuntime, wasmModule, wasmFunction, ctx := capsule.GetNewWasmRuntimeForHttp(wasmFile)

		defer wasmRuntime.Close(ctx)

		jsonStrPos, jsonStrLen, free, err := capsule.ReserveMemorySpaceFor(jsonStr, wasmModule, ctx)

		defer free.Call(ctx, jsonStrPos)

		headersStrPos, headersStrLen, free, err := capsule.ReserveMemorySpaceFor(headersStr, wasmModule, ctx)

		defer free.Call(ctx, headersStrPos)

		bytes, err := capsule.ExecHandleFunction(wasmFunction, wasmModule, ctx, jsonStrPos, jsonStrLen, headersStrPos, headersStrLen)

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
