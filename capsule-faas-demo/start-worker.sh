#!/bin/bash
capsule \
   -mode=worker \
   -reverseProxy=http://localhost:8888 \
   -backend=memory \
   -capsulePath=capsule \
   -httpPortCounter=10000 \
   -httpPort=9999
