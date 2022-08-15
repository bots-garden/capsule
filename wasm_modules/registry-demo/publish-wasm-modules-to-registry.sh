#!/bin/bash

curl -X POST http://localhost:4999/upload/k33g/hello/0.0.0 \
  -F "file=@./capsule-hello/hello.wasm" \
  -F "info=hello function from @k33g" \
  -H "Content-Type: multipart/form-data"

curl -X POST http://localhost:4999/upload/k33g/hey/0.0.0 \
  -F "file=@./capsule-hey/hey.wasm" \
  -F "info=hello hey from @k33g" \
  -H "Content-Type: multipart/form-data"
