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
	/*
	   bodyReq = {"author":"Philippe","message":"Golang 💚 wasm"}
	*/
	hf.Log("📝 body: " + req.Body)

	author := gjson.Get(req.Body, "author")
	message := gjson.Get(req.Body, "message")
	hf.Log("👋 " + message.String() + " by " + author.String() + " 😄")

	hf.Log("Content-Type: " + req.Headers["Content-Type"])
	hf.Log("Content-Length: " + req.Headers["Content-Length"])
	hf.Log("User-Agent: " + req.Headers["User-Agent"])

	envMessage, err := hf.GetEnv("MESSAGE")
	if err != nil {
		hf.Log("😡 " + err.Error())
	} else {
		hf.Log("Environment variable: " + envMessage)
	}

	headers := map[string]string{
		"Content-Type": "application/json; charset=utf-8",
		"Message":      "👋 hey people 🌍",
	}

	jsondoc := `{"message": "", "author": ""}`
	jsondoc, _ = sjson.Set(jsondoc, "message", "👋 Hey! What's up?")
	jsondoc, _ = sjson.Set(jsondoc, "author", "Bob")

	return hf.Response{Body: jsondoc, Headers: headers}, nil
}

/*
curl -v -X POST \
  http://localhost:9092 \
  -H 'content-type: application/json' \
  -d '{"message": "Golang 💚 wasm"}'
*/
