package main

// TinyGo wasm module
import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"

	/* create json string */
	"github.com/tidwall/sjson"
)

// main is required.
func main() {

	hf.SetHandleHttp(Handle)
}

func Handle(request hf.Request) (response hf.Response, errResp error) {

	headersResp := map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}

	jsondoc := `{"message": ""}`
	jsondoc, _ = sjson.Set(jsondoc, "message", "👋 hello world 🌍")

	return hf.Response{Body: jsondoc, Headers: headersResp}, nil
}

/*
curl http://localhost:7070
*/
