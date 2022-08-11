#!/bin/bash

curl -v -X POST \
http://localhost:9999/functions/set_default_revision \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hola",
    "revision": "current"
}
EOF
echo ""

