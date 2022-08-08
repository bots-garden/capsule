#!/bin/bash
cd ../../capsulelauncher

go run main.go \
   -mode=reverse-proxy \
   -config=../wasm_modules/with-proxy/config.yaml \
   -backend="memory" \
   -httpPort=8888
