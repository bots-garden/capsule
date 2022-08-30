#!/bin/bash

export MESSAGE="🖐 good morning 😄"
go run main.go \
   -url=http://localhost:4999/k33g/oups/0.0.0/oups.wasm \
   -wasm=./tmp/hello.wasm \
   -mode=http \
   -httpPort=7070
