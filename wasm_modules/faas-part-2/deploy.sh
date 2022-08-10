#!/bin/bash

curl -v -X POST \
http://localhost:9999/functions/deploy \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hello",
    "revision": "default",
    "downloadUrl": "http://localhost:4999/k33g/hello/0.0.0/hello.wasm",
    "envVariables": {
        "MESSAGE": "this is a message",
        "TOKEN": "this is not a header token"
    }
}
EOF
#-d '{"function": "hello", "revision": "default", "downloadUrl": "http://localhost:4999/k33g/hello/0.0.0/hello.wasm"}'
echo ""


curl -v -X POST \
http://localhost:9999/functions/deploy \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hey",
    "revision": "default",
    "downloadUrl": "http://localhost:4999/k33g/hey/0.0.0/hey.wasm",
    "envVariables": {
        "MESSAGE": "DEFAULT REVISION"
    }
}
EOF
#-d '{"function": "hey", "revision": "default", "downloadUrl": "http://localhost:4999/k33g/hey/0.0.0/hey.wasm"}'
echo ""


# other revision: new revision for an existing function
curl -v -X POST \
http://localhost:9999/functions/deploy \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hey",
    "revision": "blue",
    "downloadUrl": "http://localhost:4999/k33g/hey/0.0.1/hey.wasm",
    "envVariables": {
        "MESSAGE": "BLUE REVISION"
    }
}
EOF
#-d '{"function": "hey", "revision": "blue", "downloadUrl": "http://localhost:4999/k33g/hey/0.0.1/hey.wasm"}'
echo ""

# scale (same revision): add a running module to an existing revision
curl -v -X POST \
http://localhost:9999/functions/deploy \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hey",
    "revision": "default",
    "downloadUrl": "http://localhost:4999/k33g/hey/0.0.0/hey.wasm",
    "envVariables": {
        "MESSAGE": "DEFAULT REVISION (AGAIN)"
    }
}
EOF
#-d '{"function": "hey", "revision": "default", "downloadUrl": "http://localhost:4999/k33g/hey/0.0.0/hey.wasm"}'
echo ""
