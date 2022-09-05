package hostfunctions

// TODO: move this to another package: exposedFunctions
import (
	"github.com/bots-garden/capsule/capsulemodule/memory"
	"github.com/bots-garden/capsule/commons"
)

/* previous version
var handleHttpFunction func(bodyReq string, headersReq map[string]string) (
    bodyResp string, headersResp map[string]string, errResp error)
*/

var handleHttpFunction func(req Request) (resp Response, errResp error)

func SetHandleHttp(function func(request Request) (Response, error)) {
	handleHttpFunction = function
}

// The name "callHandleHttp" of the exported function is defined/declared
// in `wasmrunner.go`, function: GetNewWasmRuntimeForHttp

//export callHandleHttp
//go:linkname callHandleHttp
func callHandleHttp(bodyPtrPos, bodySize, uriPtrPos, uriSize, headersPtrPos, headersSize, methodPtrPos, methodSize uint32) (strPtrPosSize uint64) {
	//posted JSON data
	bodyParameter := memory.GetStringParam(bodyPtrPos, bodySize)
	headersParameter := memory.GetStringParam(headersPtrPos, headersSize)
	uriParameter := memory.GetStringParam(uriPtrPos, uriSize)
	methodParameter := memory.GetStringParam(methodPtrPos, methodSize)

	headersSlice := commons.CreateSliceFromString(headersParameter, commons.StrSeparator)
	headers := commons.CreateMapFromSlice(headersSlice, commons.FieldSeparator)

	var result string
	//stringReturnByHandleFunction, headersReturnByHandleFunction, errorReturnByHandleFunction := handleHttpFunction(bodyParameter, headers)
	responseReturnByHandleFunction, errorReturnByHandleFunction := handleHttpFunction(Request{bodyParameter, headers, uriParameter, methodParameter})

	returnHeaderString := commons.CreateStringFromSlice(commons.CreateSliceFromMap(responseReturnByHandleFunction.Headers), commons.StrSeparator)

	if errorReturnByHandleFunction != nil {
		result = commons.CreateStringError(errorReturnByHandleFunction.Error(), 0)
	} else {
		result = CreateBodyString(responseReturnByHandleFunction.Body)
	}

	// merge body and headers response
	pos, length := memory.GetStringPtrPositionAndSize(CreateResponseString(result, returnHeaderString))

	return memory.PackPtrPositionAndSize(pos, length)
}

func CreateBodyString(message string) string {
	return "[BODY]" + message
}

func CreateResponseString(result, headers string) string {
	return result + "[HEADERS]" + headers
}
