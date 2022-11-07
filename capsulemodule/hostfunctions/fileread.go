package hostfunctions

import (
	"errors"
	"github.com/bots-garden/capsule/commons"
	"strconv"
	_ "unsafe"
)

//export hostReadFile
//go:linkname hostReadFile
func hostReadFile(filePathPosition uint32, filePathLength uint32, returnBufferPosition **byte, returnBufferLength *int) uint32

// ReadFile :
// Call host function: hostReadFile
// Pass a string as parameter
// Get a string from the host

func ReadFile(filePath string) (string, error) {

	filePathPtrPos, size := getStringPtrPositionAndSize(filePath)

	var buffPtr *byte
	var buffSize int

	hostReadFile(filePathPtrPos, size, &buffPtr, &buffSize)

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
