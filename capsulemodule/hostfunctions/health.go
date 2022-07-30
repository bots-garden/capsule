package hostfunctions

// TODO: move this to another package: exposedFunctions
import (
	"github.com/bots-garden/capsule/capsulemodule/memory"
)

// health is an exposed wasm function, callable by the host
//export health
//go:linkname health
func health(strPtrPos, size uint32) (strPtrPosSize uint64) {
	stringParameter := memory.GetStringParam(strPtrPos, size)
	result := stringParameter + "😄"
	pos, length := memory.GetStringPtrPositionAndSize(result)
	return memory.PackPtrPositionAndSize(pos, length)
}
