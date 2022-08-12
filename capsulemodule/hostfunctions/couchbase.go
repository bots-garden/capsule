// host functions
package hostfunctions

import (
	"errors"
	"strconv"
	_ "unsafe"

	"github.com/bots-garden/capsule/capsulelauncher/commons"
	"github.com/bots-garden/capsule/capsulemodule/memory"
)

//export hostCouchBaseQuery
//go:linkname hostCouchBaseQuery
func hostCouchBaseQuery(queryPtrPos, querySize uint32, retBuffPtrPos **byte, retBuffSize *int)

// CouchBaseQuery :
// This function is called by the wasm module
func CouchBaseQuery(query string) (string, error) {

	// transform the parameter for the host function
	queryPtrPos, querySize := memory.GetStringPtrPositionAndSize(query)

	var buffPtr *byte
	var buffSize int

	// call the host function
	// the result will be available in memory thanks to ` &buffPtr, &buffSize`
	hostCouchBaseQuery(queryPtrPos, querySize, &buffPtr, &buffSize)

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
