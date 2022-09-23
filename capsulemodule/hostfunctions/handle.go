package hostfunctions

// TODO: move this to another package: exposedFunctions
import (
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
	stringParameter := getStringParam(strPtrPos, size)
	stringParameters := commons.CreateSliceFromString(stringParameter, commons.StrSeparator)
	var result string
	stringReturnByHandleFunction, errorReturnByHandleFunction := handleFunction(stringParameters)

	if errorReturnByHandleFunction != nil {
		result = commons.CreateStringError(errorReturnByHandleFunction.Error(), 0)
	} else {
		result = stringReturnByHandleFunction
	}

	pos, length := getStringPtrPositionAndSize(result)

	return packPtrPositionAndSize(pos, length)
}

/* Function Samples
//export helloWorld
func helloWorld() (strPtrPosSize uint64) {
	strPtrPos, size := helpers.getStringPtrPositionAndSize("👋 hello world, I'm very happy to meet you, I love what you are doing my friend")
	return helpers.packPtrPositionAndSize(strPtrPos, size)
}

//export sayHello
func sayHello(strPtrPos, size uint32) (strPtrPosSize uint64) {
	name := helpers.getStringParam(strPtrPos, size)
	pos, length := helpers.getStringPtrPositionAndSize("👋 hello " + name)

	return helpers.packPtrPositionAndSize(pos, length)
}
*/
