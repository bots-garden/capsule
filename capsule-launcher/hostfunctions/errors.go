package hostfunctions

import (
	"context"
	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions/memory"
	"github.com/bots-garden/capsule/commons"
	"github.com/tetratelabs/wazero/api"
	"strconv"
)

func GetExitError(ctx context.Context, module api.Module, retBuffPtrPos, retBuffSize uint32) {
	exitError := commons.GetExitError()
	memory.WriteStringToMemory(exitError, ctx, module, retBuffPtrPos, retBuffSize)
}

func GetExitCode(ctx context.Context, module api.Module, retBuffPtrPos, retBuffSize uint32) {
	exitCode := strconv.Itoa(commons.GetExitCode())
	//fmt.Println("📝", exitCode)
	memory.WriteStringToMemory(exitCode, ctx, module, retBuffPtrPos, retBuffSize)
}
