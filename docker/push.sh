#!/bin/bash

set -o allexport; source .env; set +o allexport
echo "🐋 ${IMAGE_NAME}:${IMAGE_TAG}"
 
docker tag ${IMAGE_NAME} ${DOCKER_USER}/${IMAGE_NAME}:${IMAGE_TAG}
docker push ${DOCKER_USER}/${IMAGE_NAME}:${IMAGE_TAG}
