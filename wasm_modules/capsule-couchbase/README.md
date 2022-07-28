# Capsule Redis

This wasm module is used by the `cli` mode

```bash
#!/bin/bash
COUCHBASE_CLUSTER="couchbases://127.0.0.1" \
COUCHBASE_USER="admin" \
COUCHBASE_PWD="ilovepandas" \
go run main.go \
   -wasm=../wasm_modules/capsule-couchbase/hello.wasm \
   -mode=cli
```
