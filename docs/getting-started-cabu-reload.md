# 🚀 Getting Started

## Reload a deployed wasm module function
> Introduced in Capsule v0.2.9 🦜 [parrot]

You can reload the running wasm module by using the `/load-wasm-module` route of **Capsule**, like this:

```bash
curl -v -X POST \
  http://localhost:7070/load-wasm-module \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"url": "http://localhost:9090/hello-two/hello-two.wasm", "path": "./tmp/hello-two.wasm"}'
```

## Protect `/load-wasm-module` with a token

To protect the route, start capsule with the `CAPSULE_RELOAD_TOKEN` variable:

```bash
CAPSULE_RELOAD_TOKEN="ilovepandas" capsule \
   -url=http://localhost:9090/hello-one/hello-one.wasm \
   -wasm=./tmp/hello-one.wasm \
   -mode=http \
   -httpPort=7070
```

Then, add the token to the headers of the request:

```bash
curl -v -X POST \
  http://localhost:7070/load-wasm-module \
  -H 'content-type: application/json; charset=utf-8' \
  -H 'CAPSULE_RELOAD_TOKEN: ilovepandas' \
  -d '{"url": "http://localhost:9090/hello-two/hello-two.wasm", "path": "./tmp/hello-two.wasm"}'
```
