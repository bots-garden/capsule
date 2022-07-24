#!/bin/bash
go run main.go \
  -wasm=./wasm_modules/capsule-function-template/hello.wasm \
  -mode=http \
  -httpPort=7070