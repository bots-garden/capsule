#!/bin/bash
cd ../../capsulelauncher

go run main.go \
   -mode=registry \
   -files="/Users/k33g/Documents/capsule/wasm_modules/registry-demo/wasm_modules" \
   -httpPort=4999

