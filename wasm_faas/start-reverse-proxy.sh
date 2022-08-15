#!/bin/bash
cd ../capsulelauncher
#DEBUG="true"
CAPSULE_REVERSE_PROXY_ADMIN_TOKEN="1234567890" \
go run main.go \
   -mode=reverse-proxy \
   -backend="memory" \
   -httpPort=8888
