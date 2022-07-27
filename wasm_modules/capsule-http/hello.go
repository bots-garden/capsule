// wasm module
package main

import (
	hf "github.com/bots-garden/capsule/helpers/functions"
)

// main is required.
func main() {
	hf.SetHandle(Handle)
}

func Handle(param string) (string, error) {

	hf.Log("1️⃣ parameter is: " + param)

	headers := map[string]string{"Accept": "application/json", "Content-Type": "text/html; charset=UTF-8"}

	ret, err := hf.Http("https://httpbin.org/get", "GET", headers, "")
	if err != nil {
		hf.Log("😡 This is an error:" + err.Error())
	} else {
        hf.Log("1️⃣ From module: " + ret)
    }

	ret2, err2 := hf.Http("https://httpbin.nowhere/get", "GET", headers, "🎉")
	if err2 != nil {
		hf.Log("😡 This is an error:" + err2.Error())
	} else {
        hf.Log("2️⃣ From module: " + ret2)
    }

	return ret, nil
}

// ? HandleJson, Handle<>, ...
/*
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json' \
  -d '{"message": "Golang 💚 wasm"}'
*/
