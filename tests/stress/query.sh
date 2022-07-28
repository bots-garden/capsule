#!/bin/bash
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json' \
  -H 'admin-token: AABBCCDDEEFF' \
  -d '{"message": "Golang 💚 wasm"}'
