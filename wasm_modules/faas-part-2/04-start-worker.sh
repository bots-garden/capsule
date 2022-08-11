#!/bin/bash

# bash -c "exec -a <MyProcessName> <Command>"
cd ../../capsulelauncher

go run main.go \
   -mode=worker \
   -reverseProxy=http://localhost:8888 \
   -backend=memory \
   -capsulePath=./capsule \
   -httpPortCounter=10000 \
   -httpPort=9999
