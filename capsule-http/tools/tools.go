// Package tools for the application
package tools

import (
	"github.com/bots-garden/capsule-host-sdk/helpers"
)

// GetWasmFile load or downloads the wasm file
func GetWasmFile(wasmFilePath, wasmFileURL string) ([]byte, error) {
	//TODO; add authentication with headers
	if len(wasmFileURL) == 0 {
		wasmFile, err := helpers.LoadWasmFile(wasmFilePath)
		/*
		if err != nil {
			log.Println("❌ Error while loading the wasm file:", err)
		} else {
			log.Println("✅ File loaded", wasmFilePath)
		}
		*/
		return wasmFile, err

	}
	wasmFile, err := helpers.DownloadWasmFile(wasmFileURL, wasmFilePath)
	/*
	if err != nil {
		log.Println("❌ Error while downloading the wasm file:", err)
		//os.Exit(1)
	} else {
		log.Println("✅ File downloaded", wasmFilePath)
	}
	*/
	return wasmFile, err
}