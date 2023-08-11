#!/bin/bash
tinygo build -o redis-db.wasm \
    -scheduler=none \
    --no-debug \
    -target wasi ./main.go 
ls -lh *.wasm

echo "📦 Building capsule-cli..."
cd ../..
go build -ldflags="-s -w" -o capsule
ls -lh capsule
mv capsule ./functions/redis-db
