package hostfunctions

import (
	"context"
	"log"

	"github.com/tetratelabs/wazero/api"
)

// host functions for the wasm module

// return information about the host
func GetHostInformation(ctx context.Context, module api.Module, retBuffPtrPos, retBuffSize uint32) {
	// TODO: return something more interesting
	// TODO: cpu usage, memory,...
	message := "💊 Capsule [wasm launcher] v0.0.0"
	lengthOfTheMessage := len(message)

	// TODO: create an helper from this
	// Allocate buffer in the wasm module memory
	results, err := module.ExportedFunction("allocateBuffer").Call(ctx, uint64(lengthOfTheMessage))
	if err != nil {
		log.Panicln(err)
	}

	offset := uint32(results[0])
	module.Memory().WriteUint32Le(ctx, retBuffPtrPos, offset)
	module.Memory().WriteUint32Le(ctx, retBuffSize, uint32(lengthOfTheMessage))

	// add the message to the memory of the module
	module.Memory().Write(ctx, offset, []byte(message))
}
