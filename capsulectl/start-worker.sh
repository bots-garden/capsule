#!/bin/bash
cd ../capsulelauncher
# build the capsule executable
./build.sh
cp capsule ../capsulectl/capsule

go run main.go \
   -mode=worker \
   -reverseProxy=http://localhost:8888 \
   -backend=memory \
   -capsulePath=./capsule \
   -httpPortCounter=10000 \
   -httpPort=9999

# check where to store tmp wasm downloade files
