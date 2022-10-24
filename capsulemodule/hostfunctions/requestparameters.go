// host functions
package hostfunctions

import (
	"errors"
	"fmt"
	"strconv"
	_ "unsafe"

	"github.com/bots-garden/capsule/commons"
)

//export hostRequestParamsGet
//go:linkname hostRequestParamsGet
func hostRequestParamsGet(reqId uint32, retBuffPtrPos **byte, retBuffSize *int)

// RequestParamsGet :
// This function is called by the wasm module
func RequestParamsGet(reqId uint32) ([]string, error) {

	// transform the parameter for the host function
	//reqIdPtrPos, reqIdSize := getStringPtrPositionAndSize(reqId)

	var buffPtr *byte
	var buffSize int

	// call the host function
	// the result will be available in memory thanks to ` &buffPtr, &buffSize`

	fmt.Println("🤖🐻[hostRequestParamsGet]", reqId)

	hostRequestParamsGet(reqId, &buffPtr, &buffSize)

	// transform the result to a string
	var err error
	valueStr := getStringResult(buffPtr, buffSize)

	fmt.Println("🤖🐻[valueStr]", valueStr)

	// 🖐 this string contains (in this order):
	// reqParams.JsonData,
	// reqParams.Headers,
	// reqParams.Uri,
	// reqParams.Method

	// check the return value
	if commons.IsErrorString(valueStr) {
		errorMessage, errorCode := commons.GetErrorStringInfo(valueStr)
		if errorCode == 0 {
			err = errors.New(errorMessage)
		} else {
			err = errors.New(errorMessage + " (" + strconv.Itoa(errorCode) + ")")
		}
		return nil, err

	} else {
		result := commons.CreateSliceFromString(valueStr, commons.StrSeparator)
		return result, nil
	}

}
