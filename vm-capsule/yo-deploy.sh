#!/bin/bash
#!/bin/bash
CAPSULE_WORKER_URL="https://capsdev.devsecops.fun:9999" \
CAPSULE_BACKEND="memory" \
CAPSULE_WORKER_ADMIN_TOKEN="0987654321" \
caps deploy \
-function=yo \
-revision=orange \
-downloadUrl=https://capsdev.devsecops.fun:4999/k33g/yo/0.0.0/yo.wasm \
-envVariables='{"MESSAGE": "Revision 🟠>😡"}'


CAPSULE_WORKER_URL="https://capsdev.devsecops.fun:9999" \
CAPSULE_BACKEND="memory" \
CAPSULE_WORKER_ADMIN_TOKEN="0987654321" \
caps deploy \
-function=yo \
-revision=orange \
-downloadUrl=https://capsdev.devsecops.fun:4999/k33g/yo/0.0.0/yo.wasm \
-envVariables='{"MESSAGE": "Revision 🟠>😡🤬"}'

CAPSULE_WORKER_URL="https://capsdev.devsecops.fun:9999" \
CAPSULE_BACKEND="memory" \
CAPSULE_WORKER_ADMIN_TOKEN="0987654321" \
caps deploy \
-function=yo \
-revision=orange \
-downloadUrl=https://capsdev.devsecops.fun:4999/k33g/yo/0.0.0/yo.wasm \
-envVariables='{"MESSAGE": "Revision 🟠>😡🤬🥵"}'

# https://capsdev.devsecops.fun:8888/functions/yo/orange
