
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
