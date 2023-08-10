# 🚀 Getting Started

## Use the Capsule HTTP server

First, download the last version of the Capsule HTTP server for the appropriate OS & ARCH (and release version):

```bash
VERSION="v0.4.2" OS="linux" ARCH="arm64"
wget -O capsule-http https://github.com/bots-garden/capsule/releases/download/${VERSION}/capsule-http-${VERSION}-${OS}-${ARCH}
chmod +x capsule-http
```

## Write a WASM Capsule module

```golang
package main

import (
	"github.com/bots-garden/capsule-module-sdk"
)

func main() {
	
	capsule.SetHandleHTTP(func(param capsule.HTTPRequest) (capsule.HTTPResponse, error) {
		
		response := capsule.HTTPResponse{
			JSONBody:   `{"message": "Hello World"}`,
			Headers:    `{"Content-Type": "application/json; charset=utf-8"}`,
			StatusCode: 200,
		}
		return response, nil
	})
}
```

**Build the WASM module**:

```bash
tinygo build -o hello-world.wasm \
    -scheduler=none \
    --no-debug \
    -target wasi ./main.go 
```

## Serve a WASM Capsule module

To run a WASM Capsule module you need to set 2 flags:

- `--wasm`: the path to the WASM module
- `--params`: the parameter to pass to the WASM module

```bash
./capsule-http \
--wasm=./functions/hello-world/hello-world.wasm\
--httpPort=8080
```

You can query the service like this:
```bash
curl -X POST http://localhost:8080 \
    -H 'Content-Type: application/json; charset=utf-8' \
    -d '{"name":"Bob Morane","age":42}'
```

You can remotely download  the WASM module with the `--url` flag:
```bash
./capsule-http \
--url=http://localhost:5000/hello-world.wasm \
--wasm=./tmp/hello-world.wasm 
--httpPort=8080
```

> 👋 if you need to set an authentication header you can use these flags: `--authHeaderName` and `--authHeaderValue`:
>
> ```bash
> ./capsule-http \
> --url=${DOWNLOAD_URL} \
> --authHeaderName="PRIVATE-TOKEN" \
> --authHeaderValue="${TOKEN}" \
> --wasm=${WASM_FILE} \
> --httpPort=${HTTP_PORT}
> ```

## Monitoring the service

Capsule HTTP server exposes a REST API (`/metrics` endpoint) that can be used to monitor the service. You have to write the logic to generate the metrics and expose them to Prometheus.

An HTTP request to the `/metrics` endpoint will trigger the call of the **exposed** `OnMetrics` function of the WASM module. You need to implement the `OnMetrics` function in your WASM module:

```golang
//export OnMetrics
func OnMetrics() uint64 {

	// Generate OpenText Prometheus metric
	counterMetrics := []string{
		"# HELP call counter",
		"# TYPE call_counter counter",
		"call_counter " + strconv.Itoa(counter)}

	response := capsule.HTTPResponse{
		TextBody:   strings.Join(counterMetrics, "\n"),
		Headers:    `{"Content-Type": "text/plain; charset=utf-8"}`,
		StatusCode: 200,
	}
	return capsule.Success([]byte(capsule.StringifyHTTPResponse(response)))

}
```
> - Don't forget to expose the function: `//export OnMetrics`
> - You can find a complete sample here: [hello-world sample](https://github.com/bots-garden/capsule/blob/main/capsule-http/functions/hello-world/main.go)

## Health Check

Capsule HTTP server exposes a REST API (`/health` endpoint) that can be used to teturn a health status. You have to write the logic to generate the status.

An HTTP request to the `/health` endpoint will trigger the call of the **exposed** `OnHealthCheck` function of the WASM module. You need to implement the `OnHealthCheck` function in your WASM module:

```golang
//export OnHealthCheck
func OnHealthCheck() uint64 {

	response := capsule.HTTPResponse{
		JSONBody:   `{"message": "OK"}`,
		Headers:    `{"Content-Type": "application/json; charset=utf-8"}`,
		StatusCode: 200,
	}

	return capsule.Success([]byte(capsule.StringifyHTTPResponse(response)))
}
```
> - Don't forget to expose the function: `//export OnHealthCheck`
> - You can find a complete sample here: [hello-world sample](https://github.com/bots-garden/capsule/blob/main/capsule-http/functions/hello-world/main.go)

## `OnStart` and `OnStop` functions

You can add a `OnStart` and `OnStop` function to the WASM module. These functions will be called when the service starts and stops.

```golang
//export OnStart
func OnStart() {
	capsule.Print("🚗 OnStart")
}

//export OnStop
func OnStop() {
	capsule.Print("🚙 OnStop")
}
```
> - Don't forget to expose the functions
> - You can find a complete sample here: [hello-world sample](https://github.com/bots-garden/capsule/blob/main/capsule-http/functions/hello-world/main.go)


## Develop a WASM Capsule module

Have a look to these samples:

- [Capsule MDK documentation: first HTTP module](https://bots-garden.github.io/capsule-module-sdk/first-http-module/)
- [capsule-http/functions](https://github.com/bots-garden/capsule/tree/main/capsule-http/functions)
- [Samples of the Capsule MDK](https://github.com/bots-garden/capsule-module-sdk/tree/main/samples)
