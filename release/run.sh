#!/bin/bash
HTTP_PORT="8090"
DOMAIN='localhost'
PROTOCOL='http'
WASM_FILE='./hello-world.wasm'
./capsule-http-v0.3.6-darwin-arm64 --wasm=${WASM_FILE} --httpPort=${HTTP_PORT}
