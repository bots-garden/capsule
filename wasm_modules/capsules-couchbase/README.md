# Capsule CouchBase

## Prerequisites

### Couchbase

You need a running Couchbase server with data:
```bash
cd couchbase-setup
./01-install-couchbase.sh
./02-create-couchbase-cluster.sh
./03-create-couchbase-bucket.sh
./04-insert-documents.sh
```

> - Get the url of the CouchBase web admin panel: `gp url 8091`
> - Admin user: `admin`, password: `ilovepandas`

### Wasm modules

To build the CLI wasm module:
```bash
cd cli-module
./build-module.sh
```

To build the HTTP wasm module
```bash
cd http-module
./build-module.sh
```

## Run the demos

### CLI version

This wasm module is used by the `cli` mode of the capsule launcher:

```bash
COUCHBASE_CLUSTER="couchbase://localhost" \
COUCHBASE_USER="admin" \
COUCHBASE_PWD="ilovepandas" \
COUCHBASE_BUCKET="wasm-data" \
./capsule \
   -wasm=./cli-module/hello.wasm \
   -mode=cli
```

### HTTP version

This wasm module is used by the `http` mode of the capsule launcher:
```bash
COUCHBASE_CLUSTER="couchbase://localhost" \
COUCHBASE_USER="admin" \
COUCHBASE_PWD="ilovepandas" \
COUCHBASE_BUCKET="wasm-data" \
./capsule \
   -wasm=./http-module/hello.wasm \
   -mode=http \
   -httpPort=7070
```

Then launch an HTTP request:
```bash
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{}'
```
You should get something like this:
```json
[{"docs":{"name":"this is an info","type":"info"}},{"docs":{"name":"this is another info","type":"info"}},{"docs":{"name":"👋 hello world 🌍","type":"message"}},{"docs":{"name":"👋 greetings 🎉","type":"message"}}]
```


