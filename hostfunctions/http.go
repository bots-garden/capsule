package hostfunctions

import (
	"context"
	"log"

	"github.com/tetratelabs/wazero/api"
)

func Http(ctx context.Context, module api.Module,
  urlOffset, urlByteCount, methodOffSet, methodByteCount, headersOffSet, headersByteCount, bodyOffSet, bodyByteCount,
  retBufPtrPos, retBufSize uint32) {
	// get url string from the wasm module function (from memory)
	urlBuf, ok := module.Memory().Read(ctx, urlOffset, urlByteCount)
	if !ok {
		log.Panicf("🟥 Memory.Read(%d, %d) out of range", urlOffset, urlByteCount)
	}
	urlStr := string(urlBuf)

	// get method string from the wasm module function (from memory)
	methodBuf, ok := module.Memory().Read(ctx, methodOffSet, methodByteCount)
	if !ok {
		log.Panicf("🟥 Memory.Read(%d, %d) out of range", methodOffSet, methodByteCount)
	}
	methodStr := string(methodBuf)

	// get headers string from the wasm module function (from memory)
	headersBuf, ok := module.Memory().Read(ctx, headersOffSet, headersByteCount)
	if !ok {
		log.Panicf("🟥 Memory.Read(%d, %d) out of range", headersOffSet, headersByteCount)
	}
	headersStr := string(headersBuf)

  // headers => strings.Join(headers[:], "|")

  // get body string from the wasm module function (from memory)
	bodyBuf, ok := module.Memory().Read(ctx, bodyOffSet, bodyByteCount)
	if !ok {
		log.Panicf("🟥 Memory.Read(%d, %d) out of range", bodyOffSet, bodyByteCount)
	}
	bodyStr := string(bodyBuf)

	stringMessageFromHost := "🌍 http: " + urlStr + " method: " + methodStr + " headers: " + headersStr + " body: " + bodyStr
  // 👋 Implementation: Start


  // 👋 Implementation: End

	// write the new string to the "shared memory"
	lengthOfTheMessage := len(stringMessageFromHost)
	results, err := module.ExportedFunction("allocateBuffer").Call(ctx, uint64(lengthOfTheMessage))
	if err != nil {
		log.Panicln(err)
	}

	retOffset := uint32(results[0])
	module.Memory().WriteUint32Le(ctx, retBufPtrPos, retOffset)
	module.Memory().WriteUint32Le(ctx, retBufSize, uint32(lengthOfTheMessage))

	// add the message to the memory of the module
	module.Memory().Write(ctx, retOffset, []byte(stringMessageFromHost))
}
