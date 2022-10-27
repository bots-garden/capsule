# Manual release process

- Change the version number (`vN.N.N`) in:
  - `./commons/version.go`
  - `/README.md`
  - `install-capsule-launcher.sh`
  - `sudo.install-capsule-launcher.sh`
  - All the documentation:
    - `docs/index.md`
    - `docs/install.md`
    - `docs/getting-started-cabu-inst.md`
    - `docs/getting-started-cabu-serve.md`
- 🖐 Check **every dependency** for every module
- Update and run `update-modules-for-release.sh`
- On GitHub: create a release + a tag (`vN.N.N`)
- Update https://github.com/bots-garden/capsule-docker-image
- Update https://github.com/bots-garden/capsule-function-builder
- Update the samples repositories:
  - https://github.com/bots-garden/capsule-samples
  - https://github.com/bots-garden/capsule-faas-demo
  - https://github.com/bots-garden/capsule-on-fly-dot-io
  - https://github.com/bots-garden/capsule-launcher-demo

> don't forget to update the documentation
