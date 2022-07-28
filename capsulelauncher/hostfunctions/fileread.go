package hostfunctions

import (
	"context"
	"os"

	"github.com/bots-garden/capsule/capsulelauncher/commons"

	"github.com/tetratelabs/wazero/api"
)

// ReadFile : string parameter, return string
func ReadFile(ctx context.Context, module api.Module, fileOffset, fileByteCount, retBuffPtrPos, retBuffSize uint32) {

	//=========================================================
	// Read arguments values of the function call
	//=========================================================
	// get string from the wasm module function (from memory)

	filePathStr := ReadStringFromMemory(ctx, module, fileOffset, fileByteCount)

	//fmt.Println("📝 filePathStr:", filePathStr)

	//=========================================================
	// 👋 Implementation: Start
	var stringMessageFromHost = ""
	dat, err := os.ReadFile(filePathStr)
	if err != nil {
		stringMessageFromHost = commons.CreateStringError(err.Error(), 0)
		// if code 0 don't display code in the error message
	} else {
		stringMessageFromHost = string(dat)
	}

	// 👋 Implementation: End
	//=========================================================

	// write the new string stringMessageFromHost to the "shared memory"
	// (host write string result of the function to memory)
	WriteStringToMemory(stringMessageFromHost, ctx, module, retBuffPtrPos, retBuffSize)
}
