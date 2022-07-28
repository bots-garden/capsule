#!/bin/bash
MESSAGE="🎉 Hello World" go run main.go \
   -wasm=../wasm_modules/capsule-function-template/hello.wasm \
   -mode=cli \
   -param="👋 hello world 🌍🎃"
