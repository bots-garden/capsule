package capsule

import (
	"context"
	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"log"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

// GetPackedPtrPositionAndSize :
// Get the pointer position and a size from result of ExportedFunction("F").Call().
//ex: the string pointer position (in memory) and the length of the string

func GetPackedPtrPositionAndSize(result []uint64) (ptrPos uint32, size uint32) {
	return uint32(result[0] >> 32), uint32(result[0])
}

func CreateWasmRuntime(ctx context.Context) wazero.Runtime {

	//wasmRuntime := wazero.NewRuntime(ctx)
	wasmRuntime := wazero.NewRuntimeWithConfig(ctx, wazero.NewRuntimeConfigInterpreter())

	//https://github.com/tetratelabs/wazero/blob/main/examples/allocation/tinygo/greet.go#L29

	// 🏠 Add host functions to the wasmModule (to be available from the module)
	// These functions allows the module to call functions of the host

	builder := wasmRuntime.NewHostModuleBuilder("env")

	// hostLogString
	builder.NewFunctionBuilder().
		WithGoModuleFunction(hostfunctions.LogString, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{api.ValueTypeI32}).
		Export("hostLogString")

	// hostGetHostInformation TODO: to be translated
	builder.NewFunctionBuilder().
		WithFunc(hostfunctions.GetHostInformation).
		Export("hostGetHostInformation")

	// hostHttp
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.Http,
			[]api.ValueType{
				api.ValueTypeI32, // url position (in memory)
				api.ValueTypeI32, // url length
				api.ValueTypeI32, // method position
				api.ValueTypeI32, // method length
				api.ValueTypeI32, // headers position
				api.ValueTypeI32, // headers length
				api.ValueTypeI32, // body position
				api.ValueTypeI32, // body length
				api.ValueTypeI32, // returnValue position
				api.ValueTypeI32, // returnValue length
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostHttp")

	//_, errEnv := wasmRuntime.NewHostModuleBuilder("env").

	// hostReadFile
	builder.NewFunctionBuilder().
		WithGoModuleFunction(hostfunctions.ReadFile, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{api.ValueTypeI32}).
		Export("hostReadFile")

	// hostWriteFile
	builder.NewFunctionBuilder().
		WithGoModuleFunction(hostfunctions.WriteFile, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{api.ValueTypeI32}).
		Export("hostWriteFile")

	// hostGetEnv
	builder.NewFunctionBuilder().
		WithGoModuleFunction(hostfunctions.GetEnv, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{api.ValueTypeI32}).
		Export("hostGetEnv")

	builder.NewFunctionBuilder().WithFunc(hostfunctions.RedisSet).Export("hostRedisSet").
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
		NewFunctionBuilder().WithFunc(hostfunctions.GetExitCode).Export("hostGetExitCode")

	// hostRequestParamsGet
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.RequestParamsGet,
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostRequestParamsGet")

	_, errBuilder := builder.Instantiate(ctx, wasmRuntime)
	if errBuilder != nil {
		log.Panicln("🔴 Error with env module and host function(s):", errBuilder)
	}

	_, errSnapshot := wasi_snapshot_preview1.Instantiate(ctx, wasmRuntime)
	if errSnapshot != nil {
		log.Panicln("🔴 Error with SnapShot Instantiate:", errSnapshot)
	}

	return wasmRuntime
}

func GetWasmRuntimeAndModuleInstances(wasmFile []byte) (wazero.Runtime, api.Module, context.Context) {
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

func GetModuleInstance(wasmFile []byte) (api.Module, context.Context) {
	// Choose the context to use for function calls.
	ctx := context.Background()

	//wasmRuntime := CreatePersistentWasmRuntime(ctx) // 👋 we must not close the runtime (?)
	wasmRuntime := CreateWasmRuntime(ctx)
	//defer wasmRuntime.Close(ctx) // This closes everything this Runtime created.

	// 🥚 Instantiate the wasm module (from the wasm file)
	// 🖐 The main method is called at this moment
	wasmModule, errInstanceWasmModule := wasmRuntime.InstantiateModuleFromBinary(ctx, wasmFile)

	if errInstanceWasmModule != nil {
		log.Panicln("🔴 Error while creating module instance:", errInstanceWasmModule)
	}
	return wasmModule, ctx
}
