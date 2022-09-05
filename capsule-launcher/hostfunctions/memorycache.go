package hostfunctions

import (
	"context"
	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions/memory"
	"github.com/bots-garden/capsule/commons"
	"github.com/tetratelabs/wazero/api"
)

var memoryMap = map[string]string{"capsule_version": commons.CapsuleVersion()} // my little easter 🥚

func MemorySet(ctx context.Context, module api.Module, keyOffset, keyByteCount, valueOffSet, valueByteCount, retBuffPtrPos, retBuffSize uint32) {
	keyStr := memory.ReadStringFromMemory(ctx, module, keyOffset, keyByteCount)
	valueStr := memory.ReadStringFromMemory(ctx, module, valueOffSet, valueByteCount)
	memoryMap[keyStr] = valueStr

	stringResultFromHost := "[OK](" + keyStr + ":" + valueStr + ")"

	memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)

}

func MemoryGet(ctx context.Context, module api.Module, keyOffset, keyByteCount, retBuffPtrPos, retBuffSize uint32) {
	keyStr := memory.ReadStringFromMemory(ctx, module, keyOffset, keyByteCount)
	valueStr := memoryMap[keyStr]

	var stringResultFromHost = ""

	if len(valueStr) == 0 {
		stringResultFromHost = commons.CreateStringError("key does not exist", 0)
		// if code 0 don't display code in the error message
	} else {
		stringResultFromHost = valueStr
	}

	memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)
}

func MemoryKeys(ctx context.Context, module api.Module, retBuffPtrPos, retBuffSize uint32) {
	var keys []string
	for key, _ := range memoryMap {
		keys = append(keys, key)
	}
	stringResultFromHost := commons.CreateStringFromSlice(keys, commons.StrSeparator)
	memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)

}
