#!/bin/bash

# before serve the wasm file: python3 -m http.server 8080
cd ../../capsule-launcher

go run main.go \
   -wasm=./tmp/hello.wasm \
   -url="http://localhost:8080/hello.wasm" \
   -mode=http \
   -httpPort=7070
