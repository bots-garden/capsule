#!/bin/bash
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"message": "TinyGo 💜 wasm", "author": "@k33g_org"}'
echo ""
