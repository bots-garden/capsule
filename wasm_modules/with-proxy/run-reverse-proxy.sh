#!/bin/bash
cd ../../capsulelauncher

go run main.go \
   -mode=reverse-proxy \
   -httpPort=8888
