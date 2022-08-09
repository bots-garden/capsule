#!/bin/bash

curl -X POST http://localhost:4999/upload/k33g/hello/0.0.0 \
  -F "file=@../with-proxy/capsule-hello/hello.wasm" \
  -F "info=hello function from @k33g" \
  -H "Content-Type: multipart/form-data"

curl -X POST http://localhost:4999/upload/k33g/hey/0.0.0 \
  -F "file=@../with-proxy/capsule-hey/hey.wasm" \
  -F "info=hello hey from @k33g" \
  -H "Content-Type: multipart/form-data"

curl -X POST http://localhost:4999/upload/k33g/hola/0.0.0 \
  -F "file=@../with-proxy/capsule-hola/hola.wasm" \
  -F "info=hola function from @k33g" \
  -H "Content-Type: multipart/form-data"

curl -X POST http://localhost:4999/upload/k33g/hola/orange \
  -F "file=@../with-proxy/capsule-hola-orange/hola.wasm" \
  -F "info=hola(orange) function from @k33g" \
  -H "Content-Type: multipart/form-data"

curl -X POST http://localhost:4999/upload/k33g/hola/yellow \
  -F "file=@../with-proxy/capsule-hola-yellow/hola.wasm" \
  -F "info=hola(yellow) function from @k33g" \
  -H "Content-Type: multipart/form-data"

