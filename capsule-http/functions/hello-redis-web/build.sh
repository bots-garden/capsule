#!/bin/bash
tinygo build -o hello-redis-web.wasm \
    -scheduler=none \
    --no-debug \
    -target wasi ./main.go 
ls -lh *.wasm

echo "📦 Building capsule-http..."
cd ../..
go build -ldflags="-s -w" -o capsule-http
ls -lh capsule-http
mv capsule-http ./functions/hello-redis-web