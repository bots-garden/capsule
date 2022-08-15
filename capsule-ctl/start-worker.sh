#!/bin/bash
cd ../capsulelauncher
# build the capsule executable
./build.sh
cp capsule ../capsulectl/capsule

#DEBUG="true"
CAPSULE_REVERSE_PROXY_ADMIN_TOKEN="1234567890" \
CAPSULE_WORKER_ADMIN_TOKEN="0987654321" \
go run main.go \
   -mode=worker \
   -reverseProxy=http://localhost:8888 \
   -backend=memory \
   -capsulePath=./capsule \
   -httpPortCounter=10000 \
   -httpPort=9999

# check where to store tmp wasm downloade files
