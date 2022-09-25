#!/bin/bash

TAG="v0.2.5"

git tag ${TAG}
git tag commons/${TAG}
git tag capsulemodule/${TAG}

git add .
git commit -m "📦 updates modules for ${TAG}"

git push origin main ${TAG} commons/${TAG} capsulemodule/${TAG}

