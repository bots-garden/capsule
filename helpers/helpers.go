package helpers

import (
	"unsafe"
)

//export hostLogString
func hostLogString(ptr uint32, size uint32)

func Log(message string) {
	ptr, size := stringToPtr(message)
	hostLogString(ptr, size)
}

// stringToPtr returns a pointer and size pair for the given string in a way
// compatible with WebAssembly numeric types.
func stringToPtr(s string) (uint32, uint32) {
	buf := []byte(s)
	ptr := &buf[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return uint32(unsafePtr), uint32(len(buf))
}
