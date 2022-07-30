#!/bin/bash
cd ../../../capsulelauncher
COUCHBASE_CLUSTER="couchbase://localhost" \
COUCHBASE_USER="admin" \
COUCHBASE_PWD="ilovepandas" \
COUCHBASE_BUCKET="wasm-data" \
go run main.go \
   -wasm=../wasm_modules/capsules-couchbase/cli-module/hello.wasm \
   -mode=cli
