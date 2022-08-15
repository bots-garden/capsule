#!/bin/bash
cd ../../../capsule-launcher

go run main.go \
   -wasm=../wasm_modules/reverse-proxy-demo/capsule-hola-red/hola.wasm \
   -mode=http \
   -httpPort=6062
