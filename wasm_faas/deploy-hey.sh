#!/bin/bash
# This is the CLI capsulectl
cd ../capsule-ctl
CAPSULE_WORKER_URL="http://localhost:9999" \
CAPSULE_BACKEND="memory" \
CAPSULE_WORKER_ADMIN_TOKEN="0987654321" \
go run main.go deploy \
-function=hey \
-revision=demo \
-downloadUrl=http://localhost:4999/k33g/hey/0.0.0/hey.wasm

#CAPSULE_REVERSE_PROXY_URL="http://localhost:8888" \
#CAPSULE_REVERSE_PROXY_ADMIN_TOKEN="1234567890" \
