package hostfunctions

// TODO: move this to another package: exposedFunctions
import (
	"github.com/bots-garden/capsule/capsulemodule/commons"
	"github.com/bots-garden/capsule/capsulemodule/memory"
)

var handleHttpFunction func(bodyReq string, headersReq map[string]string) (bodyResp string, headersResp map[string]string, errResp error)

func SetHandleHttp(function func(string, map[string]string) (string, map[string]string, error)) {
	handleHttpFunction = function
}

// TODO add detailed comments
//export callHandleHttp
//go:linkname callHandleHttp
func callHandleHttp(strPtrPos, size uint32, headersPtrPos, headersSize uint32) (strPtrPosSize uint64) {
	//posted JSON data
	stringParameter := memory.GetStringParam(strPtrPos, size)
	headersParameter := memory.GetStringParam(headersPtrPos, headersSize)

	headersSlice := commons.CreateSliceFromString(headersParameter, "|")
	headers := commons.CreateMapFromSlice(headersSlice, ":")

	var result string
	stringReturnByHandleFunction, headersReturnByHandleFunction, errorReturnByHandleFunction := handleHttpFunction(stringParameter, headers)

	returnHeaderString := commons.CreateStringFromSlice(commons.CreateSliceFromMap(headersReturnByHandleFunction), "|")

	if errorReturnByHandleFunction != nil {
		result = commons.CreateErrorString(errorReturnByHandleFunction.Error(), 0)
	} else {
		result = CreateBodyString(stringReturnByHandleFunction)
	}

	pos, length := memory.GetStringPtrPositionAndSize(CreateResponseString(result, returnHeaderString))

	return memory.PackPtrPositionAndSize(pos, length)
}

func CreateBodyString(message string) string {
	return "[BODY]" + message
}

func CreateResponseString(result, headers string) string {
	return result + "[HEADERS]" + headers
}
