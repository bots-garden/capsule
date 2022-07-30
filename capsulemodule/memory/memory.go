// host functions
package memory

import (
	"reflect"
	"unsafe"
)

// 🤔 TODO: is it used ?
//export allocateBuffer
//go:linkname allocateBuffer
func allocateBuffer(size uint32) *byte {
	// Allocate the in-Wasm memory region and returns its pointer to hosts.
	// The region is supposed to store random strings generated in hosts,
	// meaning that this is called "inside" of host_get_string.
	buf := make([]byte, size)
	return &buf[0]
}

/*
It returns a pointer and size pair for the given string
in a way compatible with WebAssembly numeric types.
*/
func GetStringPtrPositionAndSize(text string) (stringPointerPosition uint32, stringSize uint32) {
	if text == "" {
		text = "empty"
		//!: strange trick to avoid error when passing empty body string to hf.Http:
		//!: hf.Http("https://httpbin.org/get", "GET", headers, "")
	}
	buff := []byte(text)
	ptr := &buff[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return uint32(unsafePtr), uint32(len(buff))
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

Uses by callHandle (👀 see handle.go)
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

func GetStringResult(buffPtr *byte, buffSize int) string {
	result := *(*string)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(buffPtr)),
		Len:  uintptr(buffSize),
		Cap:  uintptr(buffSize),
	}))
	return result
}
