package hostfunctions

import (
	"context"
	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions/memory"
	"github.com/bots-garden/capsule/commons"
	"github.com/tetratelabs/wazero/api"
)

var memoryMap = map[string]string{"capsule_version": commons.CapsuleVersion()} // my little easter 🥚

var MemorySet = api.GoModuleFunc(func(ctx context.Context, module api.Module, params []uint64) []uint64 {
	keyPosition := uint32(params[0])
	keyLength := uint32(params[1])
	keyStr := memory.ReadStringFromMemory(ctx, module, keyPosition, keyLength)

	valuePosition := uint32(params[2])
	valueLength := uint32(params[3])
	valueStr := memory.ReadStringFromMemory(ctx, module, valuePosition, valueLength)

	memoryMap[keyStr] = valueStr

	stringResultFromHost := "[OK](" + keyStr + ":" + valueStr + ")"

	positionReturnBuffer := uint32(params[4])
	lengthReturnBuffer := uint32(params[5])

	memory.WriteStringToMemory(stringResultFromHost, ctx, module, positionReturnBuffer, lengthReturnBuffer)

	return []uint64{0}

})

var MemoryGet = api.GoModuleFunc(func(ctx context.Context, module api.Module, params []uint64) []uint64 {

	keyPosition := uint32(params[0])
	keyLength := uint32(params[1])
	keyStr := memory.ReadStringFromMemory(ctx, module, keyPosition, keyLength)

	valueStr := memoryMap[keyStr]

	var stringResultFromHost = ""

	if len(valueStr) == 0 {
		stringResultFromHost = commons.CreateStringError("key does not exist", 0)
		// if code 0 don't display code in the error message
	} else {
		stringResultFromHost = valueStr
	}

	positionReturnBuffer := uint32(params[2])
	lengthReturnBuffer := uint32(params[3])

	memory.WriteStringToMemory(stringResultFromHost, ctx, module, positionReturnBuffer, lengthReturnBuffer)

	return []uint64{0}

})

var MemoryKeys = api.GoModuleFunc(func(ctx context.Context, module api.Module, params []uint64) []uint64 {
	var keys []string
	for key, _ := range memoryMap {
		keys = append(keys, key)
	}
	stringResultFromHost := commons.CreateStringFromSlice(keys, commons.StrSeparator)

	positionReturnBuffer := uint32(params[0])
	lengthReturnBuffer := uint32(params[1])

	memory.WriteStringToMemory(stringResultFromHost, ctx, module, positionReturnBuffer, lengthReturnBuffer)

	return []uint64{0}

})
