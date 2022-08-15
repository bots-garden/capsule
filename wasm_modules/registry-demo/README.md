# Capsule Wasm Registry

> 🚧 wip

## Start the registry

```bash
DOWNLOADED_FILES_PATH="${PWD}/registry/functions"
echo "${DOWNLOADED_FILES_PATH}"

./capsule-registry \
   -files="${DOWNLOADED_FILES_PATH}" \
   -httpPort=4999
```

## Publish wasm modules to the registry

```bash
curl -X POST http://localhost:4999/upload/k33g/hello/0.0.0 \
  -F "file=@./capsule-hello/hello.wasm" \
  -F "info=hello function from @k33g" \
  -H "Content-Type: multipart/form-data"

curl -X POST http://localhost:4999/upload/k33g/hey/0.0.0 \
  -F "file=@./capsule-hey/hey.wasm" \
  -F "info=hello hey from @k33g" \
  -H "Content-Type: multipart/form-data"
```

## Load and serve the modules

```bash
MESSAGE="💊 Capsule Rocks 🚀" ./capsule \
   -wasm=./tmp/hello.wasm \
   -url="http://localhost:4999/k33g/hello/0.0.0/hello.wasm" \
   -mode=http \
   -httpPort=9091

MESSAGE="💊 Capsule Rocks 🚀" ./capsule \
   -wasm=./tmp/hey.wasm \
   -url="http://localhost:4999/k33g/hey/0.0.0/hey.wasm" \
   -mode=http \
   -httpPort=9092
```

## Call the functions

```bash
curl -v -X POST \
  http://localhost:9091 \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"message": "Golang 💚💜 wasm", "author": "Bob Morane"}'
```

```bash
curl -v -X POST \
  http://localhost:9092 \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"message": "Golang 💚💜 wasm", "author": "Bill Murray"}'
```
