package hostfunctions

import (
    "context"
    "github.com/bots-garden/capsule/capsule-launcher/hostfunctions/memory"
    "github.com/tetratelabs/wazero/api"
)

// HostInformation updated in `services/http/http.go`
var HostInformation = ""

// GetHostInformation returns information about the host
var GetHostInformation = api.GoModuleFunc(func(ctx context.Context, module api.Module, stack []uint64) {

    message := HostInformation

    positionReturnBuffer := uint32(stack[0])
    lengthReturnBuffer := uint32(stack[1])

    memory.WriteStringToMemory(message, ctx, module, positionReturnBuffer, lengthReturnBuffer)

    stack[0] = 0 // return 0
})
