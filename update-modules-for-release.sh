#!/bin/bash

TAG="v0.2.6"
cd commons
go mod edit -replace github.com/bots-garden/capsule@${TAG}=../

cd ..
cd capsulemodule
go mod edit -replace github.com/bots-garden/capsule@${TAG}=../
go mod edit -replace github.com/bots-garden/capsule/commons@${TAG}=../commons

cd ..
git add .
git commit -m "📦 updates modules for ${TAG}"

git tag ${TAG}
git tag commons/${TAG}
git tag capsulemodule/${TAG}

git push origin main ${TAG} commons/${TAG} capsulemodule/${TAG}
