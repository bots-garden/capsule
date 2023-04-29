#!/bin/bash
# Load the wasm module from a local location
./capsule-http \
   -wasm=../functions/hello-world/hello-world.wasm \
   -httpPort=8080
