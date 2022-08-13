#!/bin/bash
# This is the CLI capsulectl
CAPSULE_WORKER_URL="http://localhost:9999" go run main.go un-deploy \
-function=hello \
-revision=green

