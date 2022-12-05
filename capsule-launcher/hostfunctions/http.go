package hostfunctions

import (
	"context"
	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions/memory"
	"github.com/go-resty/resty/v2"
	"github.com/tetratelabs/wazero/api"

	"github.com/bots-garden/capsule/commons"
)

var Http = api.GoModuleFunc(func(ctx context.Context, module api.Module, stack []uint64) {

	// get url string from the wasm module function (from memory)
	positionUrl := uint32(stack[0])
	lengthUrl := uint32(stack[1])
	urlStr := memory.ReadStringFromMemory(ctx, module, positionUrl, lengthUrl)

	// get method string from the wasm module function (from memory)
	positionMethod := uint32(stack[2])
	lengthMethod := uint32(stack[3])
	methodStr := memory.ReadStringFromMemory(ctx, module, positionMethod, lengthMethod)

	// get headers string from the wasm module function (from memory)
	positionHeaders := uint32(stack[4])
	lengthHeaders := uint32(stack[5])
	// 🖐 headers => Accept:application/json|Content-Type: text/html; charset=UTF-8
	headersStr := memory.ReadStringFromMemory(ctx, module, positionHeaders, lengthHeaders)
	headersSlice := commons.CreateSliceFromString(headersStr, commons.StrSeparator)
	headersMap := commons.CreateMapFromSlice(headersSlice, commons.FieldSeparator)

	// get body string from the wasm module function (from memory)
	positionBody := uint32(stack[6])
	lengthBody := uint32(stack[7])
	bodyStr := memory.ReadStringFromMemory(ctx, module, positionBody, lengthBody)

	var stringResultFromHost = ""
	client := resty.New()

	for key, value := range headersMap {
		client.SetHeader(key, value)
	}

	switch what := methodStr; what {
	case "GET":

		resp, err := client.R().EnableTrace().Get(urlStr)
		if err != nil {
			stringResultFromHost = commons.CreateStringError(err.Error(), 0)
			// if code 0 don't display code in the error message
		} else {
			stringResultFromHost = resp.String()
		}

	case "POST":

		resp, err := client.R().EnableTrace().SetBody(bodyStr).Post(urlStr)
		if err != nil {
			stringResultFromHost = commons.CreateStringError(err.Error(), 0)
			// if code 0 don't display code in the error message
		} else {
			stringResultFromHost = resp.String()
		}

	default:
		stringResultFromHost = commons.CreateStringError("🔴"+methodStr+" is not yet implemented: 🚧 wip", 0)
	}

	positionReturnBuffer := uint32(stack[8])
	lengthReturnBuffer := uint32(stack[9])

	// write the new string stringMessageFromHost to the "shared memory"
	// (host write string result of the function to memory)
	memory.WriteStringToMemory(stringResultFromHost, ctx, module, positionReturnBuffer, lengthReturnBuffer)

	stack[0] = 0 // return 0
})
