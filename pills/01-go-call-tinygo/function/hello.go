package main

import (
	"reflect"
	"strconv"
	"unsafe"
)

// main is required for TinyGo to compile to Wasm.
func main() {}

//export hostLogString
//go:linkname hostLogString
func hostLogString(ptr uint32, size uint32)

//export hostLogUint32
//go:linkname hostLogUint32
func hostLogUint32(value uint32)

//export add
func add(x uint32, y uint32) uint32 {
	// 🖐 a wasm module cannot print something
	//fmt.Println(x,y)
	res := x + y

	hostLogUint32(res)

	ptr, size := stringToPtr("from wasm: " + strconv.FormatUint(uint64(res), 10))
	hostLogString(ptr, size)

	return res
}

// 🖐️ returns a pointer/size pair packed into a uint64.
// Note: This uses a uint64 instead of two result values for compatibility with
// WebAssembly 1.0.
// https://stackoverflow.com/questions/5801008/go-and-operators
// https://stackoverflow.com/questions/41790574/bitmask-multiple-values-in-int64

//export helloWorld
func helloWorld() (ptrAndSize uint64) {
	ptr, size := stringToPtr("👋 hello world, I'm very happy to meet you, I love what you are doing my friend")
	return (uint64(ptr) << uint64(32)) | uint64(size)
}

//export sayHello
func sayHello(ptr, size uint32) (ptrAndSize uint64) {
	// get the parameter
	name := ptrToString(ptr, size)

	ptr, size = stringToPtr("👋 hello " + name + " 😃")
	return (uint64(ptr) << uint64(32)) | uint64(size)
}

// stringToPtr returns a pointer and size pair for the given string in a way
// compatible with WebAssembly numeric types.
func stringToPtr(s string) (uint32, uint32) {
	buf := []byte(s)
	ptr := &buf[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return uint32(unsafePtr), uint32(len(buf))
}

// ptrToString returns a string from WebAssembly compatible numeric types
// representing its pointer and length.
func ptrToString(ptr uint32, size uint32) string {
	// Get a slice view of the underlying bytes in the stream. We use SliceHeader, not StringHeader
	// as it allows us to fix the capacity to what was allocated.
	return *(*string)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(ptr),
		Len:  uintptr(size), // Tinygo requires these as uintptrs even if they are int fields.
		Cap:  uintptr(size), // ^^ See https://github.com/tinygo-org/tinygo/issues/1284
	}))
}
