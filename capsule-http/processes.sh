#!/bin/bash
rm processes.json
curl -X GET http://localhost:8080/functions/processes > processes.json

#http://capsule-ide.local:8080/functions/index-html