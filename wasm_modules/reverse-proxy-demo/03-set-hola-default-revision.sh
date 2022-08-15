#!/bin/bash

curl -v -X POST \
  http://localhost:8888/memory/functions/hola/revision \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"revision": "default", "url": "http://localhost:6061"}'
echo ""

