#!/bin/bash
cd ../../capsulelauncher
COUCHBASE_CLUSTER="couchbases://127.0.0.1" \
COUCHBASE_USER="admin" \
COUCHBASE_PWD="ilovepandas" \
go run main.go \
   -wasm=../wasm_modules/capsule-couchbase/hello.wasm \
   -mode=cli

