#!/bin/bash
hey -n 5000 -c 50 -m POST \
-H "Content-Type: application/json" \
-d '{"message": "Golang 💚💜 wasm"}' \
"http://localhost:7070"

#  Requests/sec: 236.2816
