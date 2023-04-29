#!/bin/bash
# Load the wasm module from a local location
./capsule-http \
   -wasm=../functions/hey/hey.wasm \
   -httpPort=8080
