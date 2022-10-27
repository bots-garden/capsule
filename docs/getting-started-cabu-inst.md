# 🚀 Getting Started

> The simplest way to create your first **Capsule Function** is to use **Cabu**

To write and build wasm function for Capsule, you need to install GoLang and TinyGo. Otherwise, you can use the [capsule-function-builder](https://github.com/bots-garden/capsule-function-builder) project. It provides a very simple CLI, named **capsule-builder** or **cabu** that uses a Docker image with all the necessary resources (Golang and TinyGo compilers).

## Install **Cabu** (Capsule Builder)

```bash
CAPSULE_BUILDER_VERSION="v0.0.3"
wget -O - https://raw.githubusercontent.com/bots-garden/capsule-function-builder/${CAPSULE_BUILDER_VERSION}/install-capsule-builder.sh | bash
```

