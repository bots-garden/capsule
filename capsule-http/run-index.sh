#!/bin/bash
# Load the wasm module from a local location
./capsule-http \
   -wasm=../functions/index-html/index.wasm \
   -httpPort=8080
