package hostfunctions

import (
	"context"
	"log"

	//"github.com/go-resty/resty/v2"
	"github.com/tetratelabs/wazero/api"
)

func Http(ctx context.Context, module api.Module,
	urlOffset, urlByteCount, methodOffSet, methodByteCount, headersOffSet, headersByteCount, bodyOffSet, bodyByteCount,
	retBuffPtrPos, retBuffSize uint32) {
	// get url string from the wasm module function (from memory)
	urlBuff, ok := module.Memory().Read(ctx, urlOffset, urlByteCount)
	if !ok {
		log.Panicf("🟥 Memory.Read(%d, %d) out of range", urlOffset, urlByteCount)
	}
	urlStr := string(urlBuff)

	// get method string from the wasm module function (from memory)
	methodBuff, ok := module.Memory().Read(ctx, methodOffSet, methodByteCount)
	if !ok {
		log.Panicf("🟥 Memory.Read(%d, %d) out of range", methodOffSet, methodByteCount)
	}
	methodStr := string(methodBuff)

	// get headers string from the wasm module function (from memory)
	headersBuff, ok := module.Memory().Read(ctx, headersOffSet, headersByteCount)
	if !ok {
		log.Panicf("🟥 Memory.Read(%d, %d) out of range", headersOffSet, headersByteCount)
	}
	headersStr := string(headersBuff)

	// headers => strings.Join(headers[:], "|")

	// get body string from the wasm module function (from memory)
	bodyBuff, ok := module.Memory().Read(ctx, bodyOffSet, bodyByteCount)
	if !ok {
		log.Panicf("🟥 Memory.Read(%d, %d) out of range", bodyOffSet, bodyByteCount)
	}
	bodyStr := string(bodyBuff)

	var stringMessageFromHost = ""
	// 👋 Implementation: Start
	//client := resty.New()
	switch what := methodStr; what {
	case "GET":
		//resp, err := client.R().EnableTrace().Get(urlStr)
		stringMessageFromHost = "🌍 (GET)http: " + urlStr + " method: " + methodStr + " headers: " + headersStr + " body: " + bodyStr
	case "POST":
		stringMessageFromHost = "🌍 (POST)http: " + urlStr + " method: " + methodStr + " headers: " + headersStr + " body: " + bodyStr

	default:
		stringMessageFromHost = "[ERR]🔴 not implemented🚧 wip"
	}

	// 👋 Implementation: End

	// TODO: helper of this:
	// write the new string Result to the "shared memory"
	lengthOfTheMessage := len(stringMessageFromHost)
	results, err := module.ExportedFunction("allocateBuffer").Call(ctx, uint64(lengthOfTheMessage))
	if err != nil {
		log.Panicln(err)
	}

	retOffset := uint32(results[0])
	module.Memory().WriteUint32Le(ctx, retBuffPtrPos, retOffset)
	module.Memory().WriteUint32Le(ctx, retBuffSize, uint32(lengthOfTheMessage))

	// add the message to the memory of the module
	module.Memory().Write(ctx, retOffset, []byte(stringMessageFromHost))

}
