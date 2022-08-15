#!/bin/bash
curl -v -X DELETE \
  http://localhost:8888/memory/functions/registration \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"function": "hola"}'
echo ""
