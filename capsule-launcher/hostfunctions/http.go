package hostfunctions

import (
    "context"
    "github.com/bots-garden/capsule/capsule-launcher/hostfunctions/memory"
    "github.com/go-resty/resty/v2"
    "github.com/tetratelabs/wazero/api"

    "github.com/bots-garden/capsule/commons"
)

var Http = api.GoModuleFunc(func(ctx context.Context, module api.Module, params []uint64) []uint64 {

    // get url string from the wasm module function (from memory)
    positionUrl := uint32(params[0])
    lengthUrl := uint32(params[1])
    urlStr := memory.ReadStringFromMemory(ctx, module, positionUrl, lengthUrl)

    // get method string from the wasm module function (from memory)
    positionMethod := uint32(params[2])
    lengthMethod := uint32(params[3])
    methodStr := memory.ReadStringFromMemory(ctx, module, positionMethod, lengthMethod)

    // get headers string from the wasm module function (from memory)
    positionHeaders := uint32(params[4])
    lengthHeaders := uint32(params[5])
    // 🖐 headers => Accept:application/json|Content-Type: text/html; charset=UTF-8
    headersStr := memory.ReadStringFromMemory(ctx, module, positionHeaders, lengthHeaders)
    headersSlice := commons.CreateSliceFromString(headersStr, commons.StrSeparator)
    headersMap := commons.CreateMapFromSlice(headersSlice, commons.FieldSeparator)

    // get body string from the wasm module function (from memory)
    positionBody := uint32(params[6])
    lengthBody := uint32(params[7])
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

    positionReturnBuffer := uint32(params[8])
    lengthReturnBuffer := uint32(params[9])

    // write the new string stringMessageFromHost to the "shared memory"
    // (host write string result of the function to memory)
    memory.WriteStringToMemory(stringResultFromHost, ctx, module, positionReturnBuffer, lengthReturnBuffer)

    return []uint64{0}
})

func _Http(ctx context.Context, module api.Module,
    urlOffset, urlByteCount, methodOffSet, methodByteCount, headersOffSet, headersByteCount, bodyOffSet, bodyByteCount,
    retBuffPtrPos, retBuffSize uint32) {

    //=========================================================
    // Read arguments values of the function call
    //=========================================================

    // get url string from the wasm module function (from memory)
    urlStr := memory.ReadStringFromMemory(ctx, module, urlOffset, urlByteCount)

    // get method string from the wasm module function (from memory)
    methodStr := memory.ReadStringFromMemory(ctx, module, methodOffSet, methodByteCount)

    // get headers string from the wasm module function (from memory)
    // 🖐 headers => Accept:application/json|Content-Type: text/html; charset=UTF-8
    headersStr := memory.ReadStringFromMemory(ctx, module, headersOffSet, headersByteCount)

    //TODO: choose another separator: °
    headersSlice := commons.CreateSliceFromString(headersStr, commons.StrSeparator)

    //fmt.Println(headersSlice)

    headersMap := commons.CreateMapFromSlice(headersSlice, commons.FieldSeparator)

    //fmt.Println(headersMap)
    //fmt.Println(headersMap["Accept"])
    //fmt.Println(headersMap["Content-Type"])

    // get body string from the wasm module function (from memory)
    bodyStr := memory.ReadStringFromMemory(ctx, module, bodyOffSet, bodyByteCount)
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
    memory.WriteStringToMemory(stringMessageFromHost, ctx, module, retBuffPtrPos, retBuffSize)

}
