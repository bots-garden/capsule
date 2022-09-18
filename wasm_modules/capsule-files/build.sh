#!/bin/bash
tinygo build -o hello.wasm -scheduler=none -target wasi ./hello.go

ls -lh *.wasm
