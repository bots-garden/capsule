package capsule

import (
	"context"
	"errors"
	"fmt"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

func GetNewWasmRuntimeForHttp(wasmFile []byte) (runtime wazero.Runtime, module api.Module, function api.Function, context context.Context) {
	runtime, module, context = CreateWasmRuntimeAndModuleInstances(wasmFile)
	function = module.ExportedFunction("callHandleHttp")
	return runtime, module, function, context
}

func GetNewWasmRuntime(wasmFile []byte) (runtime wazero.Runtime, module api.Module, function api.Function, context context.Context) {
	runtime, module, context = CreateWasmRuntimeAndModuleInstances(wasmFile)
	function = module.ExportedFunction("callHandle")
	return runtime, module, function, context
}

// ReserveMemorySpaceFor :
// Reserve a place for a string parameter in the wasm module shared memory
func ReserveMemorySpaceFor(s string, wm api.Module, ctx context.Context) (pos uint64, length uint64, free api.Function, err error) {
	length = uint64(len(s))
	malloc := wm.ExportedFunction("malloc")
	free = wm.ExportedFunction("free")

	results, err := malloc.Call(ctx, length)
	if err != nil {
		//log.Panicln("😡 out of bounds memory access", err)
		return 0, 0, free, errors.New("😡 out of bounds memory access")
	}
	stringParameterPtrPosition := results[0]
	// This pointer is managed by TinyGo, but TinyGo is unaware of external usage.
	// So, we have to free it when finished
	//defer free.Call(wr.Ctx, stringParameterPtrPosition)

	// The pointer is a linear memory offset, which is where we write the name.
	if !wm.Memory().Write(ctx, uint32(stringParameterPtrPosition), []byte(s)) {
		//log.Panicf("😡 Memory.Write(%d, %d) out of range of memory size %d", stringParameterPtrPosition, stringParameterLength, wr.Module.Memory().Size(wr.Ctx))
		return 0, 0, free, errors.New("😡 Memory.Write out of range of memory size")
	} else {
		return stringParameterPtrPosition, length, free, nil
	}
}

// ExecHandleFunction :
// params: pos1, length1, pos2, length2, ...
func ExecHandleFunction(function api.Function, module api.Module, ctx context.Context, params ...uint64) (bytes []byte, err error) {
	// This shows how to
	// read-back something allocated by TinyGo.
	handleResultArray, err := function.Call(ctx, params...)
	if err != nil {
		//log.Panicln(err)
		fmt.Println("😡[execHandleFunction]", err)
		return nil, err
	}
	// Note: This pointer is still owned by TinyGo,
	// so don't try to free it!
	handleReturnPtrPos, handleReturnSize := GetPackedPtrPositionAndSize(handleResultArray)

	// The pointer is a linear memory offset,
	// which is where we write the name.
	bytes, ok := module.Memory().Read(ctx, handleReturnPtrPos, handleReturnSize)
	if !ok {
		return nil, errors.New("😡[execHandleFunction] Memory.Read out of range of memory size")
	}
	return bytes, nil
}
