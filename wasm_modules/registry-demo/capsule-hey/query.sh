#!/bin/bash
curl -v -X POST \
  http://localhost:9092 \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"message": "Golang 💚💜 wasm", "author": "Philippe"}'
