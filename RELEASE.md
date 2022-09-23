# Manual release process

- Change the version number (`vN.N.N`) in:
  - `./commons/version.go`
  - `/README.md`
  - `install-capsule-ctl.sh`
  - `install-capsule-launcher.sh`
  - `install-capsule-registry.sh`
  - `install-capsule-reverse-proxy.sh`
  - `install-capsule-worker.sh`
- Check **every dependency** for every module

```bash
TAG="vN.N.N"
cd commons
go mod edit -replace github.com/bots-garden/capsule@${TAG}=../

cd ..
cd capsulemodule
go mod edit -replace github.com/bots-garden/capsule@${TAG}=../
go mod edit -replace github.com/bots-garden/capsule/commons@${TAG}=../commons

cd ..
git add .
git commit -m "updates modules for ${TAG}"

git tag ${TAG}
git tag commons/${TAG}
git tag capsulemodule/${TAG}


git push origin main ${TAG} commons/${TAG} capsulemodule/${TAG}
```


- Commit & Push
- On GitHub: create a release + a tag (`vN.N.N`)
- Update the samples repositories:
  - https://github.com/bots-garden/capsule-samples
  - https://github.com/bots-garden/capsule-faas-demo
  - https://github.com/bots-garden/capsule-on-fly-dot-io
  - https://github.com/bots-garden/capsule-launcher-demo
