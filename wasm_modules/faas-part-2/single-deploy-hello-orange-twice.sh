#!/bin/bash

curl -v -X POST \
http://localhost:9999/functions/deploy \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hello",
    "revision": "orange",
    "downloadUrl": "http://localhost:4999/k33g/hello/0.0.0/hello.wasm",
    "envVariables": {
        "MESSAGE": "🚀 this is a message",
        "TOKEN": "😍 this is not a header token"
    }
}
EOF
echo ""

curl -v -X POST \
http://localhost:9999/functions/deploy \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hello",
    "revision": "orange",
    "downloadUrl": "http://localhost:4999/k33g/hello/0.0.0/hello.wasm",
    "envVariables": {
        "MESSAGE": "🖐 this is a message",
        "TOKEN": "😁 this is not a header token"
    }
}
EOF
echo ""
