// host functions
package hostfunctions

import _ "unsafe"

//export hostGetHostInformation
//go:linkname hostGetHostInformation
func hostGetHostInformation(retBuffPtrPos **byte, retBuffSize *int)

/*
Call host function: hostGetHostInformation
Get a string with the information about the host
*/
func GetHostInformation() string {
	var buffPtr *byte
	var buffSize int

	hostGetHostInformation(&buffPtr, &buffSize)

	// return the string result of the host function calling
	return GetStringResult(buffPtr, buffSize)
}
