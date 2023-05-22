#!/bin/bash

JSON_DATA='{"name":"Bob Morane","age":42}'

curl -X POST http://localhost:8080/functions/hello-world/default/0 \
          -H 'Content-Type: application/json; charset=utf-8' \
          -d "${JSON_DATA}"
echo "------"

curl -X POST http://localhost:8080/functions/hello-world/default/1 \
          -H 'Content-Type: application/json; charset=utf-8' \
          -d "${JSON_DATA}"
echo "------"

curl -X POST http://localhost:8080/functions/hello-world/default \
          -H 'Content-Type: application/json; charset=utf-8' \
          -d "${JSON_DATA}"
echo "------"

curl -X POST http://localhost:8080/functions/hello-world \
          -H 'Content-Type: application/json; charset=utf-8' \
          -d "${JSON_DATA}"
echo "------"

#curl -X GET http://localhost:8080/functions/hello-world/default/1
