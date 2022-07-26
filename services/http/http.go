package capsulehttp

import (
	"log"

	"github.com/gin-gonic/gin"
	"net/http"

	helpers "github.com/bots-garden/capsule/helpers/tools"
	capsulecommon "github.com/bots-garden/capsule/services/common"
)

type JsonParameter struct {
	Message string `json:"message"` // change the name ? 🤔
}

/*
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json' \
  -d '{"message": "Golang 💚 wasm"}'
*/

func callPostWasmFunctionHandler(wasmFile []byte) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		var jsonParameter JsonParameter
		// Call BindJSON to bind the received JSON to
		// jsonParameter.
		// TODO: handle json errors
		if err := c.BindJSON(&jsonParameter); err != nil {
			return
		}

		// Parameter "setup"
		stringParameterLength := uint64(len(jsonParameter.Message))
		stringParameter := jsonParameter.Message

		wasmRuntime, wasmModule, ctx := capsulecommon.CreateWasmRuntimeAndModuleInstances(wasmFile)
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

		} else {
			c.JSON(http.StatusOK, gin.H{"value": string(bytes)})
		}

	}

	return fn
}

func Serve(httpPort string, wasmFile []byte) {

	r := gin.Default()
	r.POST("/", callPostWasmFunctionHandler(wasmFile))
	r.Run(":" + httpPort)
}
