#!/bin/bash
CAPSULE_REGISTRY_ADMIN_TOKEN="AZERTYUIOP" \
go run main.go publish \
-wasmFile=./hey/hey.wasm -wasmInfo=wip \
-wasmOrg=k33g -wasmName=hey -wasmTag=0.0.0 \
-registryUrl=http://localhost:4999
