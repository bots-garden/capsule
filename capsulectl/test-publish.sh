#!/bin/bash
go run main.go publish \
-wasmFile=./hello/hello.wasm -wasmInfo=wip \
-wasmOrg=k33g -wasmName=hello -wasmTag=0.0.1 \
-registryUrl=http://localhost:4999 \
-registryToken=nothing
