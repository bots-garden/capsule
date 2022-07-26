package capsulehttpecho

import (
	"context"
	"log"

	//"math/rand"
	"net/http"
  "github.com/labstack/echo/v4"

	helpers "github.com/bots-garden/capsule/helpers/tools"
	"github.com/bots-garden/capsule/host_functions"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/wasi_snapshot_preview1"
)

type JsonParameter struct {
	Message string `json:"message"` // change the name ? 🤔
}

type JsonResult struct {
  Value string `json:"value"`
  Error string `json:"error"`
}

/*
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json' \
  -d '{"message": "Golang 💚 wasm"}'
*/
func createWasmRuntime(ctx context.Context) wazero.Runtime {

	wasmRuntime := wazero.NewRuntimeWithConfig(wazero.NewRuntimeConfig().WithWasmCore2())

	// 🏠 Add host functions
	_, errEnv := wasmRuntime.NewModuleBuilder("env").
		ExportFunction("hostLogString", host_functions.LogString).
		ExportFunction("hostGetHostInformation", host_functions.GetHostInformation).
		ExportFunction("hostPing", host_functions.Ping).
		Instantiate(ctx, wasmRuntime)

	if errEnv != nil {
		log.Panicln("🔴 Error with env module and host function(s):", errEnv)
	}

	_, errInstantiate := wasi_snapshot_preview1.Instantiate(ctx, wasmRuntime)
	if errInstantiate != nil {
		log.Panicln("🔴 Error with Instantiate:", errInstantiate)
	}

	return wasmRuntime
}

func createWasmRuntimeAndModuleInstances(wasmFile []byte) (wazero.Runtime, api.Module, context.Context) {
	// Choose the context to use for function calls.
	ctx := context.Background()

	wasmRuntime := createWasmRuntime(ctx)
	//defer wasmRuntime.Close(ctx) // This closes everything this Runtime created.

	// 🥚 Instantiate the wasm module (from the wasm file)
	wasmModule, errInstanceWasmModule := wasmRuntime.InstantiateModuleFromBinary(ctx, wasmFile)
	if errInstanceWasmModule != nil {
		log.Panicln("🔴 Error while creating module instance:", errInstanceWasmModule)
	}
	return wasmRuntime, wasmModule, ctx
}

func Serve(httpPort string, wasmFile []byte) {

	e := echo.New()

  e.POST("/", func(c echo.Context) error {
    jsonParameter :=  new(JsonParameter)
		if err := c.Bind(jsonParameter); err != nil {
			return err
		}

    jsonResult := new(JsonResult)

		// Parameter "setup"
		stringParameterLength := uint64(len(jsonParameter.Message))
		stringParameter := jsonParameter.Message

		wasmRuntime, wasmModule, ctx := createWasmRuntimeAndModuleInstances(wasmFile)
		defer wasmRuntime.Close(ctx)

		// get the function
		wasmModuleHandleFunction := wasmModule.ExportedFunction("callHandle")

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
		// Finally, we get the message "👋 hello <name>" printed. This shows how to
		// read-back something allocated by TinyGo.
		handleResultArray, err := wasmModuleHandleFunction.Call(ctx, stringParameterPtrPosition, stringParameterLength)
		if err != nil {
			log.Panicln(err)
		}
		// Note: This pointer is still owned by TinyGo, so don't try to free it!
		handleReturnPtrPos, handleReturnSize := helpers.GetPackedPtrPositionAndSize(handleResultArray)

		// The pointer is a linear memory offset, which is where we write the name.
		if bytes, ok := wasmModule.Memory().Read(ctx, handleReturnPtrPos, handleReturnSize); !ok {
			log.Panicf("Memory.Read(%d, %d) out of range of memory size %d",
				handleReturnPtrPos, handleReturnSize, wasmModule.Memory().Size(ctx))
      jsonResult.Value = ""
      jsonResult.Error = "out of range of memory size"
      return c.JSON(http.StatusConflict, jsonResult)
		} else {
      jsonResult.Value = string(bytes)
      jsonResult.Error = ""
			return c.JSON(http.StatusOK, jsonResult)
		}

  })
  //https://echo.labstack.com/guide/customization/
  e.HideBanner = true
  e.Start(":" + httpPort)

  //e.Logger.Info(e.Start(":" + httpPort))
	//e.Logger.Fatal(e.Start(":" + httpPort))

}
