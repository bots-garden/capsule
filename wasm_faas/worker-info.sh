#!/bin/bash
cd ../capsule-ctl
CAPSULE_WORKER_URL="http://localhost:9999" \
CAPSULE_BACKEND="memory" \
go run main.go worker
