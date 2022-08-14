#!/bin/bash
curl -v -X POST \
  http://localhost:8888/functions/hey/demo \
  -H 'content-type: application/json; charset=utf-8' \
  -H 'my-token: I love Pandas' \
  -d '{"message": "Golang 💚💜 wasm", "author": "k33g"}'
echo ""
