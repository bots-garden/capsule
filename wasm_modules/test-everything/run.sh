#!/bin/bash
cd ../../capsule-launcher

DEBUG=true MESSAGE="💊 Capsule Rocks 🚀" go run main.go \
   -wasm=../wasm_modules/test-everything/hello.wasm \
   -mode=http \
   -httpPort=7070
