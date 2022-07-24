package helpers

import (
	"unsafe"
)

//export hostLogString
func hostLogString(ptr, size uint32)

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
