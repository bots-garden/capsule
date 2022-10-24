#!/bin/bash
cd ../../capsule-launcher

DEBUG=true MESSAGE="💊 Capsule Rocks 🚀" go run main.go \
   -wasm=../wasm_modules/capsule-hello-post-next/hello.wasm \
   -mode=http \
   -httpPort=7070
