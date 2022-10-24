#!/bin/bash
cd ../../capsule-launcher

DEBUG=true MESSAGE="💊 Capsule Rocks 🚀" go run main.go \
   -wasm=../wasm_modules/capsule-hello-post-2/hello.wasm \
   -mode=http-next \
   -httpPort=7070
