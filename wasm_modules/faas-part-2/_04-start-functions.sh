#!/bin/bash

# bash -c "exec -a <MyProcessName> <Command>"

# Download, then start the wasm module
#capsule \
#   -wasm=./tmp/hola.wasm \
#   -url=http://localhost:4999/k33g/hola/0.0.0/hola.wasm \
#   -mode=http \
#   -httpPort=7070

cd ../../capsulelauncher


bash -c "exec -a hola-000 go run main.go \
-wasm=./tmp/hola-0.0.0.wasm \
-url=http://localhost:4999/k33g/hola/0.0.0/hola.wasm \
-mode=http \
-httpPort=7070" &

# Download, then start the wasm module
#capsule \
#   -wasm=./tmp/hola.wasm \
#   -url="http://localhost:4999/k33g/hola/0.0.1/hola.wasm" \
#   -mode=http \
#   -httpPort=7071

bash -c "exec -a hola001 go run main.go \
-wasm=./tmp/hola-0.0.1.wasm \
-url=http://localhost:4999/k33g/hola/0.0.1/hola.wasm \
-mode=http \
-httpPort=7071" &
