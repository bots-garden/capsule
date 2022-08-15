#!/bin/bash
cd ../../../capsule-launcher

export MESSAGE="💊 Capsule is Amazing 😍"
go run main.go \
   -wasm=../wasm_modules/registry-demo/capsule-hey/hey.wasm \
   -mode=http \
   -httpPort=9092
