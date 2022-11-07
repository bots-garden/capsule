#!/bin/bash
cd ../../capsule-launcher
REDIS_ADDR="localhost:6379" \
REDIS_PWD="" \
DEBUG=true MESSAGE="💊 Capsule Rocks 🚀" go run main.go \
   -wasm=../wasm_modules/test-everything/hello.wasm \
   -mode=http \
   -httpPort=7070
