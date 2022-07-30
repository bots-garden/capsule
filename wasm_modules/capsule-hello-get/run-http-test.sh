#!/bin/bash
cd ../../capsulelauncher

export MESSAGE="💊 Capsule Rocks 🚀"
go run main.go \
   -wasm=../wasm_modules/capsule-hello-get/hello.wasm \
   -mode=http \
   -httpPort=7070
