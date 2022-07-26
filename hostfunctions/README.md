# Host functions

## Add an host function

In `hostfunctions` directory, create a file `function_name.go` (use this package: `package hostfunctions`)

> Example: `http.go`

```golang
package hostfunctions

import (
	"context"
	"log"

	"github.com/tetratelabs/wazero/api"
)

func Http(ctx context.Context, module api.Module, urlOffset, urlByteCount, methodOffSet, methodByteCount, retBufPtrPos, retBufSize uint32) {
  // foo...
}
```

In `services/common/common.go` add the host funtion to the runtime:

> Example: `common.go`

```golang
// 🏠 Add host functions
_, errEnv := wasmRuntime.NewModuleBuilder("env").
  ExportFunction("hostLogString", hostfunctions.LogString).
  ExportFunction("hostGetHostInformation", hostfunctions.GetHostInformation).
  ExportFunction("hostPing", hostfunctions.Ping).
  ExportFunction("hostHttp", hostfunctions.Http). 👋👋👋
  Instantiate(ctx, wasmRuntime)
```

In `helpers/functions` directory, create a file `function_name.go` (use this package: `package helpers`

> Example: `http.go`

```golang
// host functions
package hf

//export hostHttp
func hostHttp(urlOffset, urlByteCount, methodOffSet, methodByteCount uint32, retBuffPtrPos **byte, retBuffSize *int)

func Http(url, method string) string {
  // Get parameters from the wasm module
  // Prepare parameters for the host function call
  urlStrPos, urlStrSize := GetStringPtrPositionAndSize(url)
  methodStrPos, methodStrSize := GetStringPtrPositionAndSize(method)

  // This will be used to retrieve the return value (result)
  var bufPtr *byte
	var bufSize int

  hostHttp(urlStrPos, urlStrSize, methodStrPos, methodStrSize, &bufPtr, &bufSize)

  result := GetStringResult(bufPtr, bufSize)

  return result
}

```
