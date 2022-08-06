#!/bin/bash
curl -v -X POST \
  http://localhost:8888/functions/hey \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"message": "Golang 💚💜 wasm", "author": "Philippe"}'
