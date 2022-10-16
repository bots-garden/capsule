# Capsule FaaS (experimental)

There are four additional components to use **Capsule** (the wasm module launcher/executor) in **FaaS** mode:

- [capsule-registry](https://github.com/bots-garden/capsule-registry): a wasm module registry (🚧 support of https://wapm.io/ in progress)
- [capsule-reverse-proxy](https://github.com/bots-garden/capsule-reverse-proxy): a reverse-proxy to simplify the functions (wasm modules) access
- [capsule-worker](https://github.com/bots-garden/capsule-worker): a server to start the functions (wasm modules) remotely
- [capsule-ctl](https://github.com/bots-garden/capsule-ctl) (short name: `caps`): a CLI to facilitate the interaction with the worker

> - You can use the capsule registry independently of FaaS mode, only to provide wasm modules to the capsule launcher
> - You can use the capsule reverse-proxy independently of FaaS mode, only to get only one access URL
