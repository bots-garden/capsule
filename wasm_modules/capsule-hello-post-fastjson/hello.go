package main

// TinyGo wasm module
import (
    hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"

    "github.com/valyala/fastjson"
)

// main is required.
func main() {

    hf.SetHandleHttp(Handle)
}

func Handle(request hf.Request) (response hf.Response, errResp error) {
    /*
       bodyReq = {"author":"Philippe","message":"Golang 💚 wasm"}
    */

    parser := fastjson.Parser{}
    jsonValue, err := parser.Parse(request.Body)

    capsuleVersion, _ := hf.MemoryGet("capsule_version")
    hf.Log("🖐 Version: " + capsuleVersion)

    hf.Log("📝 Body: " + request.Body)
    hf.Log("📝 URI: " + request.Uri)
    hf.Log("📝 Method: " + request.Method)

    author := string(jsonValue.GetStringBytes("author"))
    message := string(jsonValue.GetStringBytes("message"))

    //author := gjson.Get(request.Body, "author")
    //message := gjson.Get(request.Body, "message")
    hf.Log("👋 " + message + " by " + author + " 😄")

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

    jsondoc := `{"message": "👋 hey! What's up?", "author": "Bob"}`

    return hf.Response{Body: jsondoc, Headers: headersResp}, err
}

/*
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json' \
  -d '{"message": "Golang 💚 wasm"}'
*/
