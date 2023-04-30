#!/bin/bash
tinygo build -o hello.wasm -scheduler=none --no-debug -target wasi ./main.go 

ls -lh *.wasm
