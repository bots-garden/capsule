#!/bin/bash
hey -n 300 -c 100 -m POST \
-H "Content-Type: application/json; charset=utf-8" \
-d '{"name":"Bob Morane","age":42}' http://localhost:8080
