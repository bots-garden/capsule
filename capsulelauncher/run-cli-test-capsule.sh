#!/bin/bash
MESSAGE="🎉 Hello World" go run main.go \
   -wasm=../wasm_modules/capsule-function-template/hello.wasm \
   -mode=cli \
   "👋 hello world 🌍🎃" 1234 "Bob Morane"
