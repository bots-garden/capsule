#!/bin/bash
cd ../../capsule-launcher

go run main.go \
   -wasm=../wasm_modules/capsule-nats-other-subscriber/hello.wasm \
   -mode=nats \
   -natssrv=nats.devsecops.fun:4222 \
   -subject=notify

