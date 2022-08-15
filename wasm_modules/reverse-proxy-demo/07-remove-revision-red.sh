#!/bin/bash
curl -v -X DELETE \
  http://localhost:8888/memory/functions/hola/revision \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"revision": "red"}'
echo ""
