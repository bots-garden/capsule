#!/bin/bash
hey -n 5000 -c 50 -m POST \
-H "Content-Type: application/json" \
-d '{"message": "Golang 💚💜 wasm", "author": "Philippe"}' \
"http://localhost:7070"

#   Requests/sec: 241.3615
