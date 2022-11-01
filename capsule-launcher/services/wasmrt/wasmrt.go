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

	wasmRuntime := wazero.NewRuntime(ctx)
	//https://github.com/tetratelabs/wazero/blob/main/examples/allocation/tinygo/greet.go#L29

	// 🏠 Add host functions to the wasmModule (to be available from the module)
	// These functions allows the module to call functions of the host
	_, errEnv := wasmRuntime.NewHostModuleBuilder("env").
		NewFunctionBuilder().
		WithGoModuleFunction(hostfunctions.LogString, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{api.ValueTypeI32}).
		Export("hostLogString").
		NewFunctionBuilder().WithFunc(hostfunctions.GetHostInformation).Export("hostGetHostInformation").
		NewFunctionBuilder().WithFunc(hostfunctions.Ping).Export("hostPing").
		NewFunctionBuilder().WithFunc(hostfunctions.Http).Export("hostHttp").
		// hostReadFile
		NewFunctionBuilder().
		WithGoModuleFunction(hostfunctions.ReadFile, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{api.ValueTypeI32}).
		Export("hostReadFile").
		// hostWriteFile
		NewFunctionBuilder().
		WithGoModuleFunction(hostfunctions.WriteFile, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{api.ValueTypeI32}).
		Export("hostWriteFile").
		// hostGetEnv
		NewFunctionBuilder().
		WithGoModuleFunction(hostfunctions.GetEnv, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{api.ValueTypeI32}).
		Export("hostGetEnv").
		NewFunctionBuilder().WithFunc(hostfunctions.RedisSet).Export("hostRedisSet").
		NewFunctionBuilder().WithFunc(hostfunctions.RedisGet).Export("hostRedisGet").
		NewFunctionBuilder().WithFunc(hostfunctions.RedisKeys).Export("hostRedisKeys").
		NewFunctionBuilder().WithFunc(hostfunctions.MemorySet).Export("hostMemorySet").
		NewFunctionBuilder().WithFunc(hostfunctions.MemoryGet).Export("hostMemoryGet").
		NewFunctionBuilder().WithFunc(hostfunctions.MemoryKeys).Export("hostMemoryKeys").
		NewFunctionBuilder().WithFunc(hostfunctions.CouchBaseQuery).Export("hostCouchBaseQuery").
		NewFunctionBuilder().WithFunc(hostfunctions.NatsPublish).Export("hostNatsPublish").
		NewFunctionBuilder().WithFunc(hostfunctions.NatsConnectPublish).Export("hostNatsConnectPublish").
		NewFunctionBuilder().WithFunc(hostfunctions.NatsGetSubject).Export("hostNatsGetSubject").
		NewFunctionBuilder().WithFunc(hostfunctions.NatsGetServer).Export("hostNatsGetServer").
		NewFunctionBuilder().WithFunc(hostfunctions.NatsConnectRequest).Export("hostNatsConnectRequest").
		NewFunctionBuilder().WithFunc(hostfunctions.NatsReply).Export("hostNatsReply").
		NewFunctionBuilder().WithFunc(hostfunctions.MqttGetTopic).Export("hostMqttGetTopic").
		NewFunctionBuilder().WithFunc(hostfunctions.MqttGetServer).Export("hostMqttGetServer").
		NewFunctionBuilder().WithFunc(hostfunctions.MqttGetClientId).Export("hostMqttGetClientId").
		NewFunctionBuilder().WithFunc(hostfunctions.MqttPublish).Export("hostMqttPublish").
		NewFunctionBuilder().WithFunc(hostfunctions.MqttConnectPublish).Export("hostMqttConnectPublish").
		NewFunctionBuilder().WithFunc(hostfunctions.GetExitError).Export("hostGetExitError").
		NewFunctionBuilder().WithFunc(hostfunctions.GetExitCode).Export("hostGetExitCode").
		NewFunctionBuilder().WithFunc(hostfunctions.RequestParamsGet).Export("hostRequestParamsGet").
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

var persistentWasmRuntime wazero.Runtime

func CreatePersistentWasmRuntime(ctx context.Context) wazero.Runtime {
	if persistentWasmRuntime == nil {
		return CreateWasmRuntime(ctx)
	} else {
		return persistentWasmRuntime
	}
}

func GetModuleInstance(wasmFile []byte) (api.Module, context.Context) {
	// Choose the context to use for function calls.
	ctx := context.Background()

	wasmRuntime := CreatePersistentWasmRuntime(ctx) // 👋 we must not close the runtime (?)
	//defer wasmRuntime.Close(ctx) // This closes everything this Runtime created.

	// 🥚 Instantiate the wasm module (from the wasm file)
	// 🖐 The main method is called at this moment
	wasmModule, errInstanceWasmModule := wasmRuntime.InstantiateModuleFromBinary(ctx, wasmFile)

	if errInstanceWasmModule != nil {
		log.Panicln("🔴 Error while creating module instance:", errInstanceWasmModule)
	}
	return wasmModule, ctx
}
