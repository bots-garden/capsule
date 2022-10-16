#!/bin/bash
curl -v -X POST \
  http://localhost:7070/load-wasm-module \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"url": "http://localhost:9090/hello-two/hello-two.wasm", "path": "./tmp/hello-two.wasm"}'
echo ""
