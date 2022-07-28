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

func Handle(bodyReq string, headersReq map[string]string) (bodyResp string, headersResp map[string]string, errResp error) {

    hf.Log("param: " + bodyReq)
    hf.Log("Content-Type: " + headersReq["Content-Type"])
    hf.Log("Content-Length: " + headersReq["Content-Length"])
    hf.Log("User-Agent: " + headersReq["User-Agent"])

    headersResp = map[string]string{"Content-Type": "application/json; charset=utf-8"}

    bodyResp = `{"message": "👋 you sent me this:`+bodyReq+`"}`

	return bodyResp, headersResp , nil

    // [BODY]{"message": "👋 you sent me this:{"message":"Golang 💚 wasm"}"}[HEADERS]
}
// TODO: helpers: SetHeader() ...
