package capsule

import (
	"context"
	"log"

	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"

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

	// hostGetHostInformation
	builder.NewFunctionBuilder().
		WithGoModuleFunction(hostfunctions.GetHostInformation,
			[]api.ValueType{
				api.ValueTypeI32, // returnValue position
				api.ValueTypeI32, // returnValue length
			},
			[]api.ValueType{api.ValueTypeI32}).
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
		WithGoModuleFunction(
			hostfunctions.ReadFile,
			[]api.ValueType{
				api.ValueTypeI32, // positionFilePathName
				api.ValueTypeI32, // lengthFilePathName
				api.ValueTypeI32, // positionReturnBuffer
				api.ValueTypeI32, // lengthReturnBuffer
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostReadFile")

	// hostWriteFile
	builder.NewFunctionBuilder().
		WithGoModuleFunction(hostfunctions.WriteFile, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{api.ValueTypeI32}).
		Export("hostWriteFile")

	// hostGetEnv
	builder.NewFunctionBuilder().
		WithGoModuleFunction(hostfunctions.GetEnv, []api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32}, []api.ValueType{api.ValueTypeI32}).
		Export("hostGetEnv")

	// hostMemorySet
	// hostMemoryGet
	// hostMemoryKeys
	builder.
		NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.MemorySet,
			[]api.ValueType{
				api.ValueTypeI32, // keyValue position
				api.ValueTypeI32, // keyValue length
				api.ValueTypeI32, // value position
				api.ValueTypeI32, // value length
				api.ValueTypeI32, // returnValue position
				api.ValueTypeI32, // returnValue length
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostMemorySet").
		NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.MemoryGet,
			[]api.ValueType{
				api.ValueTypeI32, // keyValue position
				api.ValueTypeI32, // keyValue length
				api.ValueTypeI32, // returnValue position
				api.ValueTypeI32, // returnValue length
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostMemoryGet").
		NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.MemoryKeys,
			[]api.ValueType{
				api.ValueTypeI32, // returnValue position
				api.ValueTypeI32, // returnValue length
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostMemoryKeys")

	// hostRedisSet
	// hostRedisGet
	// hostRedisKeys
	builder.
		NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.RedisSet,
			[]api.ValueType{
				api.ValueTypeI32, // keyValue position
				api.ValueTypeI32, // keyValue length
				api.ValueTypeI32, // value position
				api.ValueTypeI32, // value length
				api.ValueTypeI32, // returnValue position
				api.ValueTypeI32, // returnValue length
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostRedisSet").
		NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.RedisGet,
			[]api.ValueType{
				api.ValueTypeI32, // keyValue position
				api.ValueTypeI32, // keyValue length
				api.ValueTypeI32, // returnValue position
				api.ValueTypeI32, // returnValue length
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostRedisGet").
		NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.RedisKeys,
			[]api.ValueType{
				api.ValueTypeI32, // pattern position
				api.ValueTypeI32, // pattern length
				api.ValueTypeI32, // returnValue position
				api.ValueTypeI32, // returnValue length
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostRedisKeys")

	builder.
		NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.CouchBaseQuery,
			[]api.ValueType{
				api.ValueTypeI32, // query position
				api.ValueTypeI32, // query length
				api.ValueTypeI32, // returnValue position
				api.ValueTypeI32, // returnValue length
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostCouchBaseQuery")

	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.NatsPublish,
			[]api.ValueType{
				api.ValueTypeI32, // subjectOffset
				api.ValueTypeI32, // subjectByteCount
				api.ValueTypeI32, // dataOffset
				api.ValueTypeI32, // dataByteCount
				api.ValueTypeI32, // retBuffPtrPos
				api.ValueTypeI32, // retBuffSize
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostNatsPublish").
		NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.NatsConnectPublish,
			[]api.ValueType{
				api.ValueTypeI32, // natsSrvOffset
				api.ValueTypeI32, // natsSrvByteCount
				api.ValueTypeI32, // subjectOffset
				api.ValueTypeI32, // subjectByteCount
				api.ValueTypeI32, // dataOffset
				api.ValueTypeI32, // dataByteCount
				api.ValueTypeI32, // retBuffPtrPos
				api.ValueTypeI32, // retBuffSize
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostNatsConnectPublish").
		NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.NatsGetSubject,
			[]api.ValueType{
				api.ValueTypeI32, // retBuffPtrPos
				api.ValueTypeI32, // retBuffSize
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostNatsGetSubject").
		NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.NatsGetServer,
			[]api.ValueType{
				api.ValueTypeI32, // retBuffPtrPos
				api.ValueTypeI32, // retBuffSize
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostNatsGetServer").
		NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.NatsConnectRequest,
			[]api.ValueType{
				api.ValueTypeI32, // natsSrvOffset
				api.ValueTypeI32, // natsSrvByteCount
				api.ValueTypeI32, // subjectOffset
				api.ValueTypeI32, // subjectByteCount
				api.ValueTypeI32, // dataOffset
				api.ValueTypeI32, // dataByteCount
				api.ValueTypeI32, // time duration
				api.ValueTypeI32, // retBuffPtrPos
				api.ValueTypeI32, // retBuffSize
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostNatsConnectRequest").
		NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.NatsReply,
			[]api.ValueType{
				api.ValueTypeI32, // dataOffset
				api.ValueTypeI32, // dataByteCount
				api.ValueTypeI32, // time duration
				api.ValueTypeI32, // retBuffPtrPos
				api.ValueTypeI32, // retBuffSize
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostNatsReply")

	builder.
		NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.MqttGetTopic,
			[]api.ValueType{
				api.ValueTypeI32, // retBuffPtrPos
				api.ValueTypeI32, // retBuffSize
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostMqttGetTopic").
		NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.MqttGetServer,
			[]api.ValueType{
				api.ValueTypeI32, // retBuffPtrPos
				api.ValueTypeI32, // retBuffSize
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostMqttGetServer").
		NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.MqttGetClientId,
			[]api.ValueType{
				api.ValueTypeI32, // retBuffPtrPos
				api.ValueTypeI32, // retBuffSize
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostMqttGetClientId").
		NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.MqttPublish,
			[]api.ValueType{
				api.ValueTypeI32, // topicOffset
				api.ValueTypeI32, // topicByteCount
				api.ValueTypeI32, // dataOffset
				api.ValueTypeI32, // dataByteCount
				api.ValueTypeI32, // retBuffPtrPos
				api.ValueTypeI32, // retBuffSize
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostMqttPublish").
		NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.MqttConnectPublish,
			[]api.ValueType{
				api.ValueTypeI32, // mqttSrvOffset
				api.ValueTypeI32, // mqttSrvByteCount
				api.ValueTypeI32, // clientIdPtrOffset
				api.ValueTypeI32, // clientIdByteCount
				api.ValueTypeI32, // topicOffset
				api.ValueTypeI32, // topicByteCount
				api.ValueTypeI32, // dataOffset
				api.ValueTypeI32, // dataByteCount
				api.ValueTypeI32, // retBuffPtrPos
				api.ValueTypeI32, // retBuffSize
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostMqttConnectPublish")

	builder.
		NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.GetExitError,
			[]api.ValueType{
				api.ValueTypeI32, // retBuffPtrPos
				api.ValueTypeI32, // retBuffSize
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostGetExitError").
		NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.GetExitCode,
			[]api.ValueType{
				api.ValueTypeI32, // retBuffPtrPos
				api.ValueTypeI32, // retBuffSize
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostGetExitCode")

	// hostRequestParamsGet
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			hostfunctions.RequestParamsGet,
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32}).
		Export("hostRequestParamsGet")

	_, errBuilder := builder.Instantiate(ctx)
	if errBuilder != nil {
		log.Panicln("🔴 Error with env module and host function(s):", errBuilder)
	}

	// This is only for TinyGo
	wasi_snapshot_preview1.MustInstantiate(ctx, wasmRuntime)
	/*
	       _, errSnapshot := wasi_snapshot_preview1.Instantiate(ctx, wasmRuntime)
	   	if errSnapshot != nil {
	   		log.Panicln("🔴 Error with SnapShot Instantiate:", errSnapshot)
	   	}
	*/

	return wasmRuntime
}

func GetWasmRuntimeAndModuleInstances(wasmFile []byte) (wazero.Runtime, api.Module, context.Context) {
	// Choose the context to use for function calls.
	ctx := context.Background()

	wasmRuntime := CreateWasmRuntime(ctx)
	//defer wasmRuntime.Close(ctx) // This closes everything this Runtime created.

	// 🥚 Instantiate the wasm module (from the wasm file)
	// 🖐 The main method is called at this moment
	wasmModule, errInstanceWasmModule := wasmRuntime.Instantiate(ctx, wasmFile)

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
	wasmModule, errInstanceWasmModule := wasmRuntime.Instantiate(ctx, wasmFile)

	if errInstanceWasmModule != nil {
		log.Panicln("🔴 Error while creating module instance:", errInstanceWasmModule)
	}
	return wasmModule, ctx
}
