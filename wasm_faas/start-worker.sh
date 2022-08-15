#!/bin/bash
cd ../capsule-launcher
# build the capsule executable
./build.sh
cp capsule ../capsule-worker/capsule

cd ../capsule-worker
#DEBUG="true"
CAPSULE_REVERSE_PROXY_ADMIN_TOKEN="1234567890" \
CAPSULE_WORKER_ADMIN_TOKEN="0987654321" \
go run main.go \
   -reverseProxy=http://localhost:8888 \
   -backend=memory \
   -capsulePath=./capsule \
   -httpPortCounter=10000 \
   -httpPort=9999

# check where to store tmp wasm downloaded files
