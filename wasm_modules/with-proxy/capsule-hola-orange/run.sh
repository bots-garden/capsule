#!/bin/bash
cd ../../../capsulelauncher

export MESSAGE="💊 Capsule is Awesome 💚"
go run main.go \
   -wasm=../wasm_modules/with-proxy/capsule-hola-orange/hola.wasm \
   -mode=http \
   -httpPort=6061
