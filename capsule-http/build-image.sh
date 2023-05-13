#!/bin/bash
IMAGE_NAME="demo-capsule-http"
docker login -u ${DOCKER_USER} -p ${DOCKER_PWD}
docker build -t ${IMAGE_NAME} . 

docker images | grep ${IMAGE_NAME}
