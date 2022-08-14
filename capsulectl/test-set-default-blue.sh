#!/bin/bash
# This is the CLI capsulectl
CAPSULE_WORKER_URL="http://localhost:9999" \
CAPSULE_REVERSE_PROXY_URL="http://localhost:8888" \
CAPSULE_BACKEND="memory" \
go run main.go set-default \
-function=hello \
-revision=blue

