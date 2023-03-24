#!/bin/bash
# -n 10000 -c 1000
#hey -n 10 -c 5 -m POST \
hey -n 300 -c 100 -m POST \
-H "Content-Type: application/json" \
-d '{"message": "Golang 💚💜 wasm"}' \
"http://localhost:7070"

# hey -n 300 -c 100 -m POST \
# Requests/sec: 30.5280
# Requests/sec: 32.1274
# Requests/sec: 31.5995
# Requests/sec: 117.4794 with wazero.NewRuntimeConfigInterpreter()
# Requests/sec: 123.6958 wasi_snapshot_preview1.MustInstantiate(ctx, wasmRuntime)
# Requests/sec: 84.0781 deps: updates wazero to 1.0.0-pre.4
# Requests/sec: 167.1686 deps: updates wazero to 1.0.0-pre.9
