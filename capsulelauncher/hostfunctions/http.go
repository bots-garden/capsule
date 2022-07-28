package hostfunctions

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/tetratelabs/wazero/api"

	"github.com/bots-garden/capsule/capsulelauncher/commons"
)

func Http(ctx context.Context, module api.Module,
	urlOffset, urlByteCount, methodOffSet, methodByteCount, headersOffSet, headersByteCount, bodyOffSet, bodyByteCount,
	retBuffPtrPos, retBuffSize uint32) {

	//=========================================================
	// Read arguments values of the function call
	//=========================================================

	// get url string from the wasm module function (from memory)
	urlStr := ReadStringFromMemory(ctx, module, urlOffset, urlByteCount)

	// get method string from the wasm module function (from memory)
	methodStr := ReadStringFromMemory(ctx, module, methodOffSet, methodByteCount)

	// get headers string from the wasm module function (from memory)
	// 🖐 headers => Accept:application/json|Content-Type: text/html; charset=UTF-8
	headersStr := ReadStringFromMemory(ctx, module, headersOffSet, headersByteCount)

	//TODO: choose another separator: °
	headersSlice := CreateSliceFromString(headersStr, "|")

	//fmt.Println(headersSlice)

	headersMap := CreateMapFromSlice(headersSlice, ":")

	//fmt.Println(headersMap)
	//fmt.Println(headersMap["Accept"])
	//fmt.Println(headersMap["Content-Type"])

	// get body string from the wasm module function (from memory)
	bodyStr := ReadStringFromMemory(ctx, module, bodyOffSet, bodyByteCount)
	//fmt.Println("==>", bodyStr)

	//=========================================================================
	// 👋 Implementation: Start
	var stringMessageFromHost = ""
	client := resty.New()

	for key, value := range headersMap {
		client.SetHeader(key, value)
	}

	switch what := methodStr; what {
	case "GET":

		resp, err := client.R().EnableTrace().Get(urlStr)
		if err != nil {
			stringMessageFromHost = commons.CreateStringError(err.Error(), 0)
			// if code 0 don't display code in the error message
		} else {
			stringMessageFromHost = resp.String()
		}

	case "POST":

		resp, err := client.R().EnableTrace().SetBody(bodyStr).Post(urlStr)
		if err != nil {
			stringMessageFromHost = commons.CreateStringError(err.Error(), 0)
			// if code 0 don't display code in the error message
		} else {
			stringMessageFromHost = resp.String()
		}

		//stringMessageFromHost = "🌍 (POST)http: " + urlStr + " method: " + methodStr + " headers: " + headersStr + " body: " + bodyStr

	default:
		stringMessageFromHost = commons.CreateStringError("🔴"+methodStr+" is not yet implemented: 🚧 wip", 0)
	}
	// 👋 Implementation: End
	//=========================================================================

	// write the new string stringMessageFromHost to the "shared memory"
	// (host write string result of the funcyion to memory)
	WriteStringToMemory(stringMessageFromHost, ctx, module, retBuffPtrPos, retBuffSize)

}
