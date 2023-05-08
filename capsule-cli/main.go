// Package main, the next generation of the Capsule project
package main

import (
	"context"
	"flag"
	"fmt"

	"log"
	"os"

	"github.com/bots-garden/capsule-host-sdk"
	"github.com/bots-garden/capsule/capsule-http/tools"
)

// CapsuleFlags handles params for the capsule-http command
type CapsuleFlags struct {
	wasm     string // wasm file location
	url      string // to download the wasm file
	registry string // url to the registry
	version  bool
	params   string
}

func main() {
	version := "v0.3.5 🍓 [strawberry]"
	args := os.Args[1:]

	if len(args) == 0 {
		log.Println("Capsule needs some args to start.")
		os.Exit(0)
	}
	// Capsule flags
	wasmFilePathPtr := flag.String("wasm", "", "wasm module file path")
	wasmFileURLPtr := flag.String("url", "", "url for downloading wasm module file")
	registryPtr := flag.String("registry", "", "url of the wasm registry")
	paramsPtr := flag.String("params", "", "parameter(s) for the capsule CLI")
	versionPtr := flag.Bool("version", false, "prints capsule CLI current version")

	flag.Parse()

	if *versionPtr {
		fmt.Println(version)
		os.Exit(0)
	}

	flags := CapsuleFlags{
		*wasmFilePathPtr,
		*wasmFileURLPtr,
		*registryPtr,
		*versionPtr,
		*paramsPtr,
	}

	// Choose the context to use for function calls.
	ctx := context.Background()

	// Create a new WebAssembly Runtime.
	runtime := capsule.GetRuntime(ctx)

	// START: host functions
	// Get the builder and load the default host functions
	builder := capsule.GetBuilder(runtime)

	// Add your host functions here
	// 🏠
	// End of of you hostfunction

	// Instantiate builder and default host functions
	_, err := builder.Instantiate(ctx)
	if err != nil {
		log.Println("❌ Error with env module and host function(s):", err)
		os.Exit(1)
	}
	// END: host functions

	// This closes everything this Runtime created.
	defer runtime.Close(ctx)

	// Load the WebAssembly module
	wasmFile, err := tools.GetWasmFile(flags.wasm, flags.url)
	if err != nil {
		log.Println("❌ Error while loading the wasm file", err)
		os.Exit(1)
	}

	mod, err := runtime.Instantiate(ctx, wasmFile)
	if err != nil {
		log.Println("❌ Error with the module instance", err)
		os.Exit(1)
	}

	// Get the reference to the WebAssembly function: "callHandle"
	// callHandle is exported by the Capsule plugin
	handleFunction := capsule.GetHandle(mod)

	// send parameter to the function
	pos, size, err := capsule.CopyDataToMemory(ctx, mod, []byte(flags.params))
	if err != nil {
		log.Println("❌ Error when copying data to memory", err)
		os.Exit(1)
	}

	// Now, we can call "callHandle" with the position and the size of "Bob Morane"
	// the result type is []uint64
	result, err := handleFunction.Call(ctx, pos, size)
	if err != nil {
		log.Println("❌ Error when calling callHandle", err)
		os.Exit(1)
	}
	// read the result of the function
	rpos, rsize := capsule.UnPackPosSize(result[0])

	bRes, err := capsule.ReadDataFromMemory(mod, rpos, rsize)
	if err != nil {
		log.Println("❌ Error when reading the memory", err)
		os.Exit(1)
	}

	res, err := capsule.Result(bRes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(string(res))
		os.Exit(0)
	}

}
