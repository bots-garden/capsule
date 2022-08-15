#!/bin/bash
DOWNLOADED_FILES_PATH="${PWD}/registry/functions"
echo "${DOWNLOADED_FILES_PATH}"
cd ../../capsule-registry

go run main.go \
   -files="${DOWNLOADED_FILES_PATH}" \
   -httpPort=4999
