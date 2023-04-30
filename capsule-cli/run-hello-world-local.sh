#!/bin/bash
# Load the wasm module from a local location
#./capsule \
export MESSAGE="Hello, World!"
go run main.go \
   --wasm=./functions/hello/hello.wasm \
   --params="Bob Morane, 42, 🤗"