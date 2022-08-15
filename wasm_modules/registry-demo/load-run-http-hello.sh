#!/bin/bash

cd ../../capsule-launcher

export MESSAGE="💊 Capsule Rocks 🚀"
go run main.go \
   -wasm=./tmp/hello.wasm \
   -url="http://localhost:4999/k33g/hello/0.0.0/hello.wasm" \
   -mode=http \
   -httpPort=9091

