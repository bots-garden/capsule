#!/bin/bash
# -n 10000 -c 1000
#hey -n 10 -c 5 -m POST \
hey -n 100 -c 5 -m POST \
-H "Content-Type: application/json" \
-d '{"message": "Golang 💚💜 wasm"}' \
"http://localhost:8888/functions/hello"

