#!/bin/bash
cd ../capsule-ctl
go build -o caps
cp caps ../vm-capsule
rm caps
