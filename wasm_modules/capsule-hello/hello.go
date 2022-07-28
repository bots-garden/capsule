package main

// TinyGo wasm module
import (
    hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

// main is required.
func main() {
    hf.SetHandleHttp(Handle)
}

func Handle(bodyReq string, headersReq map[string]string) (bodyResp string, headersResp map[string]string, errResp error) {

    hf.Log("📝 body: " + bodyReq)

    hf.Log("Content-Type: " + headersReq["Content-Type"])
    hf.Log("Content-Length: " + headersReq["Content-Length"])
    hf.Log("User-Agent: " + headersReq["User-Agent"])

    envMessage, err := hf.GetEnv("MESSAGE")
    if err != nil {
        hf.Log("😡 " + err.Error())
    } else {
        hf.Log("Environment variable: " + envMessage)
    }

    headersResp = map[string]string{
        "Content-Type": "application/json; charset=utf-8",
        "Message":      "👋 hello world 🌍",
    }

    bodyResp = `{"message": "Hello how are you?"}`

    return bodyResp, headersResp, nil
    //return bodyResp, headersResp , errors.New("😡 oups I did it again")
}

// TODO: helpers: SetHeader() ...
// TODO: be able to return a status code

/*
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json' \
  -d '{"message": "Golang 💚 wasm"}'
*/
