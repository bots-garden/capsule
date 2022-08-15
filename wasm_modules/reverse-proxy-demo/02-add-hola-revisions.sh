#!/bin/bash

curl -v -X POST \
  http://localhost:8888/memory/functions/hola/revision \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"revision": "red", "url": "http://localhost:6062"}'
echo ""

curl -v -X POST \
  http://localhost:8888/memory/functions/hola/revision \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"revision": "yellow", "url": "http://localhost:6063"}'
echo ""
