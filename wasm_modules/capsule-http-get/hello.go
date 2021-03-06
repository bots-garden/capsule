// wasm module
package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

// main is required.
func main() {
	hf.SetHandle(Handle)
}

func Handle(params []string) (string, error) {

	hf.Log("💊 Get sample: parameter is: " + params[0])

	headers := map[string]string{"Accept": "application/json", "Content-Type": "text/html; charset=UTF-8"}

	ret, err := hf.Http("https://httpbin.org/get", "GET", headers, "")
	if err != nil {
		hf.Log("😡 error:" + err.Error())
	} else {
		hf.Log("💊👋 Return value from the module: " + ret)
	}

	return ret, err
}
