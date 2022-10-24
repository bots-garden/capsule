package hostfunctions

// TODO: move this to another package: exposedFunctions
import (
	"github.com/bots-garden/capsule/commons"
	"strconv"
)

/* previous version
var handleHttpFunction func(bodyReq string, headersReq map[string]string) (
    bodyResp string, headersResp map[string]string, errResp error)
*/

var handleHttpNextFunction func(req Request) (resp Response, errResp error)

func SetHandleHttpNext(function func(request Request) (Response, error)) {
	Log("🤖🖐🎃[SetHandleHttpNext]")

	handleHttpNextFunction = function
}

// The name "callHandleHttp" of the exported function is defined/declared
// in `wasmrunner.go`, function: GetNewWasmRuntimeForHttp

//export callHandleHttpNext
//go:linkname callHandleHttpNext
func callHandleHttpNext(reqId uint32) (strPtrPosSize uint64) {
	//stringParameter := getStringParam(strPtrPos, size)

	reqParams, errReqParams := RequestParamsGet(reqId)

	Log("🤖🖐[reqParams]:" + strconv.FormatUint(uint64(reqId), 10))

	if errReqParams != nil {
		// TODO
	}

	bodyParameter := reqParams[0]
	headersParameter := reqParams[1]
	uriParameter := reqParams[2]
	methodParameter := reqParams[3]

	Log("🤖🖐[bodyParameter]:" + bodyParameter)
	Log("🤖🖐[headersParameter]:" + headersParameter)
	Log("🤖🖐[uriParameter]:" + uriParameter)
	Log("🤖🖐[methodParameter]:" + methodParameter)

	headersSlice := commons.CreateSliceFromString(headersParameter, commons.StrHeadersSeparator) //👋
	headers := commons.CreateMapFromSlice(headersSlice, commons.FieldSeparator)

	var result string
	//stringReturnByHandleFunction, headersReturnByHandleFunction, errorReturnByHandleFunction := handleHttpFunction(bodyParameter, headers)
	responseReturnByHandleFunction, errorReturnByHandleFunction := handleHttpNextFunction(Request{bodyParameter, headers, uriParameter, methodParameter})

	returnHeaderString := commons.CreateStringFromSlice(commons.CreateSliceFromMap(responseReturnByHandleFunction.Headers), commons.StrSeparator)

	if errorReturnByHandleFunction != nil {
		result = commons.CreateStringError(errorReturnByHandleFunction.Error(), 0)
	} else {
		result = CreateBodyString(responseReturnByHandleFunction.Body)
	}

	// merge body and headers response
	pos, length := getStringPtrPositionAndSize(CreateResponseString(result, returnHeaderString))

	return packPtrPositionAndSize(pos, length)
}
