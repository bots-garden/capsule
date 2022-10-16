package capsule

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
)

/*
## Remote loading of the wasm module

```bash
capsule \
   -url=http://localhost:9090/hello.wasm \
   -wasm=./tmp/hello.wasm \
   -mode=http \
   -httpPort=8080
```
- url flag: the download url
- wasm flag: the path where to save the wasm file

### Serve with python:

```bash
python3 -m http.server 9090
```

Now you can download the wasm file with this url:
http://localhost:9090/hello.wasm

*/

func loadWasmFile(path string) ([]byte, error) {
	wasmFileToLoad, errLoadWasmFile := os.ReadFile(path)

	if errLoadWasmFile != nil {
		fmt.Println("🔴 Error while loading the wasm file:", errLoadWasmFile)
		//os.Exit(1)
	}
	return wasmFileToLoad, errLoadWasmFile
}

func GetWasmFileFromUrl(wasmFileUrl, wasmFilePath string) ([]byte, error) {

	client := resty.New()
	resp, errLoadWasmFileFromUrl := client.R().
		SetOutput(wasmFilePath).
		Get(wasmFileUrl)

	if resp.IsError() {
		fmt.Println("🔴 Error while downloading the wasm file:", "empty response")
	}

	if errLoadWasmFileFromUrl != nil {
		fmt.Println("🔴 Error while downloading the wasm file:", errLoadWasmFileFromUrl)
		return nil, errLoadWasmFileFromUrl
	} else {
		fmt.Println("🙂", "file downloaded", wasmFilePath)
		return loadWasmFile(wasmFilePath)
	}

}
