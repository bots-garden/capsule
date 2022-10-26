# Developer Guide

## First HTTP function

Create a `go.mod` file: (`go mod init http-say-hello`)
```
module http-say-hello

go 1.18
```

To serve the function through http, you need to change the signature of the `Handle` function:

```golang
package main

import hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"

// main is required.
func main() {
	hf.SetHandleHttp(Handle)
}

func Handle(request hf.Request) (response hf.Response, errResp error) {
    hf.Log("📝 Body: " + request.Body)

	// Read the request headers
    hf.Log("Content-Type: " + request.Headers["Content-Type"])
    hf.Log("Content-Length: " + request.Headers["Content-Length"])
    hf.Log("User-Agent: " + request.Headers["User-Agent"])

	// Read the MESSAGE environment variable
	envMessage, err := hf.GetEnv("MESSAGE")
	if err != nil {
		hf.Log("😡 " + err.Error())
	} else {
		hf.Log("Environment variable: " + envMessage)
	}

	// Set the response content type and add a message header
	headersResp := map[string]string{
		"Content-Type": "application/json; charset=utf-8",
		"Message":      "👋 hello world 🌍",
	}

	jsonResponse := `{"message": "hey people!"}`

	return hf.Response{Body: jsonResponse, Headers: headersResp}, err
}
```
> - `hf.SetHandleHttp(Handle)` defines the called wasm function
> - `hf.Log(string)` prints a value
> - `hf.GetEnv("MESSAGE")` get the value of the `MESSAGE` environment variable

Build the wasm module:
```bash
tinygo build -o hello.wasm -scheduler=none --no-debug -target wasi ./hello.go
```

Serve the `Handle` function:
```bash
export MESSAGE="🖐 good morning 😄"
./capsule \
   -wasm=./hello.wasm \
   -mode=http \
   -httpPort=8080
```


Call the `Handle` function:
```bash
curl -v -X POST \
  http://localhost:8080 \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"message": "TinyGo 💚 wasm"}'
```

*request output:*
```bash
> POST / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.79.1
> Accept: */*
> content-type: application/json; charset=utf-8
> Content-Length: 31
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Message: 👋 hello world 🌍
< Date: Sat, 30 Jul 2022 19:17:28 GMT
< Content-Length: 26
<
{"message":"hey people!"}
```

*log server output:*
```bash
📝 body: {"message":"TinyGo 💚 wasm"}
Content-Type: application/json; charset=utf-8
Content-Length: 31
User-Agent: curl/7.79.1
Environment variable: 🖐 good morning 😄
```

### OnLoad function

If you add an `OnLoad` exported function to the module, it will be executed at the start of the HTTP launcher (capsule).
>  *the `main` function will be executed too*

```golang
//export OnLoad
func OnLoad() {
	hf.Log("👋 from the OnLoad function")
}
```
> It can be useful to register your wasm service to a backend (Redis, CouchBase, ...)

### OnExit function

If you add an `OnExit` exported function to the module, it will be executed when you stop the HTTP launcher (capsule).
>  *the `main` function will be executed too*

```golang
//export OnExit
func OnExit() {
	hf.Log("👋 from the OnExit function")
}
```
> It can be useful to unregister your wasm service from a backend (Redis, CouchBase, ...)


### GetExitError and GetExitCode function
> 🖐🚧 it's a work in progress (it's not implemented entirely)
```golang
//export OnExit
func OnExit() {
	hf.Log("👋🤗 have a nice day")
	hf.Log("Exit Error: " + hf.GetExitError())
	hf.Log("Exit Code: " + hf.GetExitCode())
}
```

## Remote loading of the wasm module

You can download the wasm module from a remote location before executing it:

For example, provide the wasm file with an HTTP server, run this command at the root of your project:
```bash
python3 -m http.server 9090
```
> Now you can download the wasm file with this url: http://localhost:9090/hello.wasm


Serve the `Handle` function:
```bash
export MESSAGE="🖐 good morning 😄"
./capsule \
   -url=http://localhost:9090/hello.wasm \
   -wasm=./tmp/hello.wasm \
   -mode=http \
   -httpPort=8080
```
> - `-url` flag: the download url
> - `-wasm` flag: the path where to save the wasm file

## GET Request

**Capsule** accept the `GET` requests, so you can serve, for example, HTML:

```golang
package main

import hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"

// main is required.
func main() {
	hf.SetHandleHttp(Handle)
}

func Handle(request hf.Request) (response hf.Response, errResp error) {
	html := `
    <html>
        <head>
            <title>Wasm is fantastic 😍</title>
        </head>

        <body>
            <h1>👋 Hello World 🌍</h1>
            <h2>Served with 💜 with Capsule 💊</h2>
        </body>

    </html>
    `

	headersResp := map[string]string{
		"Content-Type": "text/html; charset=utf-8",
	}

	return hf.Response{Body: html, Headers: headersResp}, nil
}
```

Build the wasm module:
```bash
tinygo build -o hello.wasm -scheduler=none --no-debug -target wasi ./hello.go
```

Serve the `Handle` "function page":
```bash
./capsule \
   -wasm=./hello.wasm \
   -mode=http \
   -httpPort=8080
```

Now, you can open http://localhost:8080 with your browser or run a curl request:
```bash
curl http://localhost:8080
```
