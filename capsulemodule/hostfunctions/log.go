package hostfunctions

import (
	_ "unsafe"
)

//export hostLogString
//go:linkname hostLogString
func hostLogString(ptrPos, size uint32)

// Log : call host function: hostLogString
// Print a string
func Log(message string) {
	ptr, size := getStringPtrPositionAndSize(message)
	hostLogString(ptr, size)
}
