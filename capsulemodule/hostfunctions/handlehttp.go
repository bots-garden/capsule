// host functions
package hostfunctions

var handleHttpFunction func(bodyReq string, headersReq map[string]string) (bodyResp string, headersResp map[string]string, errResp error)

func SetHandleHttp(function func(string, map[string]string) (string, map[string]string, error)) {
	handleHttpFunction = function
}

// TODO add detailed comments
//export callHandleHttp
//go:linkname callHandleHttp
func callHandleHttp(strPtrPos, size uint32, headersPtrPos, headersSize uint32) (strPtrPosSize uint64) {
	//posted JSON data
	stringParameter := GetStringParam(strPtrPos, size)
	headersParameter := GetStringParam(headersPtrPos, headersSize)

	headersSlice := CreateSliceFromString(headersParameter, "|")
	headers := CreateMapFromSlice(headersSlice, ":")

	var result string
	stringReturnByHandleFunction, headersReturnByHandleFunction, errorReturnByHandleFunction := handleHttpFunction(stringParameter, headers)

	returnHeaderString := CreateStringFromSlice(CreateSliceFromMap(headersReturnByHandleFunction), "|")

	if errorReturnByHandleFunction != nil {
		result = CreateErrorString(errorReturnByHandleFunction.Error(), 0)
	} else {
		result = CreateBodyString(stringReturnByHandleFunction)
	}

	pos, length := GetStringPtrPositionAndSize(CreateResponseString(result, returnHeaderString))

	return PackPtrPositionAndSize(pos, length)
}

func CreateBodyString(message string) string {
	return "[BODY]" + message
}

func CreateResponseString(result, headers string) string {
	return result + "[HEADERS]" + headers
}
