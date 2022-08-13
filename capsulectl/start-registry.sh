#!/bin/bash
DOWNLOADED_FILES_PATH="${PWD}/registry/functions"
cd ../capsulelauncher
go run main.go \
   -mode=registry \
   -files="${DOWNLOADED_FILES_PATH}" \
   -httpPort=4999
