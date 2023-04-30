#!/bin/bash
# Load the wasm module from a local location
#./capsule-http \
go run main.go \
   -wasm=./functions/hey/hey.wasm \
   -httpPort=8080
