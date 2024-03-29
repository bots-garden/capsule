# Release process

> Right now, it's a manual release process.

## Prerequisites

👋 **be on the main branch**

Last release: `v0.4.2 ⛱️ [beach umbrella]`
Last release: `v0.4.1 🫑 [pepper]`
Last release: `v0.4.0 🌶️ [chili pepper]`
Last release: `v0.3.9 🥒 [cucumber]`
Last release: `v0.3.8 🥬 [leafy greens]` 
Last release: `v0.3.7 🥦 [broccoli]`

### Update documentation content with the new tag

Check every documents:
- `docs/*.md`

### Update version number in Go source

- `capsule-http/tools/description.txt`
- `capsule-cli/description.txt`
- `capsctl/description.txt`

### Update version tag in Dockerfile

- `capsule-http/functions/hello-world/Dockerfile`

### Update version tag in Taskfile

- `capsule-http/Taskfile.yml` (`IMAGE_TAG: "N.N.N"`)

### Update the main Taskfile

Go to `./Taskfile.yml`

- Update `TAG` of the `release` task (`TAG: "vN.N.N"`)
- Update `TAG` of the `remove-tag` task
- Update `TAG` of the `build-releases` task
- Update `IMAGE_TAG` of the `build-push-docker-image` task (`IMAGE_TAG: "N.N.N"`)

## Publish the release

Run:

```bash
task release
task build-releases
#task build-docker-capsule-http-image-darwin-arm64
#task build-docker-capsule-http-image-darwin-amd64
task build-docker-capsule-http-image-linux-amd64
task build-docker-capsule-http-image-linux-arm64
```

- Create the release on GitHub
- Upload the files

Publish the website:

```bash
task publish-mkdocs
```

🎉