// host functions
package hf

var handleHttpFunction func(string, map[string]string) (string, error)

func SetHandleHttp(function func(string, map[string]string) (string, error)) {
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
	stringReturnByHandleFunction, errorReturnByHandleFunction := handleHttpFunction(stringParameter, headers)

	if errorReturnByHandleFunction != nil {
		result = CreateErrorString(errorReturnByHandleFunction.Error(), 0)
	} else {
		result = stringReturnByHandleFunction
	}
	//TODO: CreateJsonString
	//TODO: CreateTxtString

	pos, length := GetStringPtrPositionAndSize(result)

	return PackPtrPositionAndSize(pos, length)
}

/* Function Samples
//export helloWorld
func helloWorld() (strPtrPosSize uint64) {
	strPtrPos, size := helpers.GetStringPtrPositionAndSize("👋 hello world, I'm very happy to meet you, I love what you are doing my friend")
	return helpers.PackPtrPositionAndSize(strPtrPos, size)
}

//export sayHello
func sayHello(strPtrPos, size uint32) (strPtrPosSize uint64) {
	name := helpers.GetStringParam(strPtrPos, size)
	pos, length := helpers.GetStringPtrPositionAndSize("👋 hello " + name)

	return helpers.PackPtrPositionAndSize(pos, length)
}
*/
