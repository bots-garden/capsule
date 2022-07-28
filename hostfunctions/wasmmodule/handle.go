// host functions
package hf

var handleFunction func(string) (string, error)

func SetHandle(function func(string) (string, error)) {
	handleFunction = function
}

// TODO add detailed comments
//export callHandle
//go:linkname callHandle
func callHandle(strPtrPos, size uint32) (strPtrPosSize uint64) {
	stringParameter := GetStringParam(strPtrPos, size)

	var result string
	stringReturnByHandleFunction, errorReturnByHandleFunction := handleFunction(stringParameter)

	if errorReturnByHandleFunction != nil {
		result = CreateErrorString(errorReturnByHandleFunction.Error(), 0)
	} else {
		result = stringReturnByHandleFunction
	}

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
