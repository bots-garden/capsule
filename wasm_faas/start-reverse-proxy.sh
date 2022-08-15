#!/bin/bash
cd ../capsule-reverse-proxy
#DEBUG="true"
CAPSULE_REVERSE_PROXY_ADMIN_TOKEN="1234567890" \
go run main.go \
   -backend="memory" \
   -httpPort=8888
