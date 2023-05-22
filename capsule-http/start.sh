#!/bin/bash
# Start a function

FUNCTION_NAME="hello-world"
FUNCTION_REVISION="default"
WASM_FILE="./functions/hello-world/hello-world.wasm"
curl -X POST localhost:8080/functions/start \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"${FUNCTION_NAME}",
    "revision":"${FUNCTION_REVISION}",
    "description":"this is a description",
    "path":"./capsule-http",
    "args": ["", "-wasm=${WASM_FILE}"],
    "env": []
}
EOF