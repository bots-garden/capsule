package main

// TinyGo wasm module
import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

func main() {
	hf.SetHandleHttp(Handle)
}

func Handle(bodyReq string, headersReq map[string]string) (bodyResp string, headersResp map[string]string, errResp error) {
	message, _ := hf.GetEnv("MESSAGE")
	token, _ := hf.GetEnv("TOKEN")
	html := `
    <html>
        <head>
            <meta charset="utf-8">
            <title>Wasm is fantastic 😍</title>

			<meta name="viewport" content="width=device-width, initial-scale=1">
			
			<style>
				.container { min-height: 100vh; display: flex; justify-content: center; align-items: center; text-align: center; }
				.title { font-family: "Source Sans Pro", "Helvetica Neue", Arial, sans-serif; display: block; font-weight: 300; font-size: 100px; color: #35495e; letter-spacing: 1px; }
				.subtitle { font-family: "Source Sans Pro", "Helvetica Neue", Arial, sans-serif; font-weight: 300; font-size: 42px; color: #526488; word-spacing: 5px; padding-bottom: 15px; }
				.links { padding-top: 15px; }
			</style>

        </head>

        <body>
			<section class="container">
                <div>
                    <h1 class="title">👋 Hello World 🌍</h1>
                    <h2 class="subtitle">Served with 💜 by Capsule 💊</h2>
                    <h2 class="subtitle">` + message + `</h2>
                    <h2 class="subtitle">` + token + `</h2>
                </div>
            </section>
        </body>

    </html>
    `

	headersResp = map[string]string{
		"Content-Type": "text/html; charset=utf-8",
	}

	return html, headersResp, nil
}
