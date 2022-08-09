#!/bin/bash

# before serve the wasm file: python3 -m http.server 8080
cd ../../capsulelauncher

export MESSAGE="💊 Capsule Rocks 🚀"
go run main.go \
   -wasm=./tmp/hola-orange.wasm \
   -url="http://localhost:4999/k33g/hola/orange/hola.wasm" \
   -mode=http \
   -httpPort=7071
