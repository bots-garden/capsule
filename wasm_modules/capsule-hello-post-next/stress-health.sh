#!/bin/bash
# -n 10000 -c 1000
#hey -n 10 -c 5 -m POST \
hey -n 1000 -c 50 -m GET \
-H "Content-Type: application/json" \
"http://localhost:7070/health"

