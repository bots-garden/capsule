## Call a function with wasm types

```golang
// 1️⃣ Get references to WebAssembly function: "add"
addWasmModuleFunction := wasmModule.ExportedFunction("add")

// Now, we can call "add", which reads the string we wrote to memory!
// result []uint64
result, errCallFunction := addWasmModuleFunction.Call(ctx, 20, 22)
if errCallFunction != nil {
log.Panicln("🔴 Error while calling the function ", errCallFunction)
}

fmt.Println("result:", result[0])
```

## Call a function without a parameter

> returning a string

```golang
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
```

## Call a function with a string parameter

> returning a string

```golang
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
```





package hostfunctions

import (
	"context"
	"fmt"

	//"github.com/go-resty/resty/v2"
	"github.com/tetratelabs/wazero/api"
)

func Http(ctx context.Context, module api.Module,
	urlOffset, urlByteCount, methodOffSet, methodByteCount, headersOffSet, headersByteCount, bodyOffSet, bodyByteCount,
	retBuffPtrPos, retBuffSize uint32) {

    //=========================================================
	// Read arguments values of the function call
    //=========================================================

	// get url string from the wasm module function (from memory)
	urlStr := ReadStringFromMemory(ctx, module, urlOffset, urlByteCount)

	// get method string from the wasm module function (from memory)
	methodStr := ReadStringFromMemory(ctx, module, methodOffSet, methodByteCount)

	// get headers string from the wasm module function (from memory)
	// 🖐 headers => Accept:application/json|Content-Type: text/html; charset=UTF-8
	headersStr := ReadStringFromMemory(ctx, module, headersOffSet, headersByteCount)

    //TODO: choose another separator: °
	headersSlice := CreateSliceFromString(headersStr, "|")

	fmt.Println(headersSlice)

    headersMap := CreateMapFromSlice(headersSlice, ":")

    fmt.Println(headersMap)
    fmt.Println(headersMap["Accept"])
    fmt.Println(headersMap["Content-Type"])

	// get body string from the wasm module function (from memory)
	bodyStr := ReadStringFromMemory(ctx, module, bodyOffSet, bodyByteCount)

	// 👋 Implementation: Start
	var stringMessageFromHost = ""
	//client := resty.New()

    // set headers
    /*
    for key, value := range headersMap {
        fmt.Println("-", key, value)
    }
    */

	switch what := methodStr; what {
	case "GET":

        /*
		resp, err := client.R().EnableTrace().Get(urlStr)

        if err != nil {
            stringMessageFromHost = CreateStringError(err.Error(), 666)
        } else {
            stringMessageFromHost = resp.String()
        }
        */

		stringMessageFromHost = "🌍 (GET)http: " + urlStr + " method: " + methodStr + " headers: " + headersStr + " body: " + bodyStr


	case "POST":
		stringMessageFromHost = "🌍 (POST)http: " + urlStr + " method: " + methodStr + " headers: " + headersStr + " body: " + bodyStr

	default:
		stringMessageFromHost = CreateStringError("🔴 not implemented: 🚧 wip", 999)
	}
	// 👋 Implementation: End

	// write the new string stringMessageFromHost to the "shared memory"
	// (host write string result of the funcyion to memory)
	WriteStringToMemory(stringMessageFromHost, ctx, module, retBuffPtrPos, retBuffSize)

}
