#!/bin/bash
LAST_CAPSULE_VERSION="v0.3.3"
echo "System: ${OSTYPE} $(uname -m)"

if [ -z "$CAPSULE_PATH" ]
then
    CAPSULE_PATH="$HOME/.local/bin"
fi

if [[ $1 = "help" ]]
then
    echo "usage: $0"
    echo "The script will detect the OS & ARCH and use the last version of capsule (${LAST_CAPSULE_VERSION})"
    echo "You can force the values by setting these environment variables:"
    echo "- CAPSULE_OS (linux, darwin)"
    echo "- CAPSULE_ARCH (amd64, arm64)"
    echo "- CAPSULE_VERSION (default: ${LAST_CAPSULE_VERSION})"
    echo "- CAPSULE_PATH (default: ${CAPSULE_PATH})"
    exit 0
fi

if [ -z "$CAPSULE_VERSION" ]
then
    CAPSULE_VERSION=$LAST_CAPSULE_VERSION
fi

if [ -z "$CAPSULE_OS" ]
then
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
      CAPSULE_OS="linux"
    elif [[ "$OSTYPE" == "darwin"* ]]; then
      CAPSULE_OS="darwin"
    else
      CAPSULE_OS="linux"
    fi
fi

if [ -z "$CAPSULE_ARCH" ]
then
    if [[ "$(uname -m)" == "x86_64" ]]; then
      CAPSULE_ARCH="amd64"
    elif [[ "$(uname -m)" == "arm64" ]]; then
      CAPSULE_ARCH="arm64"
    elif [[ $(uname -m) == "aarch64" ]]; then
      CAPSULE_ARCH="arm64"
    else
      CAPSULE_ARCH="amd64"
    fi
fi

CAPSULE_MODULE="capsule"

echo "Installing ${CAPSULE_MODULE}..."
wget https://github.com/bots-garden/capsule/releases/download/${CAPSULE_VERSION}/${CAPSULE_MODULE}-${CAPSULE_VERSION}-${CAPSULE_OS}-${CAPSULE_ARCH}.tar.gz
tar -zxf ${CAPSULE_MODULE}-${CAPSULE_VERSION}-${CAPSULE_OS}-${CAPSULE_ARCH}.tar.gz --directory ${CAPSULE_PATH}
rm ${CAPSULE_MODULE}-${CAPSULE_VERSION}-${CAPSULE_OS}-${CAPSULE_ARCH}.tar.gz

echo "Capsule $(capsule version) is installed 🎉"
