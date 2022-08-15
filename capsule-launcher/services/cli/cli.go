package capsulecli

import (
	"fmt"
	capsule "github.com/bots-garden/capsule/capsule-launcher/services/wasmrt"
	"github.com/bots-garden/capsule/commons"
	"log"
	"strconv"
)

// Execute :
// Pass a string param and get a string result
func Execute(args []string, wasmFile []byte) {

	// 👋 get Wasm Module Instance (and Wasm runtime)
	wasmRuntime, wasmModule, wasmFunction, ctx := capsule.GetNewWasmRuntime(wasmFile)
	// defer must always be in the main code (to avoid go routine panic)
	defer wasmRuntime.Close(ctx)

	// TODO change the separator (same for headers etc...)
	// the string is built with args of main
	params := commons.CreateStringFromSlice(args, "°")

	paramsPos, paramsLen, free, err := capsule.ReserveMemorySpaceFor(params, wasmModule, ctx)
	defer free.Call(ctx, paramsPos)

	// get the callHandle function
	bytes, err := capsule.ExecHandleFunction(wasmFunction, wasmModule, ctx, paramsPos, paramsLen)

	if err != nil {
		log.Panicf("out of range of memory size")
	}

	returnValue := string(bytes)
	// check the return value
	if commons.IsErrorString(returnValue) {
		errorMessage, errorCode := commons.GetErrorStringInfo(returnValue)
		if errorCode == 0 {
			returnValue = errorMessage
		} else {
			returnValue = errorMessage + " (" + strconv.Itoa(errorCode) + ")"
		}
	}

	fmt.Println(returnValue) // the result
}
