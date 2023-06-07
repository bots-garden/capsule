# 🚀 Getting Started

## Use the Capsule HTTP server

First, download the last version of the Capsule HTTP server for the appropriate OS & ARCH (and release version):

```bash
VERSION="v0.3.9" OS="linux" ARCH="arm64"
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

## Develop a WASM Capsule module

Have a look to these samples:

- [Capsule MDK documentation: first HTTP module](https://bots-garden.github.io/capsule-module-sdk/first-http-module/)
- [capsule-http/functions](https://github.com/bots-garden/capsule/tree/main/capsule-http/functions)
- [Samples of the Capsule MDK](https://github.com/bots-garden/capsule-module-sdk/tree/main/samples)
