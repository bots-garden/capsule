#!/bin/bash
cd ../../capsule-launcher

DEBUG=true go run main.go \
   -wasm=../wasm_modules/capsule-hello-get/hello.wasm \
   -mode=http \
   -httpPort=7070
