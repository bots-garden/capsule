#!/bin/bash
cd ../capsule-launcher
env GOOS=linux GOARCH=arm64 go build -o capsule_arm64
cp capsule_arm64 ../vm-capsule/apps/capsule
rm capsule_arm64

