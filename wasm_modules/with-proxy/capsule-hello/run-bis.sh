#!/bin/bash
cd ../../../capsulelauncher

export MESSAGE="💊 Capsule Rocks 🚀"
go run main.go \
   -wasm=../wasm_modules/with-proxy/capsule-hello/hello.wasm \
   -mode=http \
   -httpPort=7071
