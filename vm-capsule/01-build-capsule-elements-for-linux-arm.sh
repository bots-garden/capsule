#!/bin/bash
cd ../capsule-launcher
env GOOS=linux GOARCH=arm64 go build -o capsule_arm64
cp capsule_arm64 ../vm-capsule/apps/capsule
rm capsule_arm64

cd ../capsule-registry
env GOOS=linux GOARCH=arm64 go build -o capsule-registry_arm64
cp capsule-registry_arm64 ../vm-capsule/apps/capsule-registry
rm capsule-registry_arm64

cd ../capsule-reverse-proxy
env GOOS=linux GOARCH=arm64 go build -o capsule-reverse-proxy_arm64
cp capsule-reverse-proxy_arm64 ../vm-capsule/apps/capsule-reverse-proxy
rm capsule-reverse-proxy_arm64

cd ../capsule-worker
env GOOS=linux GOARCH=arm64 go build -o capsule-worker_arm64
cp capsule-worker_arm64 ../vm-capsule/apps/capsule-worker
rm capsule-worker_arm64
