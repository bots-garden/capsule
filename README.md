# 💊 Capsule

What is **Capsule**?

**Capsule** is a WebAssembly function launcher. It means that, with **Capsule** you can:

- From your terminal, execute a function of a wasm module
- Serving a function of a wasm module through http

🖐 The functions are developed with GoLang and compiled to wasm with TinyGo

📦 Before executing or running a function, you need to download the last release of **Capsule**: https://github.com/bots-garden/capsule/releases/tag/0.1.5 (`v0.1.5 🦐`)

> - **Capsule** is developed with GoLang and thanks to the 💜 **[Wazero](https://github.com/tetratelabs/wazero)** project
> - The wasm modules are developed in GoLang and compiled with TinyGo (with the WASI specification)

👋 You will find some **running examples** with this project: https://github.com/bots-garden/capsule-samples


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

func Handle(bodyReq string, headersReq map[string]string) (bodyResp string, headersResp map[string]string, errResp error) {
	hf.Log("📝 body: " + bodyReq)

	// Read the request headers
	hf.Log("Content-Type: " + headersReq["Content-Type"])
	hf.Log("Content-Length: " + headersReq["Content-Length"])
	hf.Log("User-Agent: " + headersReq["User-Agent"])

	// Read the MESSAGE environment variable
	envMessage, err := hf.GetEnv("MESSAGE")
	if err != nil {
		hf.Log("😡 " + err.Error())
	} else {
		hf.Log("Environment variable: " + envMessage)
	}

	// Set the response content type and add a message header
	headersResp = map[string]string{
		"Content-Type": "application/json; charset=utf-8",
		"Message":      "👋 hello world 🌍",
	}

	jsonResponse := `{"message": "hey people!"}`

	return jsonResponse, headersResp, err
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

func Handle(bodyReq string, headersReq map[string]string) (bodyResp string, headersResp map[string]string, errResp error) {
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

	headersResp = map[string]string{
		"Content-Type": "text/html; charset=utf-8",
	}

	return html, headersResp, nil
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

## Use Capsule as a reverse proxy

You can use **Capsule** as a reverse proxy. Then, you can call a function by its name:
```bash
http://localhost:8888/functions/hola
```
> *The reverse proxy will serve the **default** version of the `hola` function*

Or, you can use a revision of the function (for example, if you use several version of the function):
```bash
http://localhost:8888/functions/hola/orange
```
> - *The reverse proxy will serve the `orange` revision of the `hola` function*
> - *The `default` revision is the `default` version of the function (http://localhost:8888/functions/hola)*

To run **Capsule** as a reverse proxy, run it like this:
```bash
./capsule \
   -mode=reverse-proxy \
   -config=./config.yaml \
   -httpPort=8888
```
> *You have to define a configuration yaml file.*

### Define the routes in a yaml file (static mode)

*config.yaml*
```yaml
hello:
    default:
        - http://localhost:9091
        - http://localhost:7071

hey:
    default:
        - http://localhost:9092

hola:
    default:
        - http://localhost:9093
    orange:
        - http://localhost:6061
    yellow:
        - http://localhost:6062
```
> *A revision can be a set of URLs. In this case, the Capsule reverse-proxy will use randomly one of the URLs.*

### Use the "in memory" dynamic mode of the reverse-proxy

With the reverse proxy mode of Capsule, you gain an **API** that allows to define routes dynamically (in memory). You can keep the yaml config file (it is loaded in memory at startup).

To run **Capsule** as a reverse proxy, with the "in memory" dynamic mode, add this flag: `-backend="memory"`:
```bash
./capsule \
   -mode=reverse-proxy \
   -config=./config.yaml \
   -backend="memory" \
   -httpPort=8888
```

#### Registration API

##### Register a function (add a new route to a function)

```bash
curl -v -X POST \
  http://localhost:8888/memory/functions/registration \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"function": "morgen", "revision": "default", "url": "http://localhost:5050"}'
```
> - This will add a new entry to the routes list with a `default` revision, with one url `http://localhost:5050`
> - You can call the function with this url: http://localhost:8888/function/morgen

The routes list (it's a map) will look like that:

```json
{
    "morgen": {
        "default": [
            "http://localhost:5050"
        ]
    }
}
```

You can create a new function registration with a named revision:

```bash
curl -v -X POST \
  http://localhost:8888/memory/functions/registration \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"function": "morgen", "revision": "magenta", "url": "http://localhost:5051"}'
```
> - This will add a new entry to the routes list with a `magenta` revision, with one url `http://localhost:5051`
> - You can call the function with this url: http://localhost:8888/function/morgen/magenta


##### Remove the registration

```bash
curl -v -X DELETE \
  http://localhost:8888/memory/functions/registration \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"function": "morgen"}'
```


#### Revision API

##### Add a revision to the function registration

```bash
curl -v -X POST \
  http://localhost:8888/memory/functions/morgen/revision \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"revision": "blue", "url": "http://localhost:5051"}'
```
> - The function already exists
> - The name of the function is set in the url `http://localhost:8888/memory/functions/:function_name/revision`

The routes list will look like that:

```json
{
    "morgen": {
        "blue": [
            "http://localhost:5051"
        ],
        "default": [
            "http://localhost:5050"
        ]
    }

}
```

##### Remove a revision from the function registration

```bash
curl -v -X DELETE \
  http://localhost:8888/memory/functions/morgen/revision \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"revision": "blue"}'
```

#### URL API

##### Add a URL to the revision of a function

```bash
curl -v -X POST \
  http://localhost:8888/memory/functions/morgen/blue/url \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"url": "http://localhost:5053"}'
```
> - The revision already exists
> - The name of the function and of the revision are set in the url `http://localhost:8888/memory/functions/:function_name/:function_revision/url`

The routes list will look like that:

```json
{
    "morgen": {
        "blue": [
            "http://localhost:5051",
            "http://localhost:5053"
        ],
        "default": [
            "http://localhost:5050"
        ]
    }

}
```

##### Remove a URL from the function revision

```bash
curl -v -X DELETE \
  http://localhost:8888/memory/functions/morgen/blue/url \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"url": "http://localhost:5053"}'
```

## Use Capsule as a Wasm modules registry
> 🚧 this is a work in progress

It's possible to download the wasm module from a remote location before serving it:

```bash
./capsule \
   -url=http://localhost:9090/hello.wasm \
   -wasm=./tmp/hello.wasm \
   -mode=http \
   -httpPort=8080
```

**Capsule** provides a `registry` mode allowing to upload and serve the wasm modules

### Start the wasm registry

```bash
./capsule \
   -mode=registry \
   -files="./wasm_modules" \
   -httpPort=4999
```
> The `-files` tag defines where the modules are uploaded

### Upload a wasm module

```bash
curl -X POST http://localhost:4999/upload/k33g/hola/0.0.0 \
  -F "file=@../with-proxy/capsule-hola/hola.wasm" \
  -F "info=hola function from @k33g" \
  -H "Content-Type: multipart/form-data"
```
> - The upload url is defined like this: `/upload/user_name_or_organization/module_name/tag`
> - The wasm module will be saved to `./wasm_modules/user_name_or_organization/module_name/tag/module_name.wasm`
> - This data `"info=hola function from @k33g"` will create a file `./wasm_modules/user_name_or_organization/module_name/tag/module_name.info`

Then you can download the module at http://localhost:4999/upload/k33g/hola/0.0.0/hello.wasm

### Download and start a wasm module

```bash
./capsule \
   -wasm=./tmp/hola.wasm \
   -url="http://localhost:4999/k33g/hola/0.0.0/hola.wasm" \
   -mode=http \
   -httpPort=7072
```

### Get information about a wasm module

```bash
curl http://localhost:4999/info/k33g/hola/0.0.0
```

### Get the list of all the wasm modules

```bash
curl http://localhost:4999/modules
```

## Use Capsule as Worker(s)

Then you will get an api to start functions remotely. For more details, see `/capsulelauncher/services/worker/README.md`
