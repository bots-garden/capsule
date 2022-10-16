# Capsule: the nano (wasm) functions runner

> 🚧 This is a wip

## What is **Capsule**?

**Capsule** is a WebAssembly function launcher(runner). It means that, with **Capsule** you can:

- From your terminal, execute a function of a wasm module (the **"CLI mode"**)
- Serving a function of a wasm module through http (the **"HTTP mode"**)
- Serving a function of a wasm module through NATS (the **"NATS mode"**), in this case **Capsule** is used as a NATS subscriber and can reply on a subject(topic)
- Serving a function of a wasm module through MQTT (the **"MQTT mode"**), in this case **Capsule** is used as a MQTT subscriber and can reply on a subject(topic)

> - **Capsule** is developed with GoLang and thanks to the 💜 **[Wazero](https://github.com/tetratelabs/wazero)** project
> - The wasm modules are developed in GoLang and compiled with TinyGo (with the WASI specification)



## What's new

- `v0.2.8`: Capsule uses now [Fiber](https://github.com/gofiber/fiber) instead [Gin](https://github.com/gin-gonic/gin). The size of the Capsule Runner Docker image is now 16.8M!
- `v0.2.7`:
    - The FaaS components are externalized, now, this project is **only** for the **Capsule Runner**
    - "Scratch" Docker image (18.5M) to easily use and deploy the Capsule Runner (https://github.com/bots-garden/capsule-docker-image)
    - **cabu** (or **capsule-builder**) (https://github.com/bots-garden/capsule-function-builder): a CLI using a specific Docker image allowing:
        - the creation of a wasm function project (from templates)
        - the build of the wasm function, without installing anything (TinyGo is embedded in the image) (https://github.com/bots-garden/capsule-function-builder)
- `v0.2.6`: Wazero: updates to `1.0.0-pre.2` by [@codefromthecrypt](https://github.com/codefromthecrypt) + a logo
- `v0.2.5`: Add MQTT support by [@py4mac](https://github.com/py4mac) with `MqttPublish` & `MqttPublish`
- `v0.2.4`: Add 2 wasm helper functions `flatjson.StrToMap` and `flatjson.MapToStr` (update 2022/10/10: these two helpers has been removed)
- `v0.2.3`: NATS support, 2 new functions: `NatsReply` and `NatsConnectRequest`
- `v0.2.2`: like `0.2.1` with fixed modules dependencies, and tag name start with a `v`
- `0.2.1`: NATS support (1st stage) `OnNatsMessage`, `NatsPublish`, `NatsConnectPublish`, `NatsConnectPublish`, `NatsGetSubject`, `NatsGetServer`
- `0.2.0`: `OnLoad` & `OnExit` functions + Memory cache host functions (`MemorySet`, `MemoryGet`, `MemoryKeys`)
- `0.1.9`: Add `Request` and `Response` types (for the Handle function)
- `0.1.8`: Redis host functions: add the KEYS command (`RedisKeys(pattern string)`)
