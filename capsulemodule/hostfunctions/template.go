// host functions
package hostfunctions

import (
	"errors"
	"github.com/bots-garden/capsule/capsulemodule/commons"
	"github.com/bots-garden/capsule/capsulemodule/memory"
	"strconv"
	_ "unsafe"
)

/*
   1- rename hostGetEnv
   2- rename GetEnv

*/

//export hostFunctionName
//go:linkname hostFunctionName
func hostFunctionName(paramPtrPos uint32, size uint32, retBuffPtrPos **byte, retBuffSize *int)

// FunctionName :
// Call host function: hostFunctionName
// Pass a string as parameter
// Get a string from the host
// This function is called by the wasm module
func FunctionName(param string) (string, error) {

	// transform the parameter for the host function
	paramPtrPos, size := memory.GetStringPtrPositionAndSize(param)

	var buffPtr *byte
	var buffSize int

	// call the host function
	// the result will be available in memory thanks to ` &buffPtr, &buffSize`
	hostFunctionName(paramPtrPos, size, &buffPtr, &buffSize)

	// transform the result to a string
	var resultStr = ""
	var err error
	valueStr := memory.GetStringResult(buffPtr, buffSize)

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
