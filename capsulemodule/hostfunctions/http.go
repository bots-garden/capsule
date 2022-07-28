// host functions
package hostfunctions

import (
	"errors"
	"github.com/bots-garden/capsule/capsulemodule/commons"
	"github.com/bots-garden/capsule/capsulemodule/memory"
	"strconv"
	_ "unsafe"
)

//export hostHttp
//go:linkname hostHttp
func hostHttp(urlOffset uint32, urlByteCount uint32, methodOffSet uint32, methodByteCount uint32, headersOffSet uint32, headersMethodByteCount uint32, bodyOffset uint32, bodyByteCount uint32, retBuffPtrPos **byte, retBuffSize *int)

func Http(url, method string, headers map[string]string, body string) (string, error) {
	// Get parameters from the wasm module
	// Prepare parameters for the host function call
	urlStrPos, urlStrSize := memory.GetStringPtrPositionAndSize(url)
	methodStrPos, methodStrSize := memory.GetStringPtrPositionAndSize(method)

	/*
	   headers := map[string]string{"Accept": "application/json", "Content-Type": " text/html; charset=UTF-8"}

	   headersStringSlice => ["Accept:application/json", ”Content-Type:text/html; charset=UTF-8"]

	   headerString => "Accept:application/json|Content-Type:text/html; charset=UTF-8"
	*/

	headersStringSlice := commons.CreateSliceFromMap(headers)
	headerString := commons.CreateStringFromSlice(headersStringSlice, "|")

	headersStrPos, headersStrSize := memory.GetStringPtrPositionAndSize(headerString)

	bodyStrPos, bodyStrSize := memory.GetStringPtrPositionAndSize(body)

	// This will be used to retrieve the return value (result)
	var buffPtr *byte
	var buffSize int

	// 🖐 call the host function
	// buffPtr, buffSize allows to retrieve the result of the function call
	hostHttp(urlStrPos, urlStrSize, methodStrPos, methodStrSize, headersStrPos, headersStrSize, bodyStrPos, bodyStrSize, &buffPtr, &buffSize)

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
