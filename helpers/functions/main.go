package helpers

import (
	"reflect"
	"unsafe"
)

//export allocateBuffer
func allocateBuffer(size uint32) *byte {
	// Allocate the in-Wasm memory region and returns its pointer to hosts.
	// The region is supposed to store random strings generated in hosts,
	// meaning that this is called "inside" of host_get_string.
	buf := make([]byte, size)
	return &buf[0]
}

//export hostLogString
func hostLogString(ptrPos, size uint32)

//export hostGetHostInformation
func hostGetHostInformation(retBuffPtrPos **byte, retBuffSize *int)

//export hostPing
func hostPing(ptrPos uint32, size uint32, retBuffPtrPos **byte, retBuffSize *int)

/*
Call host function: hostPing
Pass a string as parameter
Get a string from the host
*/
func Ping(message string) string {

	strPtrPos, size := GetStringPtrPositionAndSize(message)

	var bufPtr *byte
	var bufSize int

	hostPing(strPtrPos, size, &bufPtr, &bufSize)

	return *(*string)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(bufPtr)),
		Len:  uintptr(bufSize),
		Cap:  uintptr(bufSize),
	}))

}

/*
Call host function: hostGetHostInformation
Get a string with the information about the host
*/
func GetHostInformation() string {
	var bufPtr *byte
	var bufSize int
	hostGetHostInformation(&bufPtr, &bufSize)

	return *(*string)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(bufPtr)),
		Len:  uintptr(bufSize),
		Cap:  uintptr(bufSize),
	}))
}

/*
Call host function: hostLogString.
Print a string
*/
func Log(message string) {
	ptr, size := GetStringPtrPositionAndSize(message)
	hostLogString(ptr, size)
}

/*
It returns a pointer and size pair for the given string
in a way compatible with WebAssembly numeric types.
*/
func GetStringPtrPositionAndSize(text string) (stringPointerPosition uint32, stringSize uint32) {
	buf := []byte(text)
	ptr := &buf[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return uint32(unsafePtr), uint32(len(buf))
}

/*
Pack 2 uint32 values to an unique uint64 value.
ex: position of a string pointer and the size(length) of the string
*/
func PackPtrPositionAndSize(ptrPos uint32, size uint32) (packedValue uint64) {
	return (uint64(ptrPos) << uint64(32)) | uint64(size)
}

/*
GetStringParam returns a string
from WebAssembly compatible numeric types representing its pointer and length.
*/
func GetStringParam(ptrPos uint32, size uint32) string {
	// Get a slice view of the underlying bytes in the stream. We use SliceHeader, not StringHeader
	// as it allows us to fix the capacity to what was allocated.
	return *(*string)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(ptrPos),
		Len:  uintptr(size), // Tinygo requires these as uintptrs even if they are int fields.
		Cap:  uintptr(size), // ^^ See https://github.com/tinygo-org/tinygo/issues/1284
	}))
}

// TODO: Try to do the same thing with alloc 🤔
// TODO: see https://www.wasm.builders/k33g_org/an-essay-on-the-bi-directional-exchange-of-strings-between-the-wasm-module-with-tinygo-and-nodejs-with-wasi-support-3i9h


var handleFunction func(string) string

func SetHandle(function func(string) string) {
	handleFunction = function
}

/*

*/
// TODO add detailed comments
//export callHandle
func callHandle(strPtrPos, size uint32) (strPtrPosSize uint64) {
	stringParameter := GetStringParam(strPtrPos, size)

	stringReturnByHandleFunction := handleFunction(stringParameter)

	pos, length := GetStringPtrPositionAndSize(stringReturnByHandleFunction)

	return PackPtrPositionAndSize(pos, length)
}


/* Function Samples
//export helloWorld
func helloWorld() (strPtrPosSize uint64) {
	strPtrPos, size := helpers.GetStringPtrPositionAndSize("👋 hello world, I'm very happy to meet you, I love what you are doing my friend")
	return helpers.PackPtrPositionAndSize(strPtrPos, size)
}

//export sayHello
func sayHello(strPtrPos, size uint32) (strPtrPosSize uint64) {
	name := helpers.GetStringParam(strPtrPos, size)
	pos, length := helpers.GetStringPtrPositionAndSize("👋 hello " + name)

	return helpers.PackPtrPositionAndSize(pos, length)
}
*/
