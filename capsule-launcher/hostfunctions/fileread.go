package hostfunctions

import (
	"context"
	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions/memory"
	"os"

	"github.com/bots-garden/capsule/commons"

	"github.com/tetratelabs/wazero/api"
)

// ReadFile : string parameter, return string
var ReadFile = api.GoModuleFunc(func(ctx context.Context, module api.Module, stack []uint64) {

	positionFilePathName := uint32(stack[0])
	lengthFilePathName := uint32(stack[1])

	filePath := memory.ReadStringFromMemory(ctx, module, positionFilePathName, lengthFilePathName)

	var stringResultFromHost = ""

	data, err := os.ReadFile(filePath)
	if err != nil {
		stringResultFromHost = commons.CreateStringError(err.Error(), 0)
		// if code 0 don't display code in the error message
	} else {
		stringResultFromHost = string(data)
	}

	positionReturnBuffer := uint32(stack[2])
	lengthReturnBuffer := uint32(stack[3])

	// TODO: I think there is another way (with return, but let's see later with wazero sampleq)
	memory.WriteStringToMemory(stringResultFromHost, ctx, module, positionReturnBuffer, lengthReturnBuffer)

	stack[0] = 0 // return 0

})

/* old version
func ReadFile(ctx context.Context, module api.Module, fileOffset, fileByteCount, retBuffPtrPos, retBuffSize uint32) {

	//=========================================================
	// Read arguments values of the function call
	//=========================================================
	// get string from the wasm module function (from memory)

	filePathStr := memory.ReadStringFromMemory(ctx, module, fileOffset, fileByteCount)

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
	memory.WriteStringToMemory(stringMessageFromHost, ctx, module, retBuffPtrPos, retBuffSize)
}
*/
