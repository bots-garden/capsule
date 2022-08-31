#!/bin/bash
CAPSULE_REGISTRY_ADMIN_TOKEN="AZERTYUIOP" \
./caps publish \
-wasmFile="./src/functions/yo/yo.wasm" -wasmInfo="this is the yo module" \
-wasmOrg="k33g" -wasmName="yo" -wasmTag="0.0.0" \
-registryUrl="https://capsdev.devsecops.fun:4999"




