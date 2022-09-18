package hostfunctions

// TODO: move this to another package: exposedFunctions
import (
	"github.com/bots-garden/capsule/capsulemodule/memory"
	"github.com/bots-garden/capsule/commons"
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
	stringParameters := commons.CreateSliceFromString(stringParameter, commons.StrSeparator)
	var result string
	stringReturnByHandleFunction, errorReturnByHandleFunction := handleFunction(stringParameters)

	if errorReturnByHandleFunction != nil {
		result = commons.CreateStringError(errorReturnByHandleFunction.Error(), 0)
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
