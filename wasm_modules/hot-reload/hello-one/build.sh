#!/bin/bash
go mod tidy
tinygo build -o hello-one.wasm -scheduler=none -target wasi ./hello.go

ls -lh *.wasm
