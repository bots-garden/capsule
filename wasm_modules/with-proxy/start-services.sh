#!/bin/bash

# bash -c "exec -a <MyProcessName> <Command>"

cd ../../capsulelauncher

MESSAGE="💊 Capsule Rocks 🚀" bash -c "exec -a hello go run main.go \
-wasm=../wasm_modules/with-proxy/capsule-hello/hello.wasm \
-mode=http \
-httpPort=9091" &

MESSAGE="💊 Capsule Rocks 🚀" bash -c "exec -a hello go run main.go \
-wasm=../wasm_modules/with-proxy/capsule-hello/hello.wasm \
-mode=http \
-httpPort=7071" &

# pkill -f hello

MESSAGE="💊 Capsule is Amazing 😍" bash -c "exec -a hey go run main.go \
-wasm=../wasm_modules/with-proxy/capsule-hey/hey.wasm \
-mode=http \
-httpPort=9092" &

MESSAGE="💊 Capsule is Awesome 💚" bash -c "exec -a hola go run main.go \
-wasm=../wasm_modules/with-proxy/capsule-hola/hola.wasm \
-mode=http \
-httpPort=9093" &

MESSAGE="💊 Capsule is Awesome 💚" bash -c "exec -a hola_orange go run main.go \
-wasm=../wasm_modules/with-proxy/capsule-hola-orange/hola.wasm \
-mode=http \
-httpPort=6061" &

MESSAGE="💊 Capsule is Awesome 💚" bash -c "exec -a hola_yellow go run main.go \
-wasm=../wasm_modules/with-proxy/capsule-hola-yellow/hola.wasm \
-mode=http \
-httpPort=6062" &
