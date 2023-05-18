#!/bin/bash
: <<'COMMENT'
GitLab registry
https://gitlab.com/wasmkitchen/wasm.registry/-/packages

COMMENT

GITLAB_WASM_PROJECT_ID="45020306"
WASM_PACKAGE="capsule"
WASM_MODULE="hello-world.wasm"
WASM_VERSION="0.0.0"
WASM_FILE="./hello-world.wasm"

echo "📦 ${WASM_PACKAGE}"
echo "📝 ${WASM_MODULE} ${WASM_VERSION}"

curl --header "PRIVATE-TOKEN: ${GITLAB_WASM_TOKEN}" \
     --upload-file ${WASM_FILE} \
     "https://gitlab.com/api/v4/projects/${GITLAB_WASM_PROJECT_ID}/packages/generic/${WASM_PACKAGE}/${WASM_VERSION}/${WASM_MODULE}"

echo ""
echo "🌍 https://gitlab.com/bots-garden/wasm-registry/-/packages"
