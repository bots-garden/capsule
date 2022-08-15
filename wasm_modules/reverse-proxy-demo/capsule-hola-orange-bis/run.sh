#!/bin/bash
cd ../../../capsule-launcher

go run main.go \
   -wasm=../wasm_modules/reverse-proxy-demo/capsule-hola-orange-bis/hola.wasm \
   -mode=http \
   -httpPort=6064
