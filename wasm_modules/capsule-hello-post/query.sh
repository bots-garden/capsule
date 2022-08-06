#!/bin/bash
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"message": "Golang 💚💜 wasm", "author": "Philippe"}'
