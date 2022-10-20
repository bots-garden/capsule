#!/bin/bash
go mod tidy
tinygo build -o hello-two.wasm -scheduler=none -target wasi ./hello.go

ls -lh *.wasm
