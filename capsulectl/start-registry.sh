#!/bin/bash
DOWNLOADED_FILES_PATH="${PWD}/registry/functions"
cd ../capsulelauncher
CAPSULE_REGISTRY_ADMIN_TOKEN="AZERTYUIOP" \
go run main.go \
   -mode=registry \
   -files="${DOWNLOADED_FILES_PATH}" \
   -httpPort=4999
