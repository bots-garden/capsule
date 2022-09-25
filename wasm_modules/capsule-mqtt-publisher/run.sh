#!/bin/bash
cd ../../capsule-launcher

go run main.go \
   -wasm=../wasm_modules/capsule-mqtt-publisher/hello.wasm \
   -mode=cli
