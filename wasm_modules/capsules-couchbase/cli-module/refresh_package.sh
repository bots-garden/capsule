#!/bin/bash
rm ./go.sum
rm ./go.mod

echo "module github.com/bots-garden/capsule/wasm_modules/capsule-couchbase/cli-module" > go.mod
echo "" >> go.mod
echo "go 1.18" >> go.mod

go mod tidy


