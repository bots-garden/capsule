#!/bin/bash
cd ../../../capsulelauncher

export MESSAGE="💊 Capsule is Amazing 😍"
go run main.go \
   -wasm=../wasm_modules/with-proxy/capsule-hey/hey.wasm \
   -mode=http \
   -httpPort=9092
