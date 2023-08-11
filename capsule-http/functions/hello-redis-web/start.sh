#!/bin/bash
REDIS_URI="" \
./capsule-http --wasm=./hello-redis-web.wasm --httpPort=8080
