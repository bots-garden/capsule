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

    //argsWithProg := os.Args
    wasmModuleFilePath := os.Args[1:][0]

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

	// 📂 Load from file and then Instantiate a WebAssembly module
	wasmFile, errLoadWasmFile := os.ReadFile(wasmModuleFilePath)

	if errLoadWasmFile != nil {
		log.Panicln("🔴 Error while loading the wasm file", errLoadWasmFile)
	}

	// 🥚 Instantiate the wasm module (from the wasm file)
	wasmModule, errInstanceWasmModule := wasmRuntime.InstantiateModuleFromBinary(ctx, wasmFile)
	if errInstanceWasmModule != nil {
		log.Panicln("🔴 Error while creating module instance ", errInstanceWasmModule)
	}

	// 1️⃣ Get references to WebAssembly function: "add"
	addWasmModuleFunction := wasmModule.ExportedFunction("add")

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

	helloWorldwasmModuleFunction := wasmModule.ExportedFunction("helloWorld")

	resultArray, errCallFunction := helloWorldwasmModuleFunction.Call(ctx)
	if errCallFunction != nil {
		log.Panicln("🔴 Error while calling the function ", errCallFunction)
	}
	// Note: This pointer is still owned by TinyGo, so don't try to free it!
	helloWorldPtr, helloWorldSize := helpers.GetPackedPtrPositionAndSize(resultArray)

	// The pointer is a linear memory offset, which is where we write the name.
	if bytes, ok := wasmModule.Memory().Read(ctx, helloWorldPtr, helloWorldSize); !ok {
		log.Panicf("🟥 Memory.Read(%d, %d) out of range of memory size %d",
			helloWorldPtr, helloWorldSize, wasmModule.Memory().Size(ctx))
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
	sayHelloWasmModuleFunction := wasmModule.ExportedFunction("sayHello")

	// These are undocumented, but exported. See tinygo-org/tinygo#2788
	malloc := wasmModule.ExportedFunction("malloc")
	free := wasmModule.ExportedFunction("free")

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
	if !wasmModule.Memory().Write(ctx, uint32(namePtrPosition), []byte(name)) {
		log.Panicf("🟥 Memory.Write(%d, %d) out of range of memory size %d",
			namePtrPosition, nameSize, wasmModule.Memory().Size(ctx))
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
	if bytes, ok := wasmModule.Memory().Read(ctx, sayHelloPtrPos, sayHelloSize); !ok {
		log.Panicf("Memory.Read(%d, %d) out of range of memory size %d",
		sayHelloPtrPos, sayHelloSize, wasmModule.Memory().Size(ctx))
	} else {
		fmt.Println("👋👋👋:", string(bytes)) // the result
	}
	

}
