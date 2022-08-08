#!/bin/bash
curl -v -X POST \
  http://localhost:8888/memory/functions/morgen/revision \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"revision": "blue", "url": "http://localhost:5051"}'
echo ""
