// host functions
package hostfunctions

import (
	"errors"
	"strconv"
	_ "unsafe"

	"github.com/bots-garden/capsule/capsulemodule/commons"
	"github.com/bots-garden/capsule/capsulemodule/memory"
)

//export hostRedisSet
//go:linkname hostRedisSet
func hostRedisSet(keyPtrPos, keySize, valuePtrPos, valueSize uint32, retBuffPtrPos **byte, retBuffSize *int)

//export hostRedisGet
//go:linkname hostRedisGet
func hostRedisGet(keyPtrPos, keySize uint32, retBuffPtrPos **byte, retBuffSize *int)

// RedisSet :
// Call host function: hostRedisSet
// This function is called by the wasm module
func RedisSet(key string, value string) (string, error) {

	// transform the parameters for the host function
	keyPtrPos, keySize := memory.GetStringPtrPositionAndSize(key)
	valuePtrPos, valueSize := memory.GetStringPtrPositionAndSize(value)

	var buffPtr *byte
	var buffSize int

	// call the host function
	// the result will be available in memory thanks to ` &buffPtr, &buffSize`
	hostRedisSet(keyPtrPos, keySize, valuePtrPos, valueSize, &buffPtr, &buffSize)

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

// RedisGet :
// This function is called by the wasm module
func RedisGet(key string) (string, error) {

	// transform the parameter for the host function
	keyPtrPos, keySize := memory.GetStringPtrPositionAndSize(key)

	var buffPtr *byte
	var buffSize int

	// call the host function
	// the result will be available in memory thanks to ` &buffPtr, &buffSize`
	hostRedisGet(keyPtrPos, keySize, &buffPtr, &buffSize)

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
