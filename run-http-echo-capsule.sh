#!/bin/bash
go run main.go \
   -wasm=./wasm_modules/capsule-function-template/hello.wasm \
   -mode=http-echo \
   -httpPort=7070

# -wasm=./wasm_modules/capsule-function-template/hello.wasm \
# -wasm=./wasm_modules/capsule-http/hello.wasm \
