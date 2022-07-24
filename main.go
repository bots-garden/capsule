package main

import (
	"context"
	"fmt"
	"log"
	"os"

	helpers "github.com/bots-garden/capsule/helpers/tools"
	host_functions "github.com/bots-garden/capsule/host_functions"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/wasi_snapshot_preview1"
)

func main() {
	// Choose the context to use for function calls.
	ctx := context.Background()

	// Create a new WebAssembly Runtime.
	wasmRuntime := wazero.NewRuntimeWithConfig(wazero.NewRuntimeConfig().WithWasmCore2())
	defer wasmRuntime.Close(ctx) // This closes everything this Runtime created.

	_, errEnv := wasmRuntime.NewModuleBuilder("env").
		ExportFunction("hostLogString", host_functions.LogString).
		Instantiate(ctx, wasmRuntime)

	if errEnv != nil {
		log.Panicln("🔴 Error with env module and host function(s):", errEnv)
	}

	_, errInstantiate := wasi_snapshot_preview1.Instantiate(ctx, wasmRuntime)
	if errInstantiate != nil {
		log.Panicln("🔴 Error with Instantiate:", errInstantiate)
	}

	// Load then Instantiate a WebAssembly module
	helloWasm, errLoadWasmModule := os.ReadFile("./wasm_modules/03-string-as-param/hello.wasm")
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

	// ======================================================
	// 2️⃣ Get a string from wasm
	// ======================================================

	helloWorldwasmModuleFunction := mod.ExportedFunction("helloWorld")

	resultArray, errCallFunction := helloWorldwasmModuleFunction.Call(ctx)
	if errCallFunction != nil {
		log.Panicln("🔴 Error while calling the function ", errCallFunction)
	}
	// Note: This pointer is still owned by TinyGo, so don't try to free it!
	helloWorldPtr, helloWorldSize := helpers.GetPackedPtrPositionAndSize(resultArray)

	// The pointer is a linear memory offset, which is where we write the name.
	if bytes, ok := mod.Memory().Read(ctx, helloWorldPtr, helloWorldSize); !ok {
		log.Panicf("🟥 Memory.Read(%d, %d) out of range of memory size %d",
			helloWorldPtr, helloWorldSize, mod.Memory().Size(ctx))
	} else {
		//fmt.Println(bytes)
		fmt.Println("😃 the string message is:", string(bytes))
	}

	// ======================================================
	// 3️⃣ Pass a string param and get a string result
	// ======================================================

	// Parameter "setup"
	name := "Bob Morane 🎉"
	nameSize := uint64(len(name))

	// get the function
	sayHelloWasmModuleFunction := mod.ExportedFunction("sayHello")

	// These are undocumented, but exported. See tinygo-org/tinygo#2788
	malloc := mod.ExportedFunction("malloc")
	free := mod.ExportedFunction("free")

	// Instead of an arbitrary memory offset, use TinyGo's allocator.
	// 🖐 Notice there is nothing string-specific in this allocation function.
	// The same function could be used to pass binary serialized data to Wasm.
	results, err := malloc.Call(ctx, nameSize)
	if err != nil {
		log.Panicln(err)
	}
	namePtrPosition := results[0]
	// This pointer is managed by TinyGo, but TinyGo is unaware of external usage.
	// So, we have to free it when finished
	defer free.Call(ctx, namePtrPosition)

	// The pointer is a linear memory offset, which is where we write the name.
	if !mod.Memory().Write(ctx, uint32(namePtrPosition), []byte(name)) {
		log.Panicf("🟥 Memory.Write(%d, %d) out of range of memory size %d",
			namePtrPosition, nameSize, mod.Memory().Size(ctx))
	}
	// Finally, we get the message "👋 hello <name>" printed. This shows how to
	// read-back something allocated by TinyGo.
	sayHelloResultArray, err := sayHelloWasmModuleFunction.Call(ctx, namePtrPosition, nameSize)
	if err != nil {
		log.Panicln(err)
	}
	// Note: This pointer is still owned by TinyGo, so don't try to free it!
	sayHelloPtrPos, sayHelloSize := helpers.GetPackedPtrPositionAndSize(sayHelloResultArray)

	// The pointer is a linear memory offset, which is where we write the name.
	if bytes, ok := mod.Memory().Read(ctx, sayHelloPtrPos, sayHelloSize); !ok {
		log.Panicf("Memory.Read(%d, %d) out of range of memory size %d",
		sayHelloPtrPos, sayHelloSize, mod.Memory().Size(ctx))
	} else {
		fmt.Println("👋👋👋:", string(bytes)) // the result
	}
	

}
