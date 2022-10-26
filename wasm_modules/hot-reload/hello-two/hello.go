package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

//export OnLoad
func OnLoad() {
	hf.Log("🤖 I'm Hello Two")
}

func main() {
	hf.SetHandleHttp(Handle)
}

func Handle(request hf.Request) (response hf.Response, errResp error) {

	headersResp := map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}

	jsondoc := `{"message": "👋 hello world 🌍"}`

	return hf.Response{Body: jsondoc, Headers: headersResp}, nil
}

/*
curl http://localhost:7070
*/
