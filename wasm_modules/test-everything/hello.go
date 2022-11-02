package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
	"strings"

	/* string to json */
	"github.com/tidwall/gjson"
	/* create json string */
	"github.com/tidwall/sjson"
)

func main() {

	hf.SetHandleHttp(Handle)
}

func Handle(request hf.Request) (response hf.Response, errResp error) {
	/*
	   bodyReq = {"author":"Philippe","message":"Golang 💚 wasm"}
	*/

	// TODO: make an array of errors and print a test report at the end

	// +++ GetHostInformation +++
	hostInformation := hf.GetHostInformation()
	hf.Log(hostInformation)
	if strings.HasPrefix(hostInformation, `{"httpPort":`) && strings.Contains(hostInformation, "capsuleVersion") {
		hf.Log("-> Test: GetHostInformation: 🟢")
	} else {
		hf.Log("-> Test: GetHostInformation: 🔴")
	}
	// {"httpPort":7070,"capsuleVersion":"v0.2.9 🦜 [parrot]"}

	// +++ Http GET +++
	headers := map[string]string{"Accept": "application/json", "Content-Type": "text/html; charset=UTF-8"}

	retGet, err := hf.Http("https://httpbin.org/get", "GET", headers, "[GET]👋 hello world 🌍")
	if err != nil {
		hf.Log("😡 error:" + err.Error())
		hf.Log("-> Test: Http GET: 🔴")
	} else {
		hf.Log("💊👋 Return value from the module: " + retGet)
		hf.Log("-> Test: Http GET: 🟢")
	}

	// +++ Http POST +++
	retPost, err := hf.Http("https://httpbin.org/post", "POST", headers, "[POST]👋 hello world 🌍")
	if err != nil {
		hf.Log("😡 error:" + err.Error())
		hf.Log("-> Test: Http POST: 🔴")
	} else {
		hf.Log("💊👋 Return value from the module: " + retPost)
		hf.Log("-> Test: Http POST: 🟢")
	}

	// this is subject to change
	capsuleVersion, _ := hf.MemoryGet("capsule_version")
	hf.Log("🖐 Version: " + capsuleVersion)

	hf.Log("📝 Body: " + request.Body)
	hf.Log("📝 URI: " + request.Uri)
	hf.Log("📝 Method: " + request.Method)

	author := gjson.Get(request.Body, "author")
	message := gjson.Get(request.Body, "message")
	hf.Log("👋 " + message.String() + " by " + author.String() + " 😄")

	hf.Log("Content-Type: " + request.Headers["Content-Type"])
	hf.Log("Content-Length: " + request.Headers["Content-Length"])
	hf.Log("User-Agent: " + request.Headers["User-Agent"])

	envMessage, err := hf.GetEnv("MESSAGE")
	if err != nil {
		hf.Log("😡 " + err.Error())
	} else {
		hf.Log("Environment variable: " + envMessage)
	}

	headersResp := map[string]string{
		"Content-Type": "application/json; charset=utf-8",
		"Message":      "👋 hello world 🌍",
	}

	jsondoc := `{"message": "", "author": ""}`
	jsondoc, _ = sjson.Set(jsondoc, "message", "👋 hey! What's up?")
	jsondoc, _ = sjson.Set(jsondoc, "author", "Bob")

	return hf.Response{Body: jsondoc, Headers: headersResp}, err
}

/*
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json' \
  -d '{"message": "Golang 💚 wasm"}'
*/
