package hostfunctions

import "github.com/bots-garden/capsule/capsulemodule/memory"

//export health
//go:linkname health
func health() (strPtrPosSize uint64) {
	strPtrPos, size := memory.GetStringPtrPositionAndSize(`{"health":"ok"}`)
	Log("👋👋👋👋👋👋")
	return memory.PackPtrPositionAndSize(strPtrPos, size)
}
