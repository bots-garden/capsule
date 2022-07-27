// host functions
package hf

//export hostLogString
//go:linkname hostLogString
func hostLogString(ptrPos, size uint32)

/*
Call host function: hostLogString.
Print a string
*/
func Log(message string) {
	ptr, size := GetStringPtrPositionAndSize(message)
	hostLogString(ptr, size)
}
