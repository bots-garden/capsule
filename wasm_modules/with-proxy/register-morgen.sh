#!/bin/bash
curl -v -X POST \
  http://localhost:8888/memory/functions/registration \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"function": "morgen", "revision": "default", "url": "http://localhost:5050"}'

echo ""
