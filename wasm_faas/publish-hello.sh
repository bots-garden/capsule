#!/bin/bash
cd ../capsule-ctl
CAPSULE_REGISTRY_ADMIN_TOKEN="AZERTYUIOP" \
go run main.go publish \
-wasmFile=../wasm_faas/hello/hello.wasm -wasmInfo=wip \
-wasmOrg=k33g -wasmName=hello -wasmTag=0.0.0 \
-registryUrl=http://localhost:4999
