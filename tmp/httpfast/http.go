package capsulehttpfast

import (
	"context"
	"log"
	"net/http"

	helpers "github.com/bots-garden/capsule/helpers/tools"
	"github.com/bots-garden/capsule/host_functions"
	"github.com/gin-gonic/gin"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/wasi_snapshot_preview1"
)

type JsonParameter struct {
	Message string `json:"message"` // change the name ? 🤔
}

// TODO add output

// 🚧 this is a work in progress
func callWasmFunction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

//TODO handle all errors
//TODO handle errors from the wasm module too
/*
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json' \
  -d '{"message": "Golang 💚 wasm"}'
*/

func callPostWasmFunctionHandler(wasmFile []byte) gin.HandlerFunc {

	// Choose the context to use for function calls.
	ctx := context.Background()

	// Create a new WebAssembly Runtime.
	wasmRuntime := wazero.NewRuntimeWithConfig(wazero.NewRuntimeConfig().WithWasmCore2())
	defer wasmRuntime.Close(ctx) // This closes everything this Runtime created.

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

	fn := func(c *gin.Context) {

		// 🥚 Instantiate the wasm module (from the wasm file)
		wasmModule, errInstanceWasmModule := wasmRuntime.InstantiateModuleFromBinary(ctx, wasmFile)
		if errInstanceWasmModule != nil {
			log.Panicln("🔴 Error while creating module instance:", errInstanceWasmModule)
		}

		var jsonParameter JsonParameter

		// Call BindJSON to bind the received JSON to
		// jsonParameter.
		if err := c.BindJSON(&jsonParameter); err != nil {
			// TODO: DO SOMETHING WITH THE ERROR
			return
		}
		// Parameter "setup"
		stringParameterLength := uint64(len(jsonParameter.Message))

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
		if !wasmModule.Memory().Write(ctx, uint32(stringParameterPtrPosition), []byte(jsonParameter.Message)) {
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
		} else {
			//fmt.Println("🤖:", string(bytes)) // the result
			c.JSON(http.StatusOK, gin.H{"value": string(bytes)})
			//c.String(http.StatusOK, `{"value":"`+string(bytes)+`"}`)
		}

	}

	return fn
}

func Serve(httpPort string, wasmFile []byte) {

	r := gin.Default()
	r.GET("/", callWasmFunction)
	r.POST("/", callPostWasmFunctionHandler(wasmFile))
	r.Run(":" + httpPort)
}

/*
see https://github.com/bots-garden/procyon/blob/main/procyon-reverse-proxy/main.go
*/
