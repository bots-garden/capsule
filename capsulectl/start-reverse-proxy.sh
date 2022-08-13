#!/bin/bash
cd ../capsulelauncher
go run main.go \
   -mode=reverse-proxy \
   -backend="memory" \
   -httpPort=8888
