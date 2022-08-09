#!/bin/bash

# bash -c "exec -a <MyProcessName> <Command>"

cd ./src/hola-000
./build.sh

cd ../hola-001
./build.sh

cd ..

# Publish the wasm modules to the registry
curl -X POST http://localhost:4999/upload/k33g/hola/0.0.0 \
  -F "file=@./hola-000/hola.wasm" \
  -F "info=hola 0.0.0 function from @k33g" \
  -H "Content-Type: multipart/form-data"

curl -X POST http://localhost:4999/upload/k33g/hola/0.0.1 \
  -F "file=@./hola-001/hola.wasm" \
  -F "info=hola 0.0.1 function from @k33g" \
  -H "Content-Type: multipart/form-data"
