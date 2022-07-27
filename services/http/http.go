package capsulehttp

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bots-garden/capsule/services/common"
)

type JsonParameter struct {
	Message string `json:"message"` // change the na
}

type JsonResult struct {
	Value string `json:"value"`
	Error string `json:"error"`
}

//TODO: return more things from JsonResult (eg type of the response)
//! I should do this from the Handle function
//! I need to have several Handle function, or the handle function returns an interface{} (or a result object)

/*
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json' \
  -d '{"message": "Golang 💚 wasm"}'
*/

/*
curl -XPOST -H "content-type:application/json" -d '{"hoge": 1, "fuga": [2,3,4,5]}' localhost:1323

		if err := c.Bind(&json); err != nil {
			return err
		}
		return c.String(http.StatusOK, fmt.Sprintf("%v", json))
*/

func Serve(httpPort string, wasmFile []byte) {

	e := echo.New()

    //TODO: post Raw data

	e.POST("/", func(c echo.Context) error {

        jsonMap := make(map[string]interface{})

		if err := c.Bind(&jsonMap); err != nil {
			return err
		}

        // Convert map to json string
        jsonStr, err := json.Marshal(jsonMap)
        if err != nil {
            fmt.Println(err)
        }
        // TODO handle Error

		// Parameters "setup"
        // Payload
		stringParameterLength := uint64(len(jsonStr))
		stringParameter := jsonStr

        // Headers
        //headers := (map[string][]string) c.Request().Header
        var headersMap = make(map[string]string)
        for key, values := range c.Request().Header {
            headersMap[key]=values[0]
        }
        headersSlice := CreateSliceFromMap(headersMap)

        headersParameter := CreateStringFromSlice(headersSlice, "|")
        headersParameterLength := uint64(len(headersParameter))

		wasmRuntime, wasmModule, ctx := capsule.CreateWasmRuntimeAndModuleInstances(wasmFile)
		defer wasmRuntime.Close(ctx)

		// get the function
		//wasmModuleHandleFunction := wasmModule.ExportedFunction("callHandle")
        wasmModuleHandleFunction := wasmModule.ExportedFunction("callHandleHttp")


		// These are undocumented, but exported. See tinygo-org/tinygo#2788
		malloc := wasmModule.ExportedFunction("malloc")
		free := wasmModule.ExportedFunction("free")
		// https://github.com/tinygo-org/tinygo/issues/2788
		// https://github.com/tinygo-org/tinygo/issues/2787

		// Instead of an arbitrary memory offset, use TinyGo's allocator.
		// 🖐 Notice there is nothing string-specific in this allocation function.
		// The same function could be used to pass binary serialized data to Wasm.
		results, err := malloc.Call(ctx, stringParameterLength)
		if err != nil {
			log.Panicln("💥 out of bounds memory access", err)
		}
		stringParameterPtrPosition := results[0]
		// This pointer is managed by TinyGo, but TinyGo is unaware of external usage.
		// So, we have to free it when finished
		defer free.Call(ctx, stringParameterPtrPosition)

		// The pointer is a linear memory offset, which is where we write the name.
		if !wasmModule.Memory().Write(ctx, uint32(stringParameterPtrPosition), []byte(stringParameter)) {
			log.Panicf("🟥 Memory.Write(%d, %d) out of range of memory size %d",
				stringParameterPtrPosition, stringParameterLength, wasmModule.Memory().Size(ctx))
		}

        // Headers
        resultsHeader, err := malloc.Call(ctx, headersParameterLength)
		if err != nil {
			log.Panicln("💥 out of bounds memory access", err)
		}
		headersParameterPtrPosition := resultsHeader[0]

		defer free.Call(ctx, headersParameterPtrPosition)

		if !wasmModule.Memory().Write(ctx, uint32(headersParameterPtrPosition), []byte(headersParameter)) {
			log.Panicf("🟥 Memory.Write(%d, %d) out of range of memory size %d",
            headersParameterPtrPosition, headersParameterLength, wasmModule.Memory().Size(ctx))
		}
        // End of Headers


		// Finally, This shows how to
		// read-back something allocated by TinyGo.
		handleResultArray, err := wasmModuleHandleFunction.Call(ctx, stringParameterPtrPosition, stringParameterLength, headersParameterPtrPosition, headersParameterLength)
		if err != nil {
			log.Panicln(err)
		}
		// Note: This pointer is still owned by TinyGo, so don't try to free it!
		handleReturnPtrPos, handleReturnSize := capsule.GetPackedPtrPositionAndSize(handleResultArray)

		// The pointer is a linear memory offset, which is where we write the name.
		if bytes, ok := wasmModule.Memory().Read(ctx, handleReturnPtrPos, handleReturnSize); !ok {
			log.Panicf("Memory.Read(%d, %d) out of range of memory size %d",
				handleReturnPtrPos, handleReturnSize, wasmModule.Memory().Size(ctx))
                return c.String(500, "out of range of memory size")
		} else {

            valueStr := string(bytes)
            // check the return value
            if capsule.IsStringError(valueStr) {
                errorMessage, errorCode := capsule.GetStringErrorInfo(valueStr)
                if errorCode == 0 {
                    valueStr = errorMessage
                } else {
                    valueStr = errorMessage + " (" + strconv.Itoa(errorCode) + ")"
                }

                return c.String(500, valueStr)
            } else {
                // TODO: how to check id JSON or RAW (do something like with error)
                return c.String(http.StatusOK, valueStr)
            }

		}

	})
	//https://echo.labstack.com/guide/customization/
	e.HideBanner = true
	e.Start(":" + httpPort)

	//e.Logger.Info(e.Start(":" + httpPort))
	//e.Logger.Fatal(e.Start(":" + httpPort))

}
