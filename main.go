package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	capsulecli "github.com/bots-garden/capsule/services/cli"
	capsulehttp "github.com/bots-garden/capsule/services/http"

	host_functions "github.com/bots-garden/capsule/host_functions"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/wasi_snapshot_preview1"
)

type CapsuleFlags struct {
	mode  string
	param string
	wasm string
	httpPort int
}

func main() {

	//flags
	/*
	go run main.go \
		-wasm=./wasm_modules/capsule-function-template/hello.wasm \
		-mode=cli \
		-param="👋 hello world 🌍"

	go run main.go \
		-wasm=./wasm_modules/capsule-function-template/hello.wasm \
		-mode=http \
		-httpPort=7070
	*/
	capsuleModePtr := flag.String("mode", "http", "default mode is http else: cli")
	stringParamPtr := flag.String("param", "hello world", "string parameter for the cli mode")
	wasmFilePathPtr := flag.String("wasm", "", "wasm module file path")
	httpPortPtr := flag.Int("httpPort", 8080, "http port")

	flag.Parse()

	flags := CapsuleFlags{
		*capsuleModePtr,
		*stringParamPtr,
		*wasmFilePathPtr,
		*httpPortPtr,
	}
	//fmt.Println(flags)

	//argsWithProg := os.Args
	//wasmModuleFilePath := os.Args[1:][0]
	wasmModuleFilePath := flags.wasm

	// Choose the context to use for function calls.
	ctx := context.Background()

	// Create a new WebAssembly Runtime.
	wasmRuntime := wazero.NewRuntimeWithConfig(wazero.NewRuntimeConfig().WithWasmCore2())
	defer wasmRuntime.Close(ctx) // This closes everything this Runtime created.

	// 🏠 Add host functions
	_, errEnv := wasmRuntime.NewModuleBuilder("env").
		ExportFunction("hostLogString", host_functions.LogString).
		ExportFunction("hostGetHostInformation", host_functions.GetHostInformation).
		ExportFunction("hostPing", host_functions.Ping).
		Instantiate(ctx, wasmRuntime)

	if errEnv != nil {
		log.Panicln("🔴 Error with env module and host function(s):", errEnv)
	}

	_, errInstantiate := wasi_snapshot_preview1.Instantiate(ctx, wasmRuntime)
	if errInstantiate != nil {
		log.Panicln("🔴 Error with Instantiate:", errInstantiate)
	}

	// 📂 Load from file and then Instantiate a WebAssembly module
	wasmFile, errLoadWasmFile := os.ReadFile(wasmModuleFilePath)

	if errLoadWasmFile != nil {
		log.Panicln("🔴 Error while loading the wasm file:", errLoadWasmFile)
	}

	// 🥚 Instantiate the wasm module (from the wasm file)
	wasmModule, errInstanceWasmModule := wasmRuntime.InstantiateModuleFromBinary(ctx, wasmFile)
	if errInstanceWasmModule != nil {
		log.Panicln("🔴 Error while creating module instance:", errInstanceWasmModule)
	}

	switch what := flags.mode; what {
	case "http":
		fmt.Println("[http mode] 🚧 in progress", flags.param)
	case "cli":
		//fmt.Println("[cli mode] 🚧 in progress", flags.param)
		capsulecli.Execute(flags.param, wasmModule, ctx)
	default:
		log.Panicln("🔴 bad mode", *capsuleModePtr)
	}
	
}
