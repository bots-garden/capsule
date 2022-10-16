<!--
<img src="./logos/capsule-wasm-logo.png" alt="capsule-wasm-logo.png" width="40%" height="40%"/>
<img src="./logos/capsule-logo.png" alt="capsule-logo.png" width="30%" height="30%"/>
-->
<img src="./logos/capsule-logo-readme.png" alt="capsule-logo.png"  width="10%" height="10%"/>

# Capsule: the nano (wasm) functions runner

- Issues: [https://github.com/bots-garden/capsule/issues](https://github.com/bots-garden/capsule/issues)
- Last release: `v0.2.8 🦤 [dodo]`
- Dev release: `v0.2.9 🦜 [parrot]` *🚧 in progress*

## What is **Capsule**?

**Capsule** is a WebAssembly function launcher(runner). It means that, with **Capsule** you can:

- From your terminal, execute a function of a wasm module (the **"CLI mode"**)
- Serving a function of a wasm module through http (the **"HTTP mode"**)
- Serving a function of a wasm module through NATS (the **"NATS mode"**), in this case **Capsule** is used as a NATS subscriber and can reply on a subject(topic)
- Serving a function of a wasm module through MQTT (the **"MQTT mode"**), in this case **Capsule** is used as a MQTT subscriber and can reply on a subject(topic)

> - **Capsule** is developed with GoLang and thanks to the 💜 **[Wazero](https://github.com/tetratelabs/wazero)** project
> - The wasm modules are developed in GoLang and compiled with TinyGo (with the WASI specification)

