package main

// TinyGo wasm module
import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

func main() {
	hf.SetHandleHttp(Handle)
	hf.Log("🖖" + hf.GetHostInformation())

	value, err := hf.MemoryGet("message")
	if err != nil {
		hf.Log("😡 Error: " + err.Error())
	} else {
		hf.Log("📝 message: " + value)
	}
}

//export OnLoad
func OnLoad() {
	hf.Log("👋 from the OnLoad method")

	_, err := hf.MemorySet("message", "🚀 hello is started")
	if err != nil {
		hf.Log(err.Error())
	}
}

func Handle(request hf.Request) (response hf.Response, errResp error) {

	hf.Log("Body: " + request.Body)
	hf.Log("URI: " + request.Uri)
	hf.Log("Method: " + request.Method)

	params := request.ParseQueryString()

	//curl http://localhost:7070/?a=1&b=2
	for key, value := range params {
		hf.Log(key + " : " + value)
	}

	html := `
    <html>
        <head>
            <title>Wasm is fantastic 😍</title>
        </head>

        <body>
            <h1>👋 Hello World 🌍</h1>
            <h2>Served with 💚💜 with Capsule 💊</h2>
        </body>

    </html>
    `

	headersResp := map[string]string{
		"Content-Type": "text/html; charset=utf-8",
	}

	return hf.Response{Body: html, Headers: headersResp}, nil
}
