#!/bin/bash
cd ../../../capsule-launcher

export MESSAGE="💊 Capsule is Awesome 💚"
go run main.go \
   -wasm=../wasm_modules/reverse-proxy-demo/capsule-hola-yellow/hola.wasm \
   -mode=http \
   -httpPort=6063
