package main

// TinyGo wasm module
import (
	"github.com/bots-garden/capsule/capsulemodule/flatjson"
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
	"strconv"
)

// main is required.
func main() {
	hf.SetHandleHttp(Handle)
}

func Handle(request hf.Request) (response hf.Response, errResp error) {

	jsonMap := flatjson.StrToMap(request.Body)

	hf.Log("------------------------------------------------")
	hf.Log("📝 Body: " + request.Body)
	hf.Log("📝 URI: " + request.Uri)
	hf.Log("📝 Method: " + request.Method)

	author := jsonMap["author"].(string)
	age := jsonMap["age"].(int)
	weight := jsonMap["weight"].(float64)
	isHuman := jsonMap["human"].(bool)
	message := jsonMap["message"].(string)

	hf.Log("👋 " + message + " by " + author + " 😄")
	hf.Log("👋 age: " + strconv.Itoa(age))
	hf.Log("👋 weight: " + strconv.FormatFloat(weight, 'f', 6, 64))

	if isHuman {
		hf.Log("I'm not a 🤖")
	}

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

	responseMap := map[string]interface{}{
		"message": "👋 hey! What's up?",
		"author":  "Bob",
	}

	return hf.Response{Body: flatjson.MapToStr(responseMap), Headers: headersResp}, err
}

/*
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json' \
  -d '{"message": "Golang 💚 wasm"}'
*/
