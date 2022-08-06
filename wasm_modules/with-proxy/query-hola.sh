#!/bin/bash
curl -v \
  http://localhost:8888/functions/hola \
  -H 'content-type: application/json; charset=utf-8'
