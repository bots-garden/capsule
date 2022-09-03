#!/bin/bash
go mod tidy
tinygo build -o hey.wasm -scheduler=none -target wasi ./hey.go
ls -lh *.wasm
