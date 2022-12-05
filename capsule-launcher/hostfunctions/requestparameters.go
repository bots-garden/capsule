package hostfunctions

import (
	"context"
	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions/memory"
	"github.com/bots-garden/capsule/commons"
	"github.com/tetratelabs/wazero/api"
)

// RequestParamsGet gets from the host memory the values related to an HTTP request
// And then writes a string which contains: reqParams.JsonData, reqParams.Headers, reqParams.Uri, reqParams.Method

var RequestParamsGet = api.GoModuleFunc(func(ctx context.Context, module api.Module, stack []uint64) {
	reqId := uint32(stack[0])
	reqParams, err := GetRequestParams(reqId)
	// This variable will store the concatenation of reqParams.JsonData, reqParams.Headers, reqParams.Uri, reqParams.Method
	var stringResultFromHost = ""

	if err != nil {
		stringResultFromHost = commons.CreateStringError("key (requestId) does not exist", 0)
		// if code 0 don't display code in the error message
	} else {
		stringResultFromHost = commons.CreateStringFromSlice([]string{reqParams.JsonData, reqParams.Headers, reqParams.Uri, reqParams.Method}, commons.StrSeparator)
	}

	positionReturnBuffer := uint32(stack[1])
	lengthReturnBuffer := uint32(stack[2])

	// TODO: I think there is another way (with return, but let's see later with wazero sampleq)
	// Write the new string stringResultFromHost to the "shared memory"
	memory.WriteStringToMemory(stringResultFromHost, ctx, module, positionReturnBuffer, lengthReturnBuffer)

	stack[0] = 0 // return 0
})

/* old version
func _RequestParamsGet(ctx context.Context, module api.Module, reqId, retBuffPtrPos, retBuffSize uint32) {

	// reqIdOffset: it's a position
	// reqIdByteCount: it's a size

	//keyStr := memory.ReadStringFromMemory(ctx, module, reqIdOffset, reqIdByteCount)

	reqParams, err := GetRequestParams(reqId)

	//fmt.Println("🤖🔵[READ reqParams]", reqParams, reqId)

	// This variable will store the concatenation of reqParams.JsonData, reqParams.Headers, reqParams.Uri, reqParams.Method
	var stringResultFromHost = ""

	if err != nil {
		stringResultFromHost = commons.CreateStringError("key (requestId) does not exist", 0)
		// if code 0 don't display code in the error message
	} else {
		stringResultFromHost = commons.CreateStringFromSlice([]string{reqParams.JsonData, reqParams.Headers, reqParams.Uri, reqParams.Method}, commons.StrSeparator)
	}

	//fmt.Println("🤖🔵[Memory stringResultFromHost]", stringResultFromHost)

	// Write the new string stringResultFromHost to the "shared memory"
	memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)
}
*/
