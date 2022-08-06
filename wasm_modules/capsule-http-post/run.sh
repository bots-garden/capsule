#!/bin/bash
cd ../../capsulelauncher

go run main.go \
   -wasm=../wasm_modules/capsule-http-post/hello.wasm \
   -mode=cli \
   "[POST]👋 hello world 🌍"

