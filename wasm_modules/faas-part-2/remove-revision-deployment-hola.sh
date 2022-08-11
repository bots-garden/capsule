#!/bin/bash

curl -v -X DELETE \
http://localhost:9999/functions/revisions/deployments \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hola",
    "revision": "current"
}
EOF
echo ""

