package main

import (
	"context"
	"fmt"
	helpers "github.com/bots-garden/capsule/helpers/tools"
	"log"
	"os"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/wasi_snapshot_preview1"
)

func main() {
	// Choose the context to use for function calls.
	ctx := context.Background()

	// Create a new WebAssembly Runtime.
	wasmRuntime := wazero.NewRuntimeWithConfig(wazero.NewRuntimeConfig().WithWasmCore2())
	defer wasmRuntime.Close(ctx) // This closes everything this Runtime created.

	_, errEnv := wasmRuntime.NewModuleBuilder("env").
		ExportFunction("hostLogString", logString).
		Instantiate(ctx, wasmRuntime)

	if errEnv != nil {
		log.Panicln("🔴 Error with env module and host function(s):", errEnv)
	}

	_, errInstantiate := wasi_snapshot_preview1.Instantiate(ctx, wasmRuntime)
	if errInstantiate != nil {
		log.Panicln("🔴 Error with Instantiate:", errInstantiate)
	}

	// Load then Instantiate a WebAssembly module
	helloWasm, errLoadWasmModule := os.ReadFile("./wasm_modules/02-return-string/hello.wasm")
	if errLoadWasmModule != nil {
		log.Panicln("🔴 Error while loading the wasm module", errLoadWasmModule)
	}

	mod, errInstanceWasmModule := wasmRuntime.InstantiateModuleFromBinary(ctx, helloWasm)
	if errInstanceWasmModule != nil {
		log.Panicln("🔴 Error while creating module instance ", errInstanceWasmModule)
	}

	// 1️⃣ Get references to WebAssembly function: "add"
	addWasmModuleFunction := mod.ExportedFunction("add")

	// Now, we can call "add", which reads the string we wrote to memory!
	// result []uint64
	result, errCallFunction := addWasmModuleFunction.Call(ctx, 20, 22)
	if errCallFunction != nil {
		log.Panicln("🔴 Error while calling the function ", errCallFunction)
	}

	fmt.Println("result:", result[0])

	// 2️⃣ Get a string from wasm
	helloWorldwasmModuleFunction := mod.ExportedFunction("helloWorld")

	resultArray, errCallFunction := helloWorldwasmModuleFunction.Call(ctx)
	if errCallFunction != nil {
		log.Panicln("🔴 Error while calling the function ", errCallFunction)
	}
	// Note: This pointer is still owned by TinyGo, so don't try to free it!
	helloWorldPtr, helloWorldSize := helpers.PtrSizePairFromArray(resultArray)

	// The pointer is a linear memory offset, which is where we write the name.
	if bytes, ok := mod.Memory().Read(ctx, helloWorldPtr, helloWorldSize); !ok {
		log.Panicf("🟥 Memory.Read(%d, %d) out of range of memory size %d",
			helloWorldPtr, helloWorldSize, mod.Memory().Size(ctx))
	} else {
		//fmt.Println(bytes)
		fmt.Println("😃 the string message is:", string(bytes))
	}

}

// host wasm_modules
func logString(ctx context.Context, module api.Module, offset, byteCount uint32) {
	buf, ok := module.Memory().Read(ctx, offset, byteCount)
	if !ok {
		log.Panicf("🟥 Memory.Read(%d, %d) out of range", offset, byteCount)
	}
	fmt.Println(string(buf))
}
