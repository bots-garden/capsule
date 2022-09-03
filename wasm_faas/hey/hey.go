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

func Handle(req hf.Request) (resp hf.Response, errResp error) {

	hf.Log("📝 body: " + req.Body)

	author := gjson.Get(req.Body, "author")
	message := gjson.Get(req.Body, "message")
	hf.Log("👋 " + message.String() + " by " + author.String() + " 😄")

	// 👀 https://github.com/bots-garden/capsule/issues/91
	hf.Log("🟢 Content-Type: " + req.Headers["Content-Type"])
	hf.Log("🔵 Content-Length: " + req.Headers["Content-Length"])
	hf.Log("🟠 User-Agent: " + req.Headers["User-Agent"])
	hf.Log("🔴 My-Token: " + req.Headers["My-Token"])

	headers := map[string]string{
		"Content-Type": "application/json; charset=utf-8",
		"Message":      "👋 hello world 🌍",
		"MyToken":      req.Headers["My-Token"],
	}

	jsondoc := `{"message": "", "author": ""}`
	jsondoc, _ = sjson.Set(jsondoc, "message", "👋 hey! What's up?")
	jsondoc, _ = sjson.Set(jsondoc, "author", "Bob")

	return hf.Response{Body: jsondoc, Headers: headers}, nil
}
