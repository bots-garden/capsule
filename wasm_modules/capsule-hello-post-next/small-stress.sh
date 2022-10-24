#!/bin/bash
hey -n 50 -c 2 -m POST \
-H "Content-Type: application/json" \
-d '{"message": "Golang 💚💜 wasm", "author": "Philippe"}' \
"http://localhost:7070"

