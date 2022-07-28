package hostfunctions

import (
	"context"
	"github.com/tetratelabs/wazero/api"
	"os"
)

// GetEnv :
// The wasm module will call this function like this:
// `func GetEnv(param string) (string, error)`
func GetEnv(ctx context.Context, module api.Module, varNameOffset, varNameByteCount, retBuffPtrPos, retBuffSize uint32) {

	//=========================================================
	// Read arguments values of the function call
	//=========================================================
	// get string from the wasm module function (from memory)

	varNameStr := ReadStringFromMemory(ctx, module, varNameOffset, varNameByteCount)

	//==[👋 Implementation: Start]=============================
	var stringResultFromHost = ""
	// do something that returns a value(`stringResultFromHost`) and an error(`err`)
	// and that uses the parameter(`varNameStr`)
	variableValue := os.Getenv(varNameStr)
	if variableValue == "" {
		stringResultFromHost = CreateStringError(varNameStr+" is empty", 0)
	} else {
		stringResultFromHost = variableValue
	}
	//fmt.Println("✅", varNameStr, "==>", variableValue)

	//==[👋 Implementation: End]===============================

	// Write the new string stringResultFromHost to the "shared memory"
	WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)
}
