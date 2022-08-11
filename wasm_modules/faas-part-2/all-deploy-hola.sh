#!/bin/bash

curl -v -X POST \
http://localhost:9999/functions/deploy \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hola",
    "revision": "current",
    "downloadUrl": "http://localhost:4999/k33g/hola/0.0.0/hola.wasm",
    "envVariables": {
        "MESSAGE": "1️⃣🟣 Current revision of Hola",
        "TOKEN": "this is not a header token"
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
    "function": "hola",
    "revision": "current",
    "downloadUrl": "http://localhost:4999/k33g/hola/0.0.0/hola.wasm",
    "envVariables": {
        "MESSAGE": "2️⃣🟣 Current revision of Hola",
        "TOKEN": "this is not a header token"
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

curl -v -X POST \
http://localhost:9999/functions/deploy \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hola",
    "revision": "green",
    "downloadUrl": "http://localhost:4999/k33g/hola/0.0.2/hola.wasm",
    "envVariables": {
        "MESSAGE": "🟢 Green revision of Hola",
        "TOKEN": "this is not a header token"
    }
}
EOF
echo ""
