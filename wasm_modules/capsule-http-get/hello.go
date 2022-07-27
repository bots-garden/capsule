// wasm module
package main

import (
	hf "github.com/bots-garden/capsule/wasmhostfunctions"
)

// main is required.
func main() {
	hf.SetHandle(Handle)
}

func Handle(param string) (string, error) {

	hf.Log("💊 Get sample: parameter is: " + param)

	headers := map[string]string{"Accept": "application/json", "Content-Type": "text/html; charset=UTF-8"}

	ret, err := hf.Http("https://httpbin.org/get", "GET", headers, "")
	if err != nil {
		hf.Log("😡 error:" + err.Error())
	} else {
        hf.Log("💊👋 Return value from the module: " + ret)
    }

	return ret, err
}
