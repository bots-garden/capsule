#!/bin/bash

# bash -c "exec -a <MyProcessName> <Command>"

cd ./src/hello-default
./build.sh

cd ../hola-blue
./build.sh

cd ../hola-default
./build.sh

cd ../hola-green
./build.sh

cd ..

# Publish the wasm modules to the registry
curl -X POST http://localhost:4999/upload/k33g/hello/default \
  -F "file=@./hello-default/hello.wasm" \
  -F "info=hola default function from @k33g" \
  -H "Content-Type: multipart/form-data"

curl -X POST http://localhost:4999/upload/k33g/hola/default \
  -F "file=@./hola-default/hola.wasm" \
  -F "info=hola default function from @k33g" \
  -H "Content-Type: multipart/form-data"

curl -X POST http://localhost:4999/upload/k33g/hola/blue \
  -F "file=@./hola-blue/hola.wasm" \
  -F "info=hola blue function from @k33g" \
  -H "Content-Type: multipart/form-data"

curl -X POST http://localhost:4999/upload/k33g/hola/green \
  -F "file=@./hola-green/hola.wasm" \
  -F "info=hola green function from @k33g" \
  -H "Content-Type: multipart/form-data"
