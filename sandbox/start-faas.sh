#!/bin/bash
#!/bin/bash
DOWNLOADED_FILES_PATH="./registry/functions"
CAPSULE_REGISTRY_ADMIN_TOKEN="AZERTYUIOP" \
capsule-registry \
  -files="${DOWNLOADED_FILES_PATH}" \
  -httpPort=4999 &

CAPSULE_REVERSE_PROXY_ADMIN_TOKEN="1234567890" \
capsule-reverse-proxy \
  -backend="memory" \
  -httpPort=8888 &

CAPSULE_REVERSE_PROXY_ADMIN_TOKEN="1234567890" \
CAPSULE_WORKER_ADMIN_TOKEN="0987654321" \
capsule-worker \
  -reverseProxy=http://localhost:8888 \
  -backend=memory \
  -capsulePath=capsule \
  -httpPortCounter=10000 \
  -httpPort=9999 &

# get list of all modules of the registry
# curl https://capsdev.devsecops.fun:4999/modules

# get information on a module
# curl https://capsdev.devsecops.fun:4999/info/demo/hello/0.0.0

# -reverseProxy=http://localhost:8888 \
# -reverseProxy=https://capsdev.devsecops.fun:8888 \

