# 💊 Capsule
> - 🖐 I'm learning Go
> - Issues: https://github.com/bots-garden/capsule/issues
> - Last release: `v0.2.0 🪲 [beetle]`
> - Dev release (next release): `v0.2.1 TBD`

## What's new

- 0.2.0: `OnLoad` & `OnExit` functions + Memory cache host functions (`MemorySet`, `MemoryGet`, `MemoryKeys`)
- 0.1.9: Add `Request` and `Response` types (for the Handle function)
- 0.1.8: Redis host functions: add the KEYS command (`RedisKeys(pattern string))

## What is **Capsule**?

**Capsule** is a WebAssembly function launcher. It means that, with **Capsule** you can:

- From your terminal, execute a function of a wasm module (the **"CLI mode"**)
- Serving a function of a wasm module through http (the **"HTTP mode"**)

> 🖐 **The functions are developed with GoLang and compiled to wasm with TinyGo**

📦 Before executing or running a function, you need to download the last release of **Capsule**: https://github.com/bots-garden/capsule/releases/tag/0.1.9 (`v0.1.9 🐞[ladybug]`)

> - **Capsule** is developed with GoLang and thanks to the 💜 **[Wazero](https://github.com/tetratelabs/wazero)** project
> - The wasm modules are developed in GoLang and compiled with TinyGo (with the WASI specification)

👋 You will find some **running examples** with these projects:
- https://github.com/bots-garden/capsule-launcher-demo
- https://github.com/bots-garden/capsule-hello-universe

> Old samples to be updated:
> - https://github.com/bots-garden/capsule-samples
> - https://github.com/bots-garden/capsule-on-fly-dot-io

## First CLI function

Create a `go.mod` file: (`go mod init cli-say-hello`)
```
module cli-say-hello

go 1.18
```

Install the Capsule dependencies:
```bash
go get github.com/bots-garden/capsule/capsulemodule/hostfunctions
```

Create a `hello.go` file:
```go
package main

import hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"

// main is required.
func main() {
	hf.SetHandle(Handle)
}

func Handle(params []string) (string, error) {
	var err error
	for _, param := range params {
		hf.Log("- parameter is: " + param)
	}

    ret := "The first parameter is: " + params[0]

    return ret, err // err = nil
}
```
> - `hf.SetHandle(Handle)` defines the called wasm function
> - `hf.Log(string)` prints a value

Build the wasm module:
```bash
tinygo build -o hello.wasm -scheduler=none --no-debug -target wasi ./hello.go
```

Execute the `Handle` function:
```bash
./capsule \
   -wasm=./hello.wasm \
   -mode=cli \
   "👋 hello world 🌍🎃" 1234 "Bob Morane"
```
> - `-wasm` flag: the path to the wasm file
> - `-mode` execution mode


*output:*
```bash
- parameter is: 👋 hello world 🌍🎃
- parameter is: 1234
- parameter is: Bob Morane
The first parameter is: 👋 hello world 🌍🎃
```

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

## Host functions

**Capsule** offers some capabilities to the wasm modules by providing some "host functions":

### Print a message

```golang
hf.Log("👋 Hello World 🌍")
```

### Read and Write files

```golang
txt, err := hf.ReadFile("about.txt")
if err != nil {
    hf.Log(err.Error())
}
hf.Log(txt)

newFile, err := hf.WriteFile("hello.txt", "👋 HELLO WORLD 🌍")
if err != nil {
    hf.Log(err.Error())
}
hf.Log(newFile)
```

### Read value of the environment variables

```golang
message, err := hf.GetEnv("MESSAGE")
if err != nil {
    hf.Log(err.Error())
} else {
    hf.Log("MESSAGE=" + message)
}
```

### Make HTTP requests

*`GET`*
```golang
ret, err := hf.Http("https://httpbin.org/get", "GET", headers, "")
if err != nil {
    hf.Log("😡 error:" + err.Error())
} else {
    hf.Log("📝result: " + ret)
}
```

*`POST`*
```golang
headers := map[string]string{"Accept": "application/json", "Content-Type": "text/html; charset=UTF-8"}

ret, err := hf.Http("https://httpbin.org/post", "POST", headers, "👋 hello world 🌍")
if err != nil {
    hf.Log("😡 error:" + err.Error())
} else {
    hf.Log("📝result: " + ret)
}
```

### Use memory cache

*`MemorySet`*
```golang
_, err := hf.MemorySet("message", "🚀 hello is started")
```

*`MemorySet`*
```golang
value, err := hf.MemoryGet("message")
```

*`MemoryKeys`*
```golang
keys, err := hf.MemoryKeys()
// it will return an array of strings
if err != nil {
  hf.Log(err.Error())
}

for key, value := range keys {
  hf.Log(key+":"+value)
}
```

### Make Redis queries
> 🚧 this is a work in progress

You need to run **Capsule** with these two environment variables:
```bash
REDIS_ADDR="localhost:6379"
REDIS_PWD=""
```

*`SET`*
```golang
// add a key, value
res, err := hf.RedisSet("greetings", "Hello World")
if err != nil {
    hf.Log(err.Error())
} else {
    hf.Log("Value: " + res)
}
```

*`GET`*
```golang
// read the value
res, err := hf.RedisGet("greetings")
if err != nil {
    hf.Log(err.Error())
} else {
    hf.Log("Value: " + res)
}
```

*`KEYS`*
```golang
legion, err := hf.RedisKeys("bob*")
if err != nil {
    hf.Log(err.Error())
}

for _, bob := range legion {
    hf.Log(bob)
}
```

### Make CouchBase N1QL Query

You need to run **Capsule** with these four environment variables:
```bash
COUCHBASE_CLUSTER="couchbase://localhost"
COUCHBASE_USER="admin"
COUCHBASE_PWD="ilovepandas"
COUCHBASE_BUCKET="wasm-data"
```

```golang
bucketName, _ := hf.GetEnv("COUCHBASE_BUCKET")
query := "SELECT * FROM `" + bucketName + "`.data.docs"

jsonStringArray, err := hf.CouchBaseQuery(query)
```

## Capsule FaaS (experimental)

There are four additional components to use **capsule** (the wasm module launcher/executor) in **FaaS** mode:

- `capsule-reverse-proxy`: a reverse-proxy to simplify the functions (wasm modules) access
- `capsule-registry`: a wasm module registry (🚧 support of https://wapm.io/ in progress)
- `capsule-worker`: a server to start the functions (wasm modules) remotely
- `capsule-ctl` (short name: `caps`): a CLI to facilitate the interaction with the worker

See documents files in `./docs` (🚧 this is a work in progress)

👋 You will find some **running examples** with this project:
- https://github.com/bots-garden/capsule-faas-demo

> - You can use the capsule registry independently of FaaS mode, only to provide wasm modules to the capsule launcher
> - You can use the capsule reverse-proxy independently of FaaS mode, only to get only one access URL
