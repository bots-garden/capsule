# Manual release process

- Change the version number (`N.N.N`) in:
  - `./commons/version.go`
  - `/README.md`
  - `install-capsule-ctl.sh`
  - `install-capsule-launcher.sh`
  - `install-capsule-registry.sh`
  - `install-capsule-reverse-proxy.sh`
  - `install-capsule-worker.sh`
- Check **every dependency** for every module
- Commit & Push
- On GitHub: create a release + a tag (`N.N.N`)
- Update the samples repositories:
  - https://github.com/bots-garden/capsule-samples
  - https://github.com/bots-garden/capsule-faas-demo
  - https://github.com/bots-garden/capsule-on-fly-dot-io
  - https://github.com/bots-garden/capsule-launcher-demo
