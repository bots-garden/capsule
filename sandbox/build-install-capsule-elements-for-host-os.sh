#!/bin/bash
sudo rm /usr/local/bin/caps*

cd ../capsule-launcher
go build -o capsule
sudo cp capsule /usr/local/bin
rm capsule

cd ../capsule-registry
go build -o capsule-registry
sudo cp capsule-registry /usr/local/bin
rm capsule-registry

cd ../capsule-reverse-proxy
go build -o capsule-reverse-proxy
sudo cp capsule-reverse-proxy /usr/local/bin
rm capsule-reverse-proxy

cd ../capsule-worker
go build -o capsule-worker
sudo cp capsule-worker /usr/local/bin
rm capsule-worker

cd ../capsule-ctl
go build -o caps
sudo cp caps /usr/local/bin
rm caps

