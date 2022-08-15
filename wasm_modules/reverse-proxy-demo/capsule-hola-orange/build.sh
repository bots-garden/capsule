#!/bin/bash
tinygo build -o hola.wasm -scheduler=none -target wasi ./hola.go

ls -lh *.wasm
