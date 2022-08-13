#!/bin/bash

# Remove default revision if it exists
curl -v -X DELETE \
http://localhost:9999/functions/remove_default_revision \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hello"
}
EOF
echo ""

# Now the orange revision is the default revision
curl -v -X POST \
http://localhost:9999/functions/set_default_revision \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hello",
    "revision": "orange"
}
EOF
echo ""
