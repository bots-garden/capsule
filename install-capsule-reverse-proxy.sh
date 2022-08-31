#!/bin/bash
CAPSULE_VERSION="0.1.7"
CAPSULE_MODULE="capsule-reverse-proxy"

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  CAPSULE_OS="linux"
elif [[ "$OSTYPE" == "darwin"* ]]; then
  CAPSULE_OS="darwin"
else
  CAPSULE_OS="linux"
fi

if [[ "$(uname -m)" == "x86_64" ]]; then
  CAPSULE_ARCH="amd64"
elif [[ "$OSTYPE" == "arm64" ]]; then
  CAPSULE_ARCH="arm64"
else
  CAPSULE_ARCH="amd64"
fi

CAPSULE_ARCH=$(uname -m)

echo "💊 Installing ${CAPSULE_MODULE}..."
wget https://github.com/bots-garden/capsule/releases/download/${CAPSULE_VERSION}/${CAPSULE_MODULE}-${CAPSULE_VERSION}-${CAPSULE_OS}-${CAPSULE_ARCH}.tar.gz
sudo tar -zxf ${CAPSULE_MODULE}-${CAPSULE_VERSION}-${CAPSULE_OS}-${CAPSULE_ARCH}.tar.gz --directory /usr/local/bin
rm ${CAPSULE_MODULE}-${CAPSULE_VERSION}-${CAPSULE_OS}-${CAPSULE_ARCH}.tar.gz
