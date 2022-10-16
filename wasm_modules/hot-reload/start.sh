#!/bin/bash
cd ../../capsule-launcher
go run main.go \
   -url=http://localhost:9090/hello-one/hello-one.wasm \
   -wasm=./tmp/hello-one.wasm \
   -mode=http \
   -httpPort=7070
