#!/bin/bash
# -n 10000 -c 1000
hey -n 100 -c 4 -m POST \
-H "Content-Type: application/json" \
-d '{"message": "Golang 💚💜 wasm"}' \
"http://localhost:7070"

