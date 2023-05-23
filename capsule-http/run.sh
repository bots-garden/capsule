#!/bin/bash
task build-capsule-http build-hello-world build-index-html

#unset NGROK_AUTHTOKEN
WASM_FILE='./functions/hello-world/hello-world.wasm'


#NGROK_AUTH_TOKEN="${NGROK_AUTHTOKEN}" \
./capsule-http --wasm=${WASM_FILE} --httpPort=8080

#./capsule-http --wasm=${WASM_FILE} --httpPort=8080

