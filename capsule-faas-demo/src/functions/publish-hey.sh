#!/bin/bash
# Publish the wasm module to the registry
# 🖐 change the tag if you publish a new version
curl -X POST http://localhost:4999/upload/k33g/hey/0.0.0 \
  -F "file=@./hey/hey.wasm" \
  -F "info=hey function v0.0.0 from @k33g [POST]" \
  -H "Content-Type: multipart/form-data"
