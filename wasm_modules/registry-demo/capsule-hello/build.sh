#!/bin/bash
go mod tidy
tinygo build -o hello.wasm -scheduler=none -target wasi ./hello.go

ls -lh *.wasm
