#!/bin/bash

JSON_DATA='{"name":"Bob Morane","age":42}'

hey -n 300 -c 100 -m POST \
-H "Content-Type: application/json; charset=utf-8" \
-d "${JSON_DATA}" http://localhost:8080/functions/hello-world/default
