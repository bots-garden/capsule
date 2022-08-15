# Capsule hello

This wasm module is used by the `http` mode

## Build the wasm module

```bash
tinygo build -o hello.wasm -scheduler=none -target wasi ./hello.go
```

## Load (and run) the wasm file module from a local path

```bash
export MESSAGE="💊 Capsule Rocks 🚀"
./capsule \
   -wasm=./hello.wasm \
   -mode=http \
   -httpPort=7070
```

Then call the wasm function:
```bash
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"message": "Golang 💚 wasm", "author": "Philippe"}'
```

## Load (and run) the wasm file module from an URL

First serve the wasm file:
```bash
python3 -m http.server 8080
```

Then load and serve the module: *(the `wasm` file is the output of the download file)*
```bash
export MESSAGE="💊 Capsule Rocks 🚀"
./capsule \
   -wasm=./tmp/hello.wasm \
   -url="http://localhost:8080/hello.wasm" \
   -mode=http \
   -httpPort=7070
```

Then call the wasm function:
```bash
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"message": "Golang 💚 wasm", "author": "Philippe"}'
```
