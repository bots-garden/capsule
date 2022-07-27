#!/bin/bash
go run main.go \
   -wasm=./wasm_modules/capsule-http/hello.wasm \
   -mode=cli \
   -param="👋 hello world 🌍🎃"
