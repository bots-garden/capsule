#!/bin/bash
echo "http://localhost:8888/functions/hola/orange"
echo "http://localhost:8888/functions/hola/yellow"
echo "http://localhost:8888/functions/hola/red"
echo "http://localhost:8888/functions/hola/default"
echo "http://localhost:8888/functions/hola"

cd ../../capsule-reverse-proxy

go run main.go \
    -backend="memory" \
    -httpPort=8888
