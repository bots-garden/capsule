package hostfunctions

import (
	"context"
	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions/memory"
	"github.com/bots-garden/capsule/commons"
	"github.com/tetratelabs/wazero/api"
	"strconv"
)

var GetExitError = api.GoModuleFunc(func(ctx context.Context, module api.Module, stack []uint64) {
	exitError := commons.GetExitError()
	retBuffPtrPos := uint32(stack[0])
	retBuffSize := uint32(stack[1])
	memory.WriteStringToMemory(exitError, ctx, module, retBuffPtrPos, retBuffSize)
	stack[0] = 0 // return 0
})

var GetExitCode = api.GoModuleFunc(func(ctx context.Context, module api.Module, stack []uint64) {
	exitCode := strconv.Itoa(commons.GetExitCode())
	//fmt.Println("📝", exitCode)
	retBuffPtrPos := uint32(stack[0])
	retBuffSize := uint32(stack[1])
	memory.WriteStringToMemory(exitCode, ctx, module, retBuffPtrPos, retBuffSize)
	stack[0] = 0 // return 0
})
