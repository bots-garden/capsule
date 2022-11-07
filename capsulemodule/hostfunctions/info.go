// host functions
package hostfunctions

import (
	_ "unsafe"
)

//export hostGetHostInformation
//go:linkname hostGetHostInformation
func hostGetHostInformation(positionReturnBuffer **byte, lengthReturnBuffer *int) uint32

/*
Call host function: hostGetHostInformation
Get a string with the information about the host
*/
func GetHostInformation() string {
	var buffPtr *byte
	var buffSize int

	hostGetHostInformation(&buffPtr, &buffSize)

	return getStringResult(buffPtr, buffSize)
}
