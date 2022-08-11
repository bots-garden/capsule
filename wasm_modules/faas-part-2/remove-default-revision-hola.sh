#!/bin/bash

curl -v -X DELETE \
http://localhost:9999/functions/remove_default_revision \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hola"
}
EOF
echo ""

