// host functions
package hostfunctions

import (
	"errors"
	"strconv"
	_ "unsafe"

	"github.com/bots-garden/capsule/commons"
)

//export hostMemorySet
//go:linkname hostMemorySet
func hostMemorySet(keyPtrPos, keySize, valuePtrPos, valueSize uint32, retBuffPtrPos **byte, retBuffSize *int)

//export hostMemoryGet
//go:linkname hostMemoryGet
func hostMemoryGet(keyPtrPos, keySize uint32, retBuffPtrPos **byte, retBuffSize *int)

//export hostMemoryKeys
//go:linkname hostMemoryKeys
func hostMemoryKeys(retBuffPtrPos **byte, retBuffSize *int)

// MemorySet :
// Call host function: hostMemorySet
// This function is called by the wasm module
func MemorySet(key string, value string) (string, error) {

	// transform the parameters for the host function
	keyPtrPos, keySize := getStringPtrPositionAndSize(key)
	valuePtrPos, valueSize := getStringPtrPositionAndSize(value)

	var buffPtr *byte
	var buffSize int

	// call the host function
	// the result will be available in memory thanks to ` &buffPtr, &buffSize`
	hostMemorySet(keyPtrPos, keySize, valuePtrPos, valueSize, &buffPtr, &buffSize)

	// transform the result to a string
	var resultStr = ""
	var err error
	valueStr := getStringResult(buffPtr, buffSize)

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

// MemoryGet :
// This function is called by the wasm module
func MemoryGet(key string) (string, error) {

	// transform the parameter for the host function
	keyPtrPos, keySize := getStringPtrPositionAndSize(key)

	var buffPtr *byte
	var buffSize int

	// call the host function
	// the result will be available in memory thanks to ` &buffPtr, &buffSize`
	hostMemoryGet(keyPtrPos, keySize, &buffPtr, &buffSize)

	// transform the result to a string
	var resultStr = ""
	var err error
	valueStr := getStringResult(buffPtr, buffSize)

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

// MemoryKeys :
// This function is called by the wasm module
func MemoryKeys() ([]string, error) {
	// transform the parameter for the host function
	var buffPtr *byte
	var buffSize int

	// call the host function
	// the result will be available in memory thanks to ` &buffPtr, &buffSize`
	hostMemoryKeys(&buffPtr, &buffSize)

	// transform the result to a string
	var results []string
	var err error
	valueStr := getStringResult(buffPtr, buffSize)

	// check the return value
	if commons.IsErrorString(valueStr) {
		errorMessage, errorCode := commons.GetErrorStringInfo(valueStr)
		if errorCode == 0 {
			err = errors.New(errorMessage)
		} else {
			err = errors.New(errorMessage + " (" + strconv.Itoa(errorCode) + ")")
		}

	} else {
		results = commons.CreateSliceFromString(valueStr, commons.StrSeparator)
	}
	return results, err
}
