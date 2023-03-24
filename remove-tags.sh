#!/bin/bash

TAG="v0.3.2"

git tag -d ${TAG}
git tag -d commons/${TAG}
git tag -d capsulemodule/${TAG}
