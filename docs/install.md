# Install Capsule

Before executing or running a function, you need to install the last release of **Capsule**:
> The script will install **Capsule** in `$HOME/.local/bin`
> if you want to install Capsule somewhere else, override the `CAPSULE_PATH` variable (default value: `CAPSULE_PATH="$HOME/.local/bin"`)

🖐 **On Linux**:

```bash
export CAPSULE_VERSION="v0.3.2"
wget -O - https://raw.githubusercontent.com/bots-garden/capsule/${CAPSULE_VERSION}/install-capsule-launcher.sh| bash
```

🖐 **On macOS**:

- create the `$HOME/.local/bin` directory
- add it to your path:
```bash
export CAPSULE_RUNNER_PATH="$HOME/.local"
export PATH="$CAPSULE_RUNNER_PATH/bin:$PATH"
```

Then you can serve a wasm function like this:

```bash
MESSAGE="👋 Hello World 🌍" capsule \
  -wasm=./app/index.wasm \
  -mode=http \
  -httpPort=8080
```

> You can download the appropriate release of **Capsule** here: [`v0.3.2 🤗 [WASM I/O 2023]`](https://github.com/bots-garden/capsule/releases/tag/v0.3.2)

## Using the Capsule Docker image
> The documentation is a wip 🚧

A "scratch" Docker image of Capsule exists on [https://hub.docker.com/r/k33g/capsule-launcher/tags](https://hub.docker.com/r/k33g/capsule-launcher/tags). You can find more details on the [capsule-docker-image](https://github.com/bots-garden/capsule-docker-image) project.

This image will be used to deploy Capsule to CaaS or Kubernetes. You can use it directly to run a wasm function without installing Capsule:

```bash
docker run \
  -p 8080:8080 \
  -e MESSAGE="👋 Hello World 🌍" \
  -v $(pwd):/app --rm k33g/capsule-launcher:0.3.1 \
  /capsule \
  -wasm=./app/index.wasm \
  -mode=http \
  -httpPort=8080
```

👋 You will find some **running examples** with these projects:

- [https://github.com/bots-garden/capsule-launcher-demo](https://github.com/bots-garden/capsule-launcher-demo)
- [https://github.com/bots-garden/capsule-hello-universe](https://github.com/bots-garden/capsule-hello-universe)

> Old samples to be updated:

> - [https://github.com/bots-garden/capsule-samples](https://github.com/bots-garden/capsule-samples)
> - [https://github.com/bots-garden/capsule-on-fly-dot-io](https://github.com/bots-garden/capsule-on-fly-dot-io)
