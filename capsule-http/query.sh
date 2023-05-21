#!/bin/bash

JSON_DATA='{"name":"Bob Morane","age":42}'

curl -X POST http://localhost:8080 \
          -H 'Content-Type: application/json; charset=utf-8' \
          -d "${JSON_DATA}"

#curl -X POST https://e627-37-169-191-158.ngrok-free.app \
#          -H 'Content-Type: application/json; charset=utf-8' \
#          -d "${JSON_DATA}"

