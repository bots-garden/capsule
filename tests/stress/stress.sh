#!/bin/bash
# -n 10000 -c 1000
hey -n 10000 -c 50 -m POST \
-H "Content-Type: application/json" \
-d '{"message": "Golang 💚💜 wasm"}' \
"http://localhost:7070"

