package hostfunctions

import (
	_ "unsafe"
)

//export hostGetExitError
//go:linkname hostGetExitError
func hostGetExitError(retBuffPtrPos **byte, retBuffSize *int)

func GetExitError() string {
	var buffPtr *byte
	var buffSize int

	hostGetExitError(&buffPtr, &buffSize)

	// return the string result of the host function calling
	return getStringResult(buffPtr, buffSize)
}

//export hostGetExitCode
//go:linkname hostGetExitCode
func hostGetExitCode(retBuffPtrPos **byte, retBuffSize *int)

func GetExitCode() string { // I return a string because I will probably use it to return my own error codes

	var buffPtr *byte
	var buffSize int

	hostGetExitCode(&buffPtr, &buffSize)

	// return the string result of the host function calling
	strExitCode := getStringResult(buffPtr, buffSize)
	return strExitCode
}
