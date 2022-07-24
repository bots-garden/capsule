package capsulecli

import (
	"context"
	"fmt"
	"log"

	helpers "github.com/bots-garden/capsule/helpers/tools"
	"github.com/tetratelabs/wazero/api"
)

// Pass a string param and get a string result
func Execute(stringParameter string, wasmModule api.Module, ctx context.Context) {
	// Parameter "setup"
	//stringParameter := "Bob Morane 🎉"
	stringParameterLength := uint64(len(stringParameter))

	// get the function
	wasmModuleHandleFunction := wasmModule.ExportedFunction("callHandle")

	// These are undocumented, but exported. See tinygo-org/tinygo#2788
	malloc := wasmModule.ExportedFunction("malloc")
	free := wasmModule.ExportedFunction("free")

	// Instead of an arbitrary memory offset, use TinyGo's allocator.
	// 🖐 Notice there is nothing string-specific in this allocation function.
	// The same function could be used to pass binary serialized data to Wasm.
	results, err := malloc.Call(ctx, stringParameterLength)
	if err != nil {
		log.Panicln(err)
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
		fmt.Println("🤖:", string(bytes)) // the result
	}

}
