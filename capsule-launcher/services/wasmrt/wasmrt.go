package capsule

import (
	"context"
	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions"
	"log"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

// GetPackedPtrPositionAndSize :
// Get the pointer position and a size from result of ExportedFunction("F").Call().
//ex: the string pointer position (in memory) and the length of the string

func GetPackedPtrPositionAndSize(result []uint64) (ptrPos uint32, size uint32) {
	return uint32(result[0] >> 32), uint32(result[0])
}

func CreateWasmRuntime(ctx context.Context) wazero.Runtime {

	wasmRuntime := wazero.NewRuntimeWithConfig(ctx, wazero.NewRuntimeConfig().WithWasmCore2())
	//https://github.com/tetratelabs/wazero/blob/main/examples/allocation/tinygo/greet.go#L29

	// 🏠 Add host functions to the wasmModule (to be availale from the module)
	// These functions allows the module to call functions of the host
	_, errEnv := wasmRuntime.NewModuleBuilder("env").
		ExportFunction("hostLogString", hostfunctions.LogString).
		ExportFunction("hostGetHostInformation", hostfunctions.GetHostInformation).
		ExportFunction("hostPing", hostfunctions.Ping).
		ExportFunction("hostHttp", hostfunctions.Http).
		ExportFunction("hostReadFile", hostfunctions.ReadFile).
		ExportFunction("hostWriteFile", hostfunctions.WriteFile).
		ExportFunction("hostGetEnv", hostfunctions.GetEnv).
		ExportFunction("hostRedisSet", hostfunctions.RedisSet).
		ExportFunction("hostRedisGet", hostfunctions.RedisGet).
		ExportFunction("hostRedisKeys", hostfunctions.RedisKeys).
		ExportFunction("hostMemorySet", hostfunctions.MemorySet).
		ExportFunction("hostMemoryGet", hostfunctions.MemoryGet).
		ExportFunction("hostMemoryKeys", hostfunctions.MemoryKeys).
		ExportFunction("hostCouchBaseQuery", hostfunctions.CouchBaseQuery).
		ExportFunction("hostNatsPublish", hostfunctions.NatsPublish).
		ExportFunction("hostNatsConnectPublish", hostfunctions.NatsConnectPublish).
		ExportFunction("hostNatsGetSubject", hostfunctions.NatsGetSubject).
		ExportFunction("hostNatsGetServer", hostfunctions.NatsGetServer).
		ExportFunction("hostGetExitError", hostfunctions.GetExitError).
		ExportFunction("hostGetExitCode", hostfunctions.GetExitCode).
		Instantiate(ctx, wasmRuntime)

	if errEnv != nil {
		log.Panicln("🔴 Error with env module and host function(s):", errEnv)
	}

	_, errInstantiate := wasi_snapshot_preview1.Instantiate(ctx, wasmRuntime)
	if errInstantiate != nil {
		log.Panicln("🔴 Error with Instantiate:", errInstantiate)
	}

	return wasmRuntime
}

func CreateWasmRuntimeAndModuleInstances(wasmFile []byte) (wazero.Runtime, api.Module, context.Context) {
	// Choose the context to use for function calls.
	ctx := context.Background()

	wasmRuntime := CreateWasmRuntime(ctx)
	//defer wasmRuntime.Close(ctx) // This closes everything this Runtime created.

	// 🥚 Instantiate the wasm module (from the wasm file)
	// 🖐 The main method is called at this moment
	wasmModule, errInstanceWasmModule := wasmRuntime.InstantiateModuleFromBinary(ctx, wasmFile)

	if errInstanceWasmModule != nil {
		log.Panicln("🔴 Error while creating module instance:", errInstanceWasmModule)
	}
	return wasmRuntime, wasmModule, ctx
}
