package helpers

import (
	"unsafe"
)

//export hostLogString
func hostLogString(ptr uint32, size uint32)

func Log(message string) {
	ptr, size := StringToPtr(message)
	hostLogString(ptr, size)
}

// StringToPtr returns a pointer and size pair for the given string in a way
// compatible with WebAssembly numeric types.
func StringToPtr(s string) (uint32, uint32) {
	buf := []byte(s)
	ptr := &buf[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return uint32(unsafePtr), uint32(len(buf))
}

func PtrSizePair(ptr uint32, size uint32) uint64 {
	return (uint64(ptr) << uint64(32)) | uint64(size)
}
