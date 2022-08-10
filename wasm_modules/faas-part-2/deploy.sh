#!/bin/bash

curl -v -X POST \
http://localhost:9999/functions/deploy \
-H 'content-type: application/json; charset=utf-8' \
-d '{"function": "hello", "revision": "default", "downloadUrl": "http://localhost:4999/k33g/hello/0.0.0/hello.wasm"}'
echo ""

curl -v -X POST \
http://localhost:9999/functions/deploy \
-H 'content-type: application/json; charset=utf-8' \
-d '{"function": "hey", "revision": "default", "downloadUrl": "http://localhost:4999/k33g/hey/0.0.0/hey.wasm"}'
echo ""


# other revision: new revision for an existing function
curl -v -X POST \
http://localhost:9999/functions/deploy \
-H 'content-type: application/json; charset=utf-8' \
-d '{"function": "hey", "revision": "blue", "downloadUrl": "http://localhost:4999/k33g/hey/0.0.1/hey.wasm"}'
echo ""

# scale (same revision): add a running module to an existing revision
curl -v -X POST \
http://localhost:9999/functions/deploy \
-H 'content-type: application/json; charset=utf-8' \
-d '{"function": "hey", "revision": "default", "downloadUrl": "http://localhost:4999/k33g/hey/0.0.0/hey.wasm"}'
echo ""
