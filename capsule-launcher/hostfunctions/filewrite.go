package hostfunctions

import (
	"context"
	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions/memory"
	"github.com/bots-garden/capsule/commons"
	"os"

	"github.com/tetratelabs/wazero/api"
)

var WriteFile = api.GoModuleFunc(func(ctx context.Context, module api.Module, params []uint64) []uint64 {

	positionFilePathName := uint32(params[0])
	lengthFilePathName := uint32(params[1])

	filePath := memory.ReadStringFromMemory(ctx, module, positionFilePathName, lengthFilePathName)

	positionContent := uint32(params[2])
	lengthContent := uint32(params[3])

	content := memory.ReadStringFromMemory(ctx, module, positionContent, lengthContent)

	var stringResultFromHost = ""

	data := []byte(content)
	err := os.WriteFile(filePath, data, 0644)

	if err != nil {
		stringResultFromHost = commons.CreateStringError(err.Error(), 0)
		// if code 0 don't display code in the error message
	} else {
		stringResultFromHost = "file created"
	}

	positionReturnBuffer := uint32(params[4])
	lengthReturnBuffer := uint32(params[5])

	// TODO: I think there is another way (with return, but let's see later with wazero sampleq)
	memory.WriteStringToMemory(stringResultFromHost, ctx, module, positionReturnBuffer, lengthReturnBuffer)

	return []uint64{0}

})

// string parameter, return string
func _WriteFile(ctx context.Context, module api.Module, filePathOffset, filePathByteCount, contentOffset, contentByteCount, retBuffPtrPos, retBuffSize uint32) {

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
