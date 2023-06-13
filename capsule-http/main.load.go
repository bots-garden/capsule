package main

import (
	"context"
	"errors"
	"log"

	"github.com/bots-garden/capsule/capsule-http/handlers"
	"github.com/bots-garden/capsule/capsule-http/tools"
	"github.com/tetratelabs/wazero"
)

// LoadWasmFile -> load the wasm file (or download it)
// the function exits if the wasm file is not present
func LoadWasmFile(ctx context.Context, flags CapsuleFlags, runtime wazero.Runtime) ([]byte, error) {
	if flags.wasm != "" {
		wasmFile, err := tools.GetWasmFile(flags.wasm, flags.url, flags.authHeaderName, flags.authHeaderValue)
		if err != nil {
			log.Println("❌📝 Error while loading the wasm file", err)
			return nil, err
		}
		// Save the wasmFile to memory
		handlers.StoreWasmFile(wasmFile)

		return wasmFile, nil

	}
	// the wasm file is not mandatory in faas mode
	// otherwise it's an error
	if flags.faas != true {
		log.Println("❌📝 Error while loading the wasm file (empty)")
		return nil, errors.New("empty wasm file")
	}
	// when flags.faas is true, the wasm file is not mandatory
	return nil, nil
}
