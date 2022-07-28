package hostfunctions

import (
	"context"
	"os"

	"github.com/tetratelabs/wazero/api"
)

// string parameter, return string
func WriteFile(ctx context.Context, module api.Module, filePathOffset, filePathByteCount, contentOffset, contentByteCount, retBuffPtrPos, retBuffSize uint32) {

	//=========================================================
	// Read arguments values of the function call
	//=========================================================
	// get string from the wasm module function (from memory)

	filePathStr := ReadStringFromMemory(ctx, module, filePathOffset, filePathByteCount)
	contentStr := ReadStringFromMemory(ctx, module, contentOffset, contentByteCount)

	//fmt.Println("📝 filePathStr:", filePathStr)

	//=========================================================
	// 👋 Implementation: Start
	var stringMessageFromHost = ""
	//dat, err := os.ReadFile(filePathStr)

	dat := []byte(contentStr)
	err := os.WriteFile(filePathStr, dat, 0644)

	if err != nil {
		stringMessageFromHost = CreateStringError(err.Error(), 0)
		// if code 0 don't display code in the error message
	} else {
		stringMessageFromHost = "file created"
	}

	// 👋 Implementation: End
	//=========================================================

	// write the new string stringMessageFromHost to the "shared memory"
	// (host write string result of the function to memory)
	WriteStringToMemory(stringMessageFromHost, ctx, module, retBuffPtrPos, retBuffSize)
}
