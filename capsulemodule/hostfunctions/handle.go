// host functions
package hostfunctions

import (
	"github.com/bots-garden/capsule/capsulemodule/commons"
	"github.com/bots-garden/capsule/capsulemodule/memory"
)

var handleFunction func([]string) (string, error)

func SetHandle(function func([]string) (string, error)) {
	handleFunction = function
}

// TODO add detailed comments
//export callHandle
//go:linkname callHandle
func callHandle(strPtrPos, size uint32) (strPtrPosSize uint64) {
	stringParameter := memory.GetStringParam(strPtrPos, size)
	//TODO change the separator
	stringParameters := commons.CreateSliceFromString(stringParameter, "°")
	var result string
	stringReturnByHandleFunction, errorReturnByHandleFunction := handleFunction(stringParameters)

	if errorReturnByHandleFunction != nil {
		result = commons.CreateErrorString(errorReturnByHandleFunction.Error(), 0)
	} else {
		result = stringReturnByHandleFunction
	}

	pos, length := memory.GetStringPtrPositionAndSize(result)

	return memory.PackPtrPositionAndSize(pos, length)
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
