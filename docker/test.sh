#!/bin/bash
set -o allexport; source .env; set +o allexport
docker login -u ${DOCKER_USER} -p ${DOCKER_PWD}

#docker run -it ${IMAGE_NAME} bash
docker run -it ${IMAGE_NAME} sh

