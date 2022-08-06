package capsulehttp

import (
	"encoding/json"
	"github.com/bots-garden/capsule/capsulelauncher/commons"
	capsule "github.com/bots-garden/capsule/capsulelauncher/services/wasmrt"
	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/v3/mem"
	"net/http"
)

func Serve(httpPort string, wasmFile []byte) {

	v, _ := mem.VirtualMemory()

	e := echo.New()

	e.GET("/host-metrics", func(c echo.Context) error {

		jsonMap := make(map[string]interface{})
		json.Unmarshal([]byte(v.String()), &jsonMap)

		return c.JSON(http.StatusOK, jsonMap)
	})

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	//TODO: be able to get the query string from the wasm module
	// we need to be able to return json, html, txt
	e.GET("/", func(c echo.Context) error {
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
			return c.String(500, "out of range of memory size")
		}
		bodyStr, headers := GetBodyAndHeaders(bytes, c)

		// check the return value
		if commons.IsErrorString(bodyStr) {
			return SendErrorMessage(bodyStr, headers, c)
		}

		if IsBodyString(bodyStr) {
			return SendBodyMessage(bodyStr, headers, c)
		}

		return c.String(http.StatusOK, bodyStr)

	})

	e.POST("/", func(c echo.Context) error {

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
			return c.String(500, "out of range of memory size")
		}

		bodyStr, headers := GetBodyAndHeaders(bytes, c)

		// check the return value
		if commons.IsErrorString(bodyStr) {
			return SendErrorMessage(bodyStr, headers, c)
		}

		if IsBodyString(bodyStr) {
			return SendJsonMessage(bodyStr, headers, c)
		}

		return c.String(http.StatusOK, bodyStr)

	})
	//https://echo.labstack.com/guide/customization/
	e.HideBanner = true
	e.Start(":" + httpPort)

	//e.Logger.Info(e.Start(":" + httpPort))
	//e.Logger.Fatal(e.Start(":" + httpPort))

}
