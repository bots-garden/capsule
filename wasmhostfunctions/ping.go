// host functions
package hf

import _ "unsafe"

//export hostPing
//go:linkname hostPing
func hostPing(ptrPos uint32, size uint32, retBuffPtrPos **byte, retBuffSize *int)

/*
Call host function: hostPing
Pass a string as parameter
Get a string from the host
*/
func Ping(message string) string {

	strPtrPos, size := GetStringPtrPositionAndSize(message)

	var buffPtr *byte
	var buffSize int

	hostPing(strPtrPos, size, &buffPtr, &buffSize)

	// return the string result of the host function calling
	return GetStringResult(buffPtr, buffSize)

}
