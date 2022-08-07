#!/bin/bash
cd ../../capsulelauncher

go run main.go \
   -mode=reverse-proxy \
   -config=../wasm_modules/with-proxy/config.yaml \
   -httpPort=8888
