#!/bin/bash
cd ../capsule-ctl
CAPSULE_REVERSE_PROXY_URL="http://localhost:8888" \
CAPSULE_WORKER_URL="http://localhost:9999" \
CAPSULE_BACKEND="memory" \
go run main.go reverse-proxy

