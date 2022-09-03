#!/bin/bash
cd ../../capsule-launcher
REDIS_ADDR="localhost:6379" \
REDIS_PWD="" \
go run main.go \
   -wasm=../wasm_modules/capsule-redis-keys/hello.wasm \
   -mode=cli
