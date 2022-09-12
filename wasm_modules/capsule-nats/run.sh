#!/bin/bash
cd ../../capsule-launcher

NATS_SRV="nats.devsecops.fun:4222" \
go run main.go \
   -wasm=../wasm_modules/capsule-nats/hello.wasm \
   -mode=cli
