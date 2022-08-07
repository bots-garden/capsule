package hostfunctions

import (
	"context"
	"log"

	"github.com/tetratelabs/wazero/api"
)

var HostInformation = ""

// GetHostInformation returns information about the host
func GetHostInformation(ctx context.Context, module api.Module, retBuffPtrPos, retBuffSize uint32) {

	message := HostInformation
	lengthOfTheMessage := len(message)

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
