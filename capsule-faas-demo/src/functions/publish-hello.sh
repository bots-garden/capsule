#!/bin/bash
# Publish the wasm module to the registry
# 🖐 change the tag if you publish a new version
curl -X POST http://localhost:4999/upload/k33g/hello/0.0.0 \
  -F "file=@./hello/hello.wasm" \
  -F "info=hello function v0.0.0 from @k33g [GET]" \
  -H "Content-Type: multipart/form-data"
