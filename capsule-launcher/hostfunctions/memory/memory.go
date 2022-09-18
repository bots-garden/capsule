package memory

import (
	"context"
	"github.com/tetratelabs/wazero/api"
	"log"
)

// WriteStringToMemory :
// Write string to the memory of the module
// The Host writes to memory
func WriteStringToMemory(text string, ctx context.Context, module api.Module,
	retBuffPtrPos, retBuffSize uint32) {

	stringMessageFromHost := text
	lengthOfTheMessage := len(stringMessageFromHost)
	results, err := module.ExportedFunction("allocateBuffer").Call(ctx, uint64(lengthOfTheMessage))

	//TODO handle this in another way?
	if err != nil {
		log.Panicln(err)
	}

	retOffset := uint32(results[0])
	module.Memory().WriteUint32Le(ctx, retBuffPtrPos, retOffset)
	module.Memory().WriteUint32Le(ctx, retBuffSize, uint32(lengthOfTheMessage))

	// add the message to the memory of the module
	module.Memory().Write(ctx, retOffset, []byte(stringMessageFromHost))

}

// ReadStringFromMemory :
// Get string from the module's memory (written by the module)
// (argument of a function)
func ReadStringFromMemory(ctx context.Context, module api.Module, contentOffset, contentByteCount uint32) string {
	contentBuff, ok := module.Memory().Read(ctx, contentOffset, contentByteCount)
	if !ok {
		log.Panicf("🟥 Memory.Read(%d, %d) out of range", contentOffset, contentByteCount)
	}
	contentStr := string(contentBuff)
	return contentStr
}
