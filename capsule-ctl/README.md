# Capsule CLI, aka `cc`

## Publish a wasm file to the Capsule registry

```bash
CAPSULE_REGISTRY_ADMIN_TOKEN="AZERTYUIOP" \
./caps publish \
-wasmFile=./hey/hey.wasm -wasmInfo=wip \
-wasmOrg=k33g -wasmName=hey -wasmTag=0.0.0 \
-registryUrl=http://localhost:4999
```

## Deploy a function (with a revision) to the worker

```bash
CAPSULE_WORKER_URL="http://localhost:9999" \
CAPSULE_BACKEND="memory" \
CAPSULE_WORKER_ADMIN_TOKEN="0987654321" \
./caps deploy \
-function=hey \
-revision=demo \
-downloadUrl=http://localhost:4999/k33g/hey/0.0.0/hey.wasm
```
> 🖐 if you deploy the same revision, you scale the function

## Set an existing revision as the default revision

```bash
CAPSULE_WORKER_URL="http://localhost:9999" \
CAPSULE_BACKEND="memory" \
CAPSULE_WORKER_ADMIN_TOKEN="0987654321" \
./caps set-default \
-function=hello \
-revision=orange
```

## Remove the default revision of a function

```bash
CAPSULE_WORKER_URL="http://localhost:9999" \
CAPSULE_BACKEND="memory" \
CAPSULE_WORKER_ADMIN_TOKEN="0987654321" \
./caps unset-default \
-function=hello
```

## Un-deploy the revision of a function

```bash
CAPSULE_WORKER_URL="http://localhost:9999" \
CAPSULE_BACKEND="memory" \
CAPSULE_WORKER_ADMIN_TOKEN="0987654321" \
./caps un-deploy \
-function=hello \
-revision=green
```
