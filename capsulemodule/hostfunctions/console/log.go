package hf_console

import (
	"github.com/bots-garden/capsule/capsulemodule/memory"
	_ "unsafe"
)

//export hostLogString
//go:linkname hostLogString
func hostLogString(ptrPos, size uint32)

// Log : call host function: hostLogString
// Print a string
func Log(message string) {
	ptr, size := memory.GetStringPtrPositionAndSize(message)
	hostLogString(ptr, size)
}
