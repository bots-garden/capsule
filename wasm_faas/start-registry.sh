#!/bin/bash
DOWNLOADED_FILES_PATH="${PWD}/registry/functions"
cd ../capsule-registry
CAPSULE_REGISTRY_ADMIN_TOKEN="AZERTYUIOP" \
go run main.go \
   -files="${DOWNLOADED_FILES_PATH}" \
   -httpPort=4999
