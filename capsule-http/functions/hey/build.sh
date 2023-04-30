#!/bin/bash
tinygo build -o hey.wasm -scheduler=none --no-debug -target wasi ./main.go 

ls -lh *.wasm
