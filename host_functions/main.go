package host_functions

import (
	"context"
	"fmt"
	"github.com/tetratelabs/wazero/api"
	"log"
)

// host wasm_modules
func LogString(ctx context.Context, module api.Module, offset, byteCount uint32) {
	buf, ok := module.Memory().Read(ctx, offset, byteCount)
	if !ok {
		log.Panicf("🟥 Memory.Read(%d, %d) out of range", offset, byteCount)
	}
	fmt.Println(string(buf))
}
