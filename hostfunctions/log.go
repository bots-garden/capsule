package hostfunctions

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
