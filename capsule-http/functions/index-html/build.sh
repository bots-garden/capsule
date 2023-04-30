#!/bin/bash
tinygo build -o index.wasm -scheduler=none --no-debug -target wasi ./main.go 

ls -lh *.wasm
