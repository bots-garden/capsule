# 🚀 Getting Started

## Use the Capsule HTTP server

First, download the last version of the Capsule HTTP server for the appropriate OS & ARCH:

```bash
VERSION="v0.3.5" OS="linux" ARCH="arm64"
wget -O capsule-http https://github.com/bots-garden/capsule/releases/download/${VERSION}/capsule-http-${VERSION}-${OS}-${ARCH}
chmod +x capsule-http
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

## Monitoring the service

Capsule HTTP server exposes a REST API that can be used to monitor the service. It's a Prometheus route. You only need to call the `/metrics` endpoint.

> This feature is provided thanks to the [FiberPrometheus](https://github.com/ansrivas/fiberprometheus) library.

Following metrics are available by default:

- `http_requests_total`
- `http_request_duration_seconds`
- `http_requests_in_progress_total`

## Develop a WASM Capsule module

Have a look to these samples:

- [Capsule MDK documentation: first HTTP module](https://bots-garden.github.io/capsule-module-sdk/first-http-module/)
- [capsule-http/functions](https://github.com/bots-garden/capsule/tree/main/capsule-http/functions)
- [Samples of the Capsule MDK](https://github.com/bots-garden/capsule-module-sdk/tree/main/samples)
