#!/bin/bash
curl -v -X DELETE \
  http://localhost:8888/memory/functions/morgen/blue/url \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"url": "http://localhost:5053"}'
echo ""
