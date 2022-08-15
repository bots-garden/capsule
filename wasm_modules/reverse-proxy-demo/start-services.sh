#!/bin/bash

# bash -c "exec -a <MyProcessName> <Command>"

cd ../../capsule-launcher

bash -c "exec -a hola_orange go run main.go \
-wasm=../wasm_modules/reverse-proxy-demo/capsule-hola-orange/hola.wasm \
-mode=http \
-httpPort=6061" &

bash -c "exec -a hola_red go run main.go \
-wasm=../wasm_modules/reverse-proxy-demo/capsule-hola-red/hola.wasm \
-mode=http \
-httpPort=6062" &

bash -c "exec -a hola_yellow go run main.go \
-wasm=../wasm_modules/reverse-proxy-demo/capsule-hola-yellow/hola.wasm \
-mode=http \
-httpPort=6063" &
