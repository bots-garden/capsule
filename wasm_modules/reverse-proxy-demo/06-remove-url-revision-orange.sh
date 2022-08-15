#!/bin/bash
curl -v -X DELETE \
  http://localhost:8888/memory/functions/hola/orange/url \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"url": "http://localhost:6064"}'
echo ""

curl -v -X DELETE \
  http://localhost:8888/memory/functions/hola/default/url \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"url": "http://localhost:6064"}'
echo ""
