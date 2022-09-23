package hostfunctions

// TODO: move this to another package: exposedFunctions
import (
	"github.com/bots-garden/capsule/capsulemodule/memory"
	"github.com/bots-garden/capsule/commons"
)

var mqttHandleFunction func([]string)

func OnMqttMessage(function func([]string)) {
	mqttHandleFunction = function
}

//export callMqttMessageHandle
//go:linkname callMqttMessageHandle
func callMqttMessageHandle(strPtrPos, size uint32) (strPtrPosSize uint64) {
	stringParameter := memory.GetStringParam(strPtrPos, size)
	//fmt.Println("🤗 stringParameter", stringParameter)
	stringParameters := commons.CreateSliceFromString(stringParameter, commons.StrSeparator)
	var result string
	mqttHandleFunction(stringParameters)

	pos, length := memory.GetStringPtrPositionAndSize(result)

	return memory.PackPtrPositionAndSize(pos, length)
}
