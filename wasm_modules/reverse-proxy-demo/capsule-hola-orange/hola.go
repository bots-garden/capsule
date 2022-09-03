package main

// TinyGo wasm module
import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

func main() {
	hf.SetHandleHttp(Handle)
}

func Handle(bodyReq string, headersReq map[string]string) (resp hf.Response, errResp error) {
	html := `
    <html>
        <head>
            <title>Wasm is fantastic 😍</title>
        </head>

        <body>
            <h1>👋 Hola Mundo 🌍</h1>
            <h2>Served with 💚💜 with Capsule 💊</h2>
            <h3>🟠🟠🟠</h3>
        </body>

    </html>
    `

	headers := map[string]string{
		"Content-Type": "text/html; charset=utf-8",
	}

	return hf.Response{Body: html, Headers: headers}, nil
}
