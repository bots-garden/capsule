// host functions
package hostfunctions

import (
	"errors"
	"strconv"
	_ "unsafe"
)

//export hostWriteFile
//go:linkname hostWriteFile
func hostWriteFile(filePathPtrPos uint32, size uint32, contentPtrPos uint32, contentSize uint32, retBuffPtrPos **byte, retBuffSize *int)

// WriteFile : Call host function: hostReadFile
// Pass a string as parameter
//Get a string from the host

func WriteFile(filePath string, content string) (string, error) {

	filePathPtrPos, size := GetStringPtrPositionAndSize(filePath)
	contentPtrPos, contentSize := GetStringPtrPositionAndSize(content)

	var buffPtr *byte
	var buffSize int

	hostWriteFile(filePathPtrPos, size, contentPtrPos, contentSize, &buffPtr, &buffSize)

	var resultStr = ""
	var err error
	valueStr := GetStringResult(buffPtr, buffSize)

	// check the return value
	if IsErrorString(valueStr) {
		errorMessage, errorCode := GetErrorStringInfo(valueStr)
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
