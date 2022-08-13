#!/bin/bash
capsule \
   -mode=reverse-proxy \
   -backend="memory" \
   -httpPort=8888
