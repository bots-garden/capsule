package hostfunctions

// TODO: move this to another package: exposedFunctions
import (
	"github.com/bots-garden/capsule/capsulemodule/memory"
	"github.com/bots-garden/capsule/commons"
)

var natsHandleFunction func([]string)

func OnNatsMessage(function func([]string)) {
	natsHandleFunction = function
}

//export callNatsMessageHandle
//go:linkname callNatsMessageHandle
func callNatsMessageHandle(strPtrPos, size uint32) (strPtrPosSize uint64) {
	stringParameter := memory.GetStringParam(strPtrPos, size)
	//fmt.Println("🤗 stringParameter", stringParameter)
	stringParameters := commons.CreateSliceFromString(stringParameter, commons.StrSeparator)
	var result string
	natsHandleFunction(stringParameters)

	pos, length := memory.GetStringPtrPositionAndSize(result)

	return memory.PackPtrPositionAndSize(pos, length)
}
