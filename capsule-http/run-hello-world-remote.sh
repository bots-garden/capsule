#!/bin/bash
# Load the wasm module from a remote location
./capsule-http \
   -url=http://localhost:5000/hello-world.wasm \
   -wasm=./tmp/hello-world.wasm \
   -httpPort=8080
