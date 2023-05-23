#!/bin/bash

#curl -X PUT http://localhost:8080/functions/switch/hello-world/default/saved


#http://capsule-ide.local:8080/functions/index-html


JSON_DATA='{"name":"Bob Morane 🤣","age":42}'

curl -X POST http://localhost:8080/functions/hello-world/saved \
          -H 'Content-Type: application/json; charset=utf-8' \
          -d "${JSON_DATA}"
