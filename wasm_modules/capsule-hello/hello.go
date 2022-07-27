package main

import (
	//"errors"
	"github.com/bots-garden/capsule/hostfunctions/wasmmodule"
)

// main is required.
func main() {
	hf.SetHandleHttp(Handle)
}

/*
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json' \
  -d '{"message": "Golang 💚 wasm"}'
*/

func Handle(param string, headers map[string]string) (string, error) {

    hf.Log("param: " + param)
    hf.Log("Content-Type: " + headers["Content-Type"])
    hf.Log("Content-Length: " + headers["Content-Length"])
    hf.Log("User-Agent: " + headers["User-Agent"])

	return "Hello, you send me this: " + param, nil
    //TODO: return a contentType
}

// ? HandleJson, Handle<>, ...
