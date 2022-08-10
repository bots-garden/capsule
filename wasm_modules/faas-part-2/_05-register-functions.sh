#!/bin/bash

# bash -c "exec -a <MyProcessName> <Command>"

# Register the function with the 0.0.0 revision
curl -v -X POST \
  http://localhost:8888/memory/functions/registration \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"function": "hola", "revision": "0.0.0", "url": "http://localhost:7070"}'

# Add the default revision
curl -v -X POST \
  http://localhost:8888/memory/functions/hola/revision \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"function": "hola", "revision": "default", "url": "http://localhost:7070"}'


# Add the 0.0.1 revision
curl -v -X POST \
  http://localhost:8888/memory/functions/hola/revision \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"function": "hola", "revision": "0.0.1", "url": "http://localhost:7071"}'

