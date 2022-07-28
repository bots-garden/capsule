#!/bin/bash
export MESSAGE="💊 Capsule Rocks 🚀"
go run main.go \
   -wasm=./wasm_modules/capsule-hello/hello.wasm \
   -mode=http \
   -httpPort=7070

# -wasm=./wasm_modules/capsule-function-template/hello.wasm \
# -wasm=./wasm_modules/capsule-http/hello.wasm \
