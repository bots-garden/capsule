#!/bin/bash

export MESSAGE="🖐 good morning 😄"
go run main.go \
   -url=http://localhost:4999/k33g/hello/0.0.0/hello.wasm \
   -wasm=./tmp/hello.wasm \
   -mode=http \
   -httpPort=7070
