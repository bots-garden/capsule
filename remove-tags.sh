#!/bin/bash

TAG="v0.2.9"

git tag -d ${TAG}
git tag -d commons/${TAG}
git tag -d capsulemodule/${TAG}
