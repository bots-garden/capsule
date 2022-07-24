package host_functions

import (
	"context"
	"fmt"
	"log"

	"github.com/tetratelabs/wazero/api"
)

// host functions for the wasm module

// print a string to the console
func LogString(ctx context.Context, module api.Module, offset, byteCount uint32) {
	buf, ok := module.Memory().Read(ctx, offset, byteCount)
	if !ok {
		log.Panicf("🟥 Memory.Read(%d, %d) out of range", offset, byteCount)
	}
	fmt.Println(string(buf))
}

// return information about the host
func GetHostInformation(ctx context.Context, module api.Module, retBufPtrPos, retBufSize uint32) {

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
	module.Memory().WriteUint32Le(ctx, retBufPtrPos, offset)
	module.Memory().WriteUint32Le(ctx, retBufSize, uint32(lengthOfTheMessage))

	// add the message to the memory of the module
	module.Memory().Write(ctx, offset, []byte(message))

}


func Ping(ctx context.Context, module api.Module, offset, byteCount, retBufPtrPos, retBufSize uint32 ) {
	// get string from the wasm module function (from memory)
	buf, ok := module.Memory().Read(ctx, offset, byteCount)
	  if !ok {
		  log.Panicf("🟥 Memory.Read(%d, %d) out of range", offset, byteCount)
	  }
	stringMessageFromFunction := string(buf)
	stringMessageFromHost :=  "🏓 pong: " + stringMessageFromFunction
  
	// write the new string to the "shared memory"
	lengthOfTheMessage := len(stringMessageFromHost)
	results, err := module.ExportedFunction("allocateBuffer").Call(ctx, uint64(lengthOfTheMessage))
	if err != nil {
	  log.Panicln(err)
	}
  
	retOffset := uint32(results[0])
	module.Memory().WriteUint32Le(ctx, retBufPtrPos, retOffset)
	module.Memory().WriteUint32Le(ctx, retBufSize, uint32(lengthOfTheMessage))
  
	// add the message to the memory of the module
	module.Memory().Write(ctx, retOffset, []byte(stringMessageFromHost))
  }
  