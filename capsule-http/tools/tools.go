// Package tools for the application
package tools

import (
	"net"
	"os"
	"strconv"

	"github.com/bots-garden/capsule-host-sdk/helpers"
)

// TODO:
// Create several function to download the wasm files
// one function per use case
// then chose the appropriate function with a flag

// GetWasmFile load or downloads the wasm file
func GetWasmFile(wasmFilePath, wasmFileURL, authHeaderName, authHeaderValue string) ([]byte, error) {
	//TODO: add authentication with headers
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
	wasmFile, err := helpers.DownloadWasmFile(wasmFileURL, wasmFilePath, authHeaderName, authHeaderValue)
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

// GetEnv returns the environment variable
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetNewHTTPPort returns a unique http port
func GetNewHTTPPort() string {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	httpPort := strconv.Itoa(listener.Addr().(*net.TCPAddr).Port)
	listener.Close()
	return httpPort
}
