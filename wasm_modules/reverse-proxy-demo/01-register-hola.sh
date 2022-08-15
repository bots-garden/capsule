#!/bin/bash
# Always register a function giving a revision (the revision will be created)
curl -v -X POST \
  http://localhost:8888/memory/functions/registration \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"function": "hola", "revision": "orange", "url": "http://localhost:6061"}'

echo ""
