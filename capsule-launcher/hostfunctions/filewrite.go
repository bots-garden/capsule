package hostfunctions

import (
	"context"
	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions/memory"
	"github.com/bots-garden/capsule/commons"
	"os"

	"github.com/tetratelabs/wazero/api"
)

// string parameter, return string
func WriteFile(ctx context.Context, module api.Module, filePathOffset, filePathByteCount, contentOffset, contentByteCount, retBuffPtrPos, retBuffSize uint32) {

	//=========================================================
	// Read arguments values of the function call
	//=========================================================
	// get string from the wasm module function (from memory)

	filePathStr := memory.ReadStringFromMemory(ctx, module, filePathOffset, filePathByteCount)
	contentStr := memory.ReadStringFromMemory(ctx, module, contentOffset, contentByteCount)

	//fmt.Println("📝 filePathStr:", filePathStr)

	//=========================================================
	// 👋 Implementation: Start
	var stringMessageFromHost = ""
	//dat, err := os.ReadFile(filePathStr)

	dat := []byte(contentStr)
	err := os.WriteFile(filePathStr, dat, 0644)

	if err != nil {
		stringMessageFromHost = commons.CreateStringError(err.Error(), 0)
		// if code 0 don't display code in the error message
	} else {
		stringMessageFromHost = "file created"
	}

	// 👋 Implementation: End
	//=========================================================

	// write the new string stringMessageFromHost to the "shared memory"
	// (host write string result of the function to memory)
	memory.WriteStringToMemory(stringMessageFromHost, ctx, module, retBuffPtrPos, retBuffSize)
}
