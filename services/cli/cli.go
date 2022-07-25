package capsulecli


import (
	"context"
	"fmt"
	"log"

	helpers "github.com/bots-garden/capsule/helpers/tools"
	"github.com/bots-garden/capsule/host_functions"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/wasi_snapshot_preview1"
)

// if it works, try with a pool of wasmRuntime or WasmModule
func getWasmRuntime(ctx context.Context) wazero.Runtime {
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


func getWasmRuntimeAndModuleInstances(wasmFile []byte) (wazero.Runtime, api.Module, context.Context) {
  // Choose the context to use for function calls.
	ctx := context.Background()

	wasmRuntime := getWasmRuntime(ctx)
	//defer wasmRuntime.Close(ctx) // This closes everything this Runtime created.

	// 🥚 Instantiate the wasm module (from the wasm file)
	wasmModule, errInstanceWasmModule := wasmRuntime.InstantiateModuleFromBinary(ctx, wasmFile)
	if errInstanceWasmModule != nil {
		log.Panicln("🔴 Error while creating module instance:", errInstanceWasmModule)
	}
	return wasmRuntime, wasmModule, ctx
}

// Pass a string param and get a string result
func Execute(stringParameter string, wasmFile []byte) {

	// Choose the context to use for function calls.
	//ctx := context.Background()

  // 👋 get Wasm Module Instance (and Wasm runtime)
  wasmRuntime, wasmModule, ctx := getWasmRuntimeAndModuleInstances(wasmFile)
  // defer must always be in the main code (to avoid go routine panic)
  defer wasmRuntime.Close(ctx)


	// Parameter "setup" / stringParameter comes from argument
	stringParameterLength := uint64(len(stringParameter))

	// get the callHandle function
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
