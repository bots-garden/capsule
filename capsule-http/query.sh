#!/bin/bash

JSON_DATA='{"name":"Bob Morane","age":42}'

curl -X POST http://localhost:8080 \
          -H 'Content-Type: application/json; charset=utf-8' \
          -d "${JSON_DATA}"

#curl -X POST https://e627-37-169-191-158.ngrok-free.app \
#          -H 'Content-Type: application/json; charset=utf-8' \
#          -d "${JSON_DATA}"

#curl -X POST http://localhost:8080/functions/yo/blue

curl -X POST http://localhost:38771/

42723

curl -X POST http://localhost:42723 \
          -H 'Content-Type: application/json; charset=utf-8' \
          -d '{"name":"Bob Morane","age":42}'


curl -X POST http://0.0.0.0:8080/functions/yo/blue \
          -H 'Content-Type: application/json; charset=utf-8' \
          -d '{"name":"Bob Morane","age":42}'