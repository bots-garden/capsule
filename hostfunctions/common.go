package hostfunctions

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/tetratelabs/wazero/api"
)

// Write string to the memory of the module
// The Host writes to memory
func WriteStringToMemory(text string, ctx context.Context, module api.Module,
	retBuffPtrPos, retBuffSize uint32){

    stringMessageFromHost := text
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

// Get string from the module's memory (written by the module)
// (argument of a function)
func ReadStringFromMemory(ctx context.Context, module api.Module, contentOffset, contentByteCount uint32) string {
  contentBuff, ok := module.Memory().Read(ctx, contentOffset, contentByteCount)
	if !ok {
		log.Panicf("🟥 Memory.Read(%d, %d) out of range", contentOffset, contentByteCount)
	}
	contentStr := string(contentBuff)
  return contentStr
}

// to not display an error code at the end of the error message, code == 0
func CreateStringError(message string, code int) string {
    return "[ERR]["+strconv.Itoa(code)+"]:"+message
    // "[ERR][200]:hello world"
    // message: e.split("]:")[1]
    // code: e.split("]")[1].split("[")[1]
}

func CreateSliceFromString(str string, separator string) []string {
    return strings.Split(str, separator)
}

func CreateMapFromSlice(strSlice []string, separator string) map[string]string {
    strMap := make(map[string]string)
    for _, item := range strSlice {
        res := strings.Split(item, separator)
        strMap[res[0]] = res[1]
    }
    return strMap
}
