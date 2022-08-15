#!/bin/bash
# This is the CLI capsulectl
cd ../capsule-ctl

CAPSULE_WORKER_URL="http://localhost:9999" \
CAPSULE_BACKEND="memory" \
CAPSULE_WORKER_ADMIN_TOKEN="0987654321" \
go run main.go deploy \
-function=hello \
-revision=blue \
-downloadUrl=http://localhost:4999/k33g/hello/0.0.0/hello.wasm \
-envVariables='{"MESSAGE": "Revision 🔵","TOKEN": "👩‍🔧🧑‍🔧👨‍🔧"}'

CAPSULE_WORKER_URL="http://localhost:9999" \
CAPSULE_BACKEND="memory" \
CAPSULE_WORKER_ADMIN_TOKEN="0987654321" \
go run main.go deploy \
-function=hello \
-revision=green \
-downloadUrl=http://localhost:4999/k33g/hello/0.0.0/hello.wasm \
-envVariables='{"MESSAGE": "Revision 🟢","TOKEN": "🍏🥝🍉"}'

CAPSULE_WORKER_URL="http://localhost:9999" \
CAPSULE_BACKEND="memory" \
CAPSULE_WORKER_ADMIN_TOKEN="0987654321" \
go run main.go deploy \
-function=hello \
-revision=orange \
-downloadUrl=http://localhost:4999/k33g/hello/0.0.0/hello.wasm \
-envVariables='{"MESSAGE": "Revision 🟠","TOKEN": "😡"}'

CAPSULE_WORKER_URL="http://localhost:9999" \
CAPSULE_BACKEND="memory" \
CAPSULE_WORKER_ADMIN_TOKEN="0987654321" \
go run main.go deploy \
-function=hello \
-revision=orange \
-downloadUrl=http://localhost:4999/k33g/hello/0.0.0/hello.wasm \
-envVariables='{"MESSAGE": "Revision 🟠","TOKEN": "😡🤬"}'

CAPSULE_WORKER_URL="http://localhost:9999" \
CAPSULE_BACKEND="memory" \
CAPSULE_WORKER_ADMIN_TOKEN="0987654321" \
go run main.go deploy \
-function=hello \
-revision=orange \
-downloadUrl=http://localhost:4999/k33g/hello/0.0.0/hello.wasm \
-envVariables='{"MESSAGE": "Revision 🟠","TOKEN": "😡🤬🥵"}'
