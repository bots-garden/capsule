# Tooling
> The documentation is a wip 🚧

To write and build wasm function for Capsule, you need to install GoLang and TinyGo. Otherwise, you can use the [capsule-function-builder](https://github.com/bots-garden/capsule-function-builder) project. It provides a very simple CLI, named **capsule-builder** or **cabu** that uses a Docker image with all the necessary resources (Golang and TinyGo compilers).

## Install Capsule Builder

```bash
CAPSULE_BUILDER_VERSION="v0.0.2"
wget -O - https://raw.githubusercontent.com/bots-garden/capsule-function-builder/${CAPSULE_BUILDER_VERSION}/install-capsule-builder.sh | bash
```

Then you can generate a new project from a template:

```bash
# template name: `service-get`
# function project name `hello-world`
cabu generate service-get hello-world
```


Then, build it easily:

```bash
cd hello-world
cabu build . hello-world.go hello-world.wasm
```

And, finally, serve it:

```bash
capsule \
  -wasm=./hello-world.wasm \
  -mode=http \
  -httpPort=8080
```
