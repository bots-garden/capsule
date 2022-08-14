#!/bin/bash
# This is the CLI capsulectl
CAPSULE_WORKER_URL="http://localhost:9999" \
CAPSULE_REVERSE_PROXY_URL="http://localhost:8888" \
CAPSULE_BACKEND="memory" \
go run main.go deploy \
-function=hello \
-revision=blue \
-downloadUrl=http://localhost:4999/k33g/hello/0.0.0/hello.wasm \
-envVariables='{"MESSAGE": "Revision 🔵","TOKEN": "👩‍🔧🧑‍🔧👨‍🔧"}'

CAPSULE_WORKER_URL="http://localhost:9999" \
CAPSULE_REVERSE_PROXY_URL="http://localhost:8888" \
CAPSULE_BACKEND="memory" \
go run main.go deploy \
-function=hello \
-revision=green \
-downloadUrl=http://localhost:4999/k33g/hello/0.0.0/hello.wasm \
-envVariables='{"MESSAGE": "Revision 🟢","TOKEN": "🍏🥝🍉"}'

CAPSULE_WORKER_URL="http://localhost:9999" \
CAPSULE_REVERSE_PROXY_URL="http://localhost:8888" \
CAPSULE_BACKEND="memory" \
go run main.go deploy \
-function=hello \
-revision=orange \
-downloadUrl=http://localhost:4999/k33g/hello/0.0.0/hello.wasm \
-envVariables='{"MESSAGE": "Revision 🟠","TOKEN": "😡"}'

CAPSULE_WORKER_URL="http://localhost:9999" \
CAPSULE_REVERSE_PROXY_URL="http://localhost:8888" \
CAPSULE_BACKEND="memory" \
go run main.go deploy \
-function=hello \
-revision=orange \
-downloadUrl=http://localhost:4999/k33g/hello/0.0.0/hello.wasm \
-envVariables='{"MESSAGE": "Revision 🟠","TOKEN": "😡🤬"}'

CAPSULE_WORKER_URL="http://localhost:9999" \
CAPSULE_REVERSE_PROXY_URL="http://localhost:8888" \
CAPSULE_BACKEND="memory" \
go run main.go deploy \
-function=hello \
-revision=orange \
-downloadUrl=http://localhost:4999/k33g/hello/0.0.0/hello.wasm \
-envVariables='{"MESSAGE": "Revision 🟠","TOKEN": "😡🤬🥵"}'
