package hostfunctions

import (
	_ "unsafe"
)

//export hostLogString
//go:linkname hostLogString
func hostLogString(position, length uint32) uint32

// Log : call host function: hostLogString
// Print a string
func Log(message string) {
	position, length := getStringPtrPositionAndSize(message)
	hostLogString(position, length)
}
