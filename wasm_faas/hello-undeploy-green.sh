#!/bin/bash
# This is the CLI capsulectl
cd ../capsule-ctl
CAPSULE_WORKER_URL="http://localhost:9999" \
CAPSULE_BACKEND="memory" \
CAPSULE_WORKER_ADMIN_TOKEN="0987654321" \
go run main.go un-deploy \
-function=hello \
-revision=green

