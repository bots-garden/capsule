package hostfunctions

import (
	"context"
	"errors"
	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions/memory"
	"github.com/bots-garden/capsule/commons"
	"github.com/tetratelabs/wazero/api"
)

/*
    0- copy this template(`template.go`) to a new file functionname.go
    1- change the name of FunctionName
    2- add a reference to this function to `services/wasmrt/wasmrt.go` inside the `CreateWasmRuntime` function
    ```golang
	// 🏠 Add host functions
	_, errEnv := wasmRuntime.NewHostModuleBuilder("env").
		NewFunctionBuilder().WithFunc(hostfunctions.LogString).Export("hostLogString").
		NewFunctionBuilder().WithFunc(hostfunctions.GetHostInformation).Export("hostGetHostInformation").
		NewFunctionBuilder().WithFunc(hostfunctions.Ping).Export("hostPing").
		NewFunctionBuilder().WithFunc(hostfunctions.Http).Export("hostHttp").
		NewFunctionBuilder().WithFunc(hostfunctions.ReadFile).Export("hostReadFile").
		NewFunctionBuilder().WithFunc(hostfunctions.WriteFile).Export("hostWriteFile").
		NewFunctionBuilder().WithFunc(hostfunctions.FunctionName).Export("hostFunctionName").
		Instantiate(ctx, wasmRuntime)
    ```
    3- Implement the feature
    4- look at `./capsulemodule/hostfunctions/template.go`
*/

// FunctionName :
// The wasm module will call this function like this:
// `func FunctionName(param string) (string, error)`
func FunctionName(ctx context.Context, module api.Module, paramOffset, paramByteCount, retBuffPtrPos, retBuffSize uint32) {

	//=========================================================
	// Read arguments values of the function call
	//=========================================================
	// get string from the wasm module function (from memory)

	paramStr := memory.ReadStringFromMemory(ctx, module, paramOffset, paramByteCount)

	//==[👋 Implementation: Start]=============================
	var stringResultFromHost = ""

	// do something that returns a value(`stringResultFromHost`) and an error(`err`)
	// and that uses the parameter(`paramStr`)
	err := errors.New("errorMessage")
	something := "something:" + paramStr

	if err != nil {
		stringResultFromHost = commons.CreateStringError(err.Error(), 0)
		// if code 0 don't display code in the error message
	} else {
		stringResultFromHost = something
	}
	//==[👋 Implementation: End]===============================

	// Write the new string stringResultFromHost to the "shared memory"
	memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)
}
