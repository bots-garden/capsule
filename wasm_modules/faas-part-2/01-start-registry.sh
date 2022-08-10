#!/bin/bash

# bash -c "exec -a <MyProcessName> <Command>"
cd ../../capsulelauncher

go run main.go \
   -mode=registry \
   -files="/Users/k33g/Documents/capsule/wasm_modules/faas-part-2/functions" \
   -httpPort=4999
