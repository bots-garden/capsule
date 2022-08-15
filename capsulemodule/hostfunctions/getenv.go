// Package hf host functions
package hostfunctions

import (
	"errors"
	"github.com/bots-garden/capsule/capsulemodule/memory"
	"github.com/bots-garden/capsule/commons"
	"strconv"
	_ "unsafe"
)

//export hostGetEnv
//go:linkname hostGetEnv
func hostGetEnv(varNamePtrPos uint32, size uint32, retBuffPtrPos **byte, retBuffSize *int)

// GetEnv :
// Call host function: hostGetEnv
// Pass a string as parameter
// Get a string from the host
// This function is called by the wasm module
func GetEnv(varName string) (string, error) {

	// transform the parameter for the host function
	varNamePtrPos, size := memory.GetStringPtrPositionAndSize(varName)

	var buffPtr *byte
	var buffSize int

	// call the host function
	// the result will be available in memory thanks to ` &buffPtr, &buffSize`
	hostGetEnv(varNamePtrPos, size, &buffPtr, &buffSize)

	// transform the result to a string
	var resultStr = ""
	var err error
	valueStr := memory.GetStringResult(buffPtr, buffSize)
	//Log("✅ " + valueStr)

	// check the return value
	if commons.IsErrorString(valueStr) {
		errorMessage, errorCode := commons.GetErrorStringInfo(valueStr)
		if errorCode == 0 {
			err = errors.New(errorMessage)
		} else {
			err = errors.New(errorMessage + " (" + strconv.Itoa(errorCode) + ")")
		}

	} else {
		resultStr = valueStr
	}
	return resultStr, err

}
