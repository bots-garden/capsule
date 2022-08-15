#!/bin/bash

cd ../../capsule-launcher

export MESSAGE="💊 Capsule Rocks 🚀"
go run main.go \
   -wasm=./tmp/hey.wasm \
   -url="http://localhost:4999/k33g/hey/0.0.0/hey.wasm" \
   -mode=http \
   -httpPort=9092


