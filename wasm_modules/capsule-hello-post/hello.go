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

func Handle(bodyReq string, headersReq map[string]string) (bodyResp string, headersResp map[string]string, errResp error) {
    /*
       bodyReq = {"author":"Philippe","message":"Golang 💚 wasm"}
    */
    hf.Log("📝 body: " + bodyReq)

    author := gjson.Get(bodyReq, "author")
    message := gjson.Get(bodyReq, "message")
    hf.Log("👋 " + message.String() + " by " + author.String() + " 😄")

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

    jsondoc := `{"message": "", "author": ""}`
    jsondoc, _ = sjson.Set(jsondoc, "message", "👋 hey! What's up?")
    jsondoc, _ = sjson.Set(jsondoc, "author", "Bob")

    return jsondoc, headersResp, nil
    //return jsondoc, headersResp , errors.New("😡 oups I did it again")
}

// TODO: helpers: SetHeader() ...
// TODO: be able to return a status code

/*
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json' \
  -d '{"message": "Golang 💚 wasm"}'
*/
