#!/bin/bash
go mod tidy
tinygo build -o hola.wasm -scheduler=none -target wasi ./hola.go

ls -lh *.wasm
