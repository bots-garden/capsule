#!/bin/bash

JSON_DATA='{"name":"Bob Morane","age":42}'
HTTP_PORT='8090'
DOMAIN='localhost'
PROTOCOL='http'

curl -X POST ${PROTOCOL}://${DOMAIN}:${HTTP_PORT} \
  -H 'Content-Type: application/json; charset=utf-8' \
  -d "${JSON_DATA}"
