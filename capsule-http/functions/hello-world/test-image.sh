#!/bin/bash
IMAGE_NAME="demo-capsule-http"
docker run \
  -p 8080:8080 \
  --rm ${IMAGE_NAME}
  