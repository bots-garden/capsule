#!/bin/bash
#hey -n 5000 -c 50 -m POST \
hey -n 300 -c 100 -m POST \
-H "Content-Type: application/json" \
-d '{"message": "Golang 💚💜 wasm", "author": "Philippe"}' \
"http://localhost:7070"

# Requests/sec: 485.4062 deps: updates wazero to 1.0.0-pre.9
