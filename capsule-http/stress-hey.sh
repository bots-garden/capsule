#!/bin/bash
hey -n 300 -c 100 -m POST \
-H "Content-Type: text/plain; charset=utf-8" \
-d 'Bob Morane' http://localhost:8080
