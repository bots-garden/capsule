#!/bin/bash

# bash -c "exec -a <MyProcessName> <Command>"
cd ../../capsulelauncher

go run main.go \
   -mode=reverse-proxy \
   -backend="memory" \
   -httpPort=8888

#    -config=../wasm_modules/faas-part-1/config.yaml \
