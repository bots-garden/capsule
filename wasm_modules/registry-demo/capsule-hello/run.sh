#!/bin/bash
cd ../../../capsule-launcher

export MESSAGE="💊 Capsule Rocks 🚀"
go run main.go \
   -wasm=../wasm_modules/registry-demo/capsule-hello/hello.wasm \
   -mode=http \
   -httpPort=9091
