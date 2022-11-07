package hostfunctions

import (
	"context"
	"fmt"
	"log"

	"github.com/tetratelabs/wazero/api"
)

// LogString : print a string to the console
var LogString = api.GoModuleFunc(func(ctx context.Context, module api.Module, params []uint64) []uint64 {

	//fmt.Println("🌺", params)
	//fmt.Println("🖐 position:", params[0])
	//fmt.Println("🖐 length:", params[1])

	position := uint32(params[0])
	length := uint32(params[1])

	buffer, ok := module.Memory().Read(ctx, position, length)
	if !ok {
		log.Panicf("🟥 Memory.Read(%d, %d) out of range", position, length)
	}
	//fmt.Println("🍭 ", string(buffer))
	fmt.Println(string(buffer))

	return []uint64{0}
})

/* old version
func LogString(ctx context.Context, module api.Module, offset, byteCount uint32) {
	buf, ok := module.Memory().Read(ctx, offset, byteCount)
	if !ok {
		log.Panicf("🟥 Memory.Read(%d, %d) out of range", offset, byteCount)
	}
	fmt.Println(string(buf))
}
*/
