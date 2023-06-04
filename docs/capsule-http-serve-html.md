# Serve HTML with Capsule HTTP, step by step

## Create an HTML file: `index.html`

```html
<html>
  <head>
    <meta charset="utf-8">
    <title>Capsule 💜 Wasm & Wazero</title>
      <meta name="viewport" content="width=device-width, initial-scale=1">
      <style>
        .container { min-height: 100vh; display: flex; justify-content: center; align-items: center; text-align: center; }
        .title { font-family: "Source Sans Pro", "Helvetica Neue", Arial, sans-serif; display: block; font-weight: 300; font-size: 80px; color: #35495e; letter-spacing: 1px; }
        .subtitle { font-family: "Source Sans Pro", "Helvetica Neue", Arial, sans-serif; font-weight: 300; font-size: 32px; color: #526488; word-spacing: 5px; padding-bottom: 15px; }
        .links { padding-top: 15px; }
      </style>
  </head>

  <body>
    <section class="container">
      <div>
        <h1 class="title">👋 Hello World 🌍</h1>
        <h2 class="subtitle">Served with 💜 by Capsule [HTTP] v0.3.9 🥒 [cucumber] 💊</h2>
        <h2 class="subtitle">🎉 Happily built thanks to Wazero</h2>
      </div>
    </section>
  </body>

</html>
```

## Create a new WASM module

```bash
go mod init index
touch main.go
```

> `main.go`
```golang
package main

import (
	_ "embed"
	capsule "github.com/bots-garden/capsule-module-sdk"
)

var (
	//go:embed index.html
	html []byte
)

func main() {
	capsule.SetHandleHTTP(func (param capsule.HTTPRequest) (capsule.HTTPResponse, error) {
		return capsule.HTTPResponse{
			TextBody: string(html),
			Headers: `{
				"Content-Type": "text/html; charset=utf-8",
				"Cache-Control": "no-cache",
				"X-Powered-By": "capsule-module-sdk"
			}`,
			StatusCode: 200,
		}, nil
	})
}
```

## Build the module

```bash
tinygo build -o index.wasm \
    -scheduler=none \
    --no-debug \
    -target wasi ./main.go 
```

## Serve the module

```bash
capsule-http --wasm=./index.wasm --httpPort=7070
```

Go to [http://localhost:7070](http://localhost:7070) with your favorite browser.
