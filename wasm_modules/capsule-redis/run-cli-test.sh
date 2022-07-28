#!/bin/bash
cd ../../capsulelauncher
REDIS_ADDR="localhost:6379" \
REDIS_PWD="" \
go run main.go \
   -wasm=../wasm_modules/capsule-redis/hello.wasm \
   -mode=cli
