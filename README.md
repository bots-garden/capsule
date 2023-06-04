# Capsule, next generation

> the nano wasm functions runners

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/bots-garden/capsule)

## What is the **Capsule** project?

Capsule is a set of **WASM runners**. Right now, the Capsule project is composed of:

- Capsule **CLI**: to simply execute a **WebAssembly module** in a terminal
- Capsule **HTTP** server to serve a **WebAssembly module** like a micro service or a function.

> - **Capsule** applications are developed with GoLang and thanks to the 💜 **[Wazero](https://github.com/tetratelabs/wazero)** project. 
> - The wasm modules are developed in GoLang and compiled with **[TinyGo](https://tinygo.org/)** 💜 (with the WASI specification)

### Host DK & Module DK

- The **Capsule** applications are developed thanks to the [Capsule Host SDK (HDK)](https://bots-garden.github.io/capsule-host-sdk/)
- The **Capsule** modules executed by the The **Capsule** applications are developed thanks to the [Capsule Module SDK (MDK)](https://bots-garden.github.io/capsule-module-sdk/)

**🎉 That means, since now, it's possible to develop various runners thanks to the Capsule HostSDK**
