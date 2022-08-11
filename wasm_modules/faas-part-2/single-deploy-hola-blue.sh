#!/bin/bash

curl -v -X POST \
http://localhost:9999/functions/deploy \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hola",
    "revision": "blue",
    "downloadUrl": "http://localhost:4999/k33g/hola/0.0.1/hola.wasm",
    "envVariables": {
        "MESSAGE": "🔵 Blue revision of Hola",
        "TOKEN": "this is not a header token"
    }
}
EOF
echo ""

