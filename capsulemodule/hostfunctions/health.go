package hostfunctions

import (
    "github.com/bots-garden/capsule/capsulemodule/memory"
)

//export health
//go:linkname health
func health(strPtrPos, size uint32) (strPtrPosSize uint64) {
    stringParameter := memory.GetStringParam(strPtrPos, size)
    result := stringParameter
    pos, length := memory.GetStringPtrPositionAndSize(result)
    return memory.PackPtrPositionAndSize(pos, length)
}
