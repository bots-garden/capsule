#!/bin/bash
CAPSULE_MAIN_PROCESS_URL="http://localhost:8080"
CAPSULE_INSTALL_PATH="/usr/local/bin/capsule-http" # on the target platform
GITLAB_WASM_PROJECT_ID="45020306"

FUNCTION_NAME="hello"
WASM_PACKAGE="capsule"

WASM_FILES_LOCATION="/home/ubuntu/workspaces/capsule/capsule-http/tests/faas/functions"

# Download wasm file from the GitLab generic package registry
WASM_VERSION="blue"
WASM_MODULE="${FUNCTION_NAME}-${WASM_VERSION}.wasm"
GITLAB_REGISTRY_URL="https://gitlab.com/api/v4/projects/${GITLAB_WASM_PROJECT_ID}/packages/generic/${WASM_PACKAGE}/${WASM_VERSION}/${WASM_MODULE}"
SAVE_DIRECTORY="/home/ubuntu/workspaces/capsule/capsule-http/tests/faas/tmp/${WASM_MODULE}"
go run main.go --cmd=start \
    --name=${FUNCTION_NAME} \
	--revision=${WASM_VERSION} \
	--description="hello blue deployment" \
	--path=${CAPSULE_INSTALL_PATH} \
	--stopAfter=30 \
	--env='["MESSAGE=hello blue"]'\
    --wasm=${SAVE_DIRECTORY} \
	--url=${GITLAB_REGISTRY_URL}


# The wasm file is already available locally
WASM_VERSION="orange"
WASM_MODULE="${FUNCTION_NAME}-${WASM_VERSION}.wasm"

go run main.go --cmd=start \
    --name=${FUNCTION_NAME} \
	--revision=${WASM_VERSION} \
	--description="hello orange deployment" \
	--path=${CAPSULE_INSTALL_PATH} \
	--stopAfter=30 \
	--env='["MESSAGE=hello orange"]'\
    --wasm=${WASM_FILES_LOCATION}/${FUNCTION_NAME}-${WASM_VERSION}/${WASM_MODULE}

# The wasm file is already available locally
WASM_VERSION="green"
WASM_MODULE="${FUNCTION_NAME}-${WASM_VERSION}.wasm"

go run main.go --cmd=start \
    --name=${FUNCTION_NAME} \
	--revision=${WASM_VERSION} \
	--description="hello green deployment" \
	--path=${CAPSULE_INSTALL_PATH} \
	--env='["MESSAGE=hello green"]'\
    --wasm=${WASM_FILES_LOCATION}/${FUNCTION_NAME}-${WASM_VERSION}/${WASM_MODULE}

# The wasm file is already available locally
WASM_VERSION="default"
WASM_MODULE="${FUNCTION_NAME}-${WASM_VERSION}.wasm"

go run main.go --cmd=start \
    --name=${FUNCTION_NAME} \
	--revision=${WASM_VERSION} \
	--description="hello default deployment" \
	--path=${CAPSULE_INSTALL_PATH} \
	--env='["MESSAGE=hello default"]'\
    --wasm=${WASM_FILES_LOCATION}/${FUNCTION_NAME}-${WASM_VERSION}/${WASM_MODULE}

