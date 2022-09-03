#!/bin/bash
go mod tidy
tinygo build -o yo.wasm -scheduler=none -target wasi ./yo.go

ls -lh *.wasm
