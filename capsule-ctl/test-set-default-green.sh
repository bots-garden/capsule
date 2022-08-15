#!/bin/bash
# This is the CLI capsulectl
CAPSULE_WORKER_URL="http://localhost:9999" \
CAPSULE_REVERSE_PROXY_URL="http://localhost:8888" \
CAPSULE_BACKEND="memory" \
CAPSULE_WORKER_ADMIN_TOKEN="0987654321" \
go run main.go set-default \
-function=hello \
-revision=green

