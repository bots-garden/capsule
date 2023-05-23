#!/bin/bash

curl -X DELETE http://localhost:8080/functions/stop/hello-world
curl -X DELETE http://localhost:8080/functions/stop/index-html

#http://capsule-ide.local:8080/functions/index-html