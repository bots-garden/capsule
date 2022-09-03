package main

// TinyGo wasm module
import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
	/* string to json */
	"github.com/tidwall/gjson"
	/* create json string */
	"github.com/tidwall/sjson"
)

// main is required.
func main() {
	hf.SetHandleHttp(Handle)
}

func Handle(bodyReq string, headersReq map[string]string) (rest hf.Response, errResp error) {

	hf.Log("📝 body: " + bodyReq)

	author := gjson.Get(bodyReq, "author")
	message := gjson.Get(bodyReq, "message")
	hf.Log("👋 " + message.String() + " by " + author.String() + " 😄")

	// 👀 https://github.com/bots-garden/capsule/issues/91
	hf.Log("🟢 Content-Type: " + headersReq["Content-Type"])
	hf.Log("🔵 Content-Length: " + headersReq["Content-Length"])
	hf.Log("🟠 User-Agent: " + headersReq["User-Agent"])
	hf.Log("🔴 My-Token: " + headersReq["My-Token"])

	headers := map[string]string{
		"Content-Type": "application/json; charset=utf-8",
		"Message":      "👋 hello world 🌍",
		"MyToken":      headersReq["My-Token"],
	}

	jsondoc := `{"message": "", "author": ""}`
	jsondoc, _ = sjson.Set(jsondoc, "message", "👋 hey! What's up?")
	jsondoc, _ = sjson.Set(jsondoc, "author", "Bob")

	return hf.Response{Body: jsondoc, Headers: headers}, nil
}
