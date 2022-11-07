package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

func main() {
	// Choose the context to use for function calls.
	ctx := context.Background()

	// Create a new WebAssembly Runtime.
	wasmRuntime := wazero.NewRuntime(ctx)
	defer wasmRuntime.Close(ctx) // This closes everything this Runtime created.

	// host functions

	doSomethingWithNumbers := api.GoFunc(func(ctx context.Context, params []uint64) []uint64 {
		fmt.Println("🎃", params)
		return []uint64{42}
	})

	logStringPlus := api.GoModuleFunc(func(ctx context.Context, module api.Module, params []uint64) []uint64 {
		fmt.Println("🌺", params)
		fmt.Println("🖐 position:", params[0])
		fmt.Println("🖐 length:", params[1])

		position := uint32(params[0])
		length := uint32(params[1])

		buffer, ok := module.Memory().Read(ctx, position, length)
		if !ok {
			log.Panicf("🟥 Memory.Read(%d, %d) out of range", position, length)
		}
		fmt.Println("🐼:", string(buffer))

		return []uint64{0}
	})

	helloWorld := api.GoModuleFunc(func(ctx context.Context, module api.Module, params []uint64) []uint64 {
		fmt.Println("🌸", params)
		fmt.Println("🖐 position:", params[0])
		fmt.Println("🖐 length:", params[1])

		position := uint32(params[0])
		length := uint32(params[1])

		// Read the parameters (position and length)
		buffer, ok := module.Memory().Read(ctx, position, length)
		if !ok {
			log.Panicf("🟥 Memory.Read(%d, %d) out of range", position, length)
		}
		fmt.Println("🐼🎃:", string(buffer))

		return []uint64{42, 53} // 🤔
	})

	_, errEnv := wasmRuntime.NewHostModuleBuilder("env").
		NewFunctionBuilder().WithFunc(func(value uint32) {
		fmt.Println("🤖:", value)
	}).Export("hostLogUint32").
		NewFunctionBuilder().WithFunc(logString).Export("hostLogString").
		NewFunctionBuilder().WithGoFunction(doSomethingWithNumbers, []api.ValueType{api.ValueTypeI32}, []api.ValueType{api.ValueTypeI32}).Export("hostDoSomethingWithNumbers").
		NewFunctionBuilder().WithGoModuleFunction(logStringPlus, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{api.ValueTypeI32}).Export("hostLogStringPlus").
		NewFunctionBuilder().WithGoModuleFunction(helloWorld, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{api.ValueTypeI32}).Export("hostHelloWorld").
		Instantiate(ctx, wasmRuntime)

	//wasmRuntime.NewHostModuleBuilder().NewFunctionBuilder().WithGoFunction()

	/*
	   builder.WithGoFunction(api.GoFunc(func(ctx context.Context, params []uint64) []uint64 {
	   	x, y := uint32(params[0]), uint32(params[1])
	   	sum := x + y
	   	return []uint64{sum}
	   }, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{api.ValueTypeI32})
	*/

	/*
	   func logString(ctx context.Context, module api.Module, offset, byteCount uint32) {
	   	buf, ok := module.Memory().Read(ctx, offset, byteCount)
	   	if !ok {
	   		log.Panicf("🟥 Memory.Read(%d, %d) out of range", offset, byteCount)
	   	}
	   	fmt.Println("👽:", string(buf))
	   }
	*/

	if errEnv != nil {
		log.Panicln("🔴 Error with env module and host function(s):", errEnv)
	}

	_, errInstantiate := wasi_snapshot_preview1.Instantiate(ctx, wasmRuntime)
	if errInstantiate != nil {
		log.Panicln("🔴 Error with Instantiate:", errInstantiate)
	}

	// Load then Instantiate a WebAssembly module
	helloWasm, errLoadWasmModule := os.ReadFile("./function/hello.wasm")
	if errLoadWasmModule != nil {
		log.Panicln("🔴 Error while loading the wasm module", errLoadWasmModule)
	}

	mod, errInstanceWasmModule := wasmRuntime.InstantiateModuleFromBinary(ctx, helloWasm)
	if errInstanceWasmModule != nil {
		log.Panicln("🔴 Error while creating module instance ", errInstanceWasmModule)
	}

	// Get references to WebAssembly function: "add"
	addWasmModuleFunction := mod.ExportedFunction("add")

	// Now, we can call "add", which reads the string we wrote to memory!
	// result []uint64
	result, errCallFunction := addWasmModuleFunction.Call(ctx, 20, 22)
	if errCallFunction != nil {
		log.Panicln("🔴 Error while calling the function ", errCallFunction)
	}

	fmt.Println("result:", result[0])

	// ======================================================
	// Get a string from wasm
	// ======================================================
	helloWorldWasmModuleFunction := mod.ExportedFunction("helloWorld")

	ptrSize, errCallFunction := helloWorldWasmModuleFunction.Call(ctx)
	if errCallFunction != nil {
		log.Panicln("🔴 Error while calling the function ", errCallFunction)
	}
	// Note: This pointer is still owned by TinyGo, so don't try to free it!
	helloWorldPtr := uint32(ptrSize[0] >> 32)
	helloWorldSize := uint32(ptrSize[0])

	// The pointer is a linear memory offset, which is where we write the name.
	if bytes, ok := mod.Memory().Read(ctx, helloWorldPtr, helloWorldSize); !ok {
		log.Panicf("🟥 Memory.Read(%d, %d) out of range of memory size %d",
			helloWorldPtr, helloWorldSize, mod.Memory().Size(ctx))
	} else {
		fmt.Println("😃 the string message is:", string(bytes))
	}

	// ======================================================
	// Pass a string param and get a string result
	// ======================================================
	// Let's use the argument to this main function in Wasm.
	name := "Bob Morane"
	nameSize := uint64(len(name))
	// Get references to WebAssembly functions we'll use in this example.

	sayHelloWasmModuleFunction := mod.ExportedFunction("sayHello")

	// These are undocumented, but exported. See tinygo-org/tinygo#2788
	malloc := mod.ExportedFunction("malloc")
	free := mod.ExportedFunction("free")

	// Instead of an arbitrary memory offset, use TinyGo's allocator. Notice
	// there is nothing string-specific in this allocation function. The same
	// function could be used to pass binary serialized data to Wasm.
	results, err := malloc.Call(ctx, nameSize)
	if err != nil {
		log.Panicln(err)
	}
	namePtr := results[0]
	// This pointer is managed by TinyGo, but TinyGo is unaware of external usage.
	// So, we have to free it when finished
	defer free.Call(ctx, namePtr)

	// The pointer is a linear memory offset, which is where we write the name.
	if !mod.Memory().Write(ctx, uint32(namePtr), []byte(name)) {
		log.Panicf("🟥 Memory.Write(%d, %d) out of range of memory size %d",
			namePtr, nameSize, mod.Memory().Size(ctx))
	}

	// Finally, we get the greeting message "greet" printed. This shows how to
	// read-back something allocated by TinyGo.
	sayHelloPtrSize, err := sayHelloWasmModuleFunction.Call(ctx, namePtr, nameSize)
	if err != nil {
		log.Panicln(err)
	}
	// Note: This pointer is still owned by TinyGo, so don't try to free it!
	sayHelloPtr := uint32(sayHelloPtrSize[0] >> 32)
	sayHelloSize := uint32(sayHelloPtrSize[0])
	// The pointer is a linear memory offset, which is where we write the name.
	if bytes, ok := mod.Memory().Read(ctx, sayHelloPtr, sayHelloSize); !ok {
		log.Panicf("Memory.Read(%d, %d) out of range of memory size %d",
			sayHelloPtr, sayHelloSize, mod.Memory().Size(ctx))
	} else {
		fmt.Println("👋 saying hello :", string(bytes))
	}

}

func logString(ctx context.Context, module api.Module, offset, byteCount uint32) {
	buf, ok := module.Memory().Read(ctx, offset, byteCount)
	if !ok {
		log.Panicf("🟥 Memory.Read(%d, %d) out of range", offset, byteCount)
	}
	fmt.Println("👽:", string(buf))
}
