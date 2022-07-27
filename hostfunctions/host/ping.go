package hostfunctions

import (
	"context"
	"log"

	"github.com/tetratelabs/wazero/api"
)

// host functions for the wasm module

// string parameter, return string
func Ping(ctx context.Context, module api.Module, offset, byteCount, retBuffPtrPos, retBuffSize uint32) {
	// get string from the wasm module function (from memory)
	buf, ok := module.Memory().Read(ctx, offset, byteCount)
	if !ok {
		log.Panicf("🟥 Memory.Read(%d, %d) out of range", offset, byteCount)
	}
	stringMessageFromFunction := string(buf)
	stringMessageFromHost := "🏓 pong: " + stringMessageFromFunction

	// write the new string to the "shared memory"
	lengthOfTheMessage := len(stringMessageFromHost)
	results, err := module.ExportedFunction("allocateBuffer").Call(ctx, uint64(lengthOfTheMessage))
	if err != nil {
		log.Panicln(err)
	}

	retOffset := uint32(results[0])
	module.Memory().WriteUint32Le(ctx, retBuffPtrPos, retOffset)
	module.Memory().WriteUint32Le(ctx, retBuffSize, uint32(lengthOfTheMessage))

	// add the message to the memory of the module
	module.Memory().Write(ctx, retOffset, []byte(stringMessageFromHost))
}
