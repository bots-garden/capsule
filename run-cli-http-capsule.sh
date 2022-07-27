#!/bin/bash
go run main.go \
   -wasm=./wasm_modules/capsule-http-get/hello.wasm \
   -mode=cli \
   -param="[GET]👋 hello world 🌍"


go run main.go \
   -wasm=./wasm_modules/capsule-http-post/hello.wasm \
   -mode=cli \
   -param="[POST]👋 hello world 🌍"
