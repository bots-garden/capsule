# Capsule: the nano (wasm) functions runner

!!! info "What's new?"
    - `v0.3.0 🎄 [Christmas tree]`: Refactoring (Capsule is at least 4 times faster than the previous version).
    - `v0.2.9 🦜 [parrot]`: Hot reloading of the wasm module [see the "Reload the module" section](getting-started-cabu-reload.md) and HTTP service refactoring.
    - `v0.2.8 🦤 [dodo]`: Capsule uses now [Fiber](https://github.com/gofiber/fiber) instead [Gin](https://github.com/gin-gonic/gin). The size of the Capsule Runner Docker image is now 16.8M!

## What is **Capsule**?

**Capsule** is a **WebAssembly Function Runner**. It means that **Capsule** is both:

- An **HTTP server** that serves **WebAssembly functions**
- A **NATS** subscriber and publisher (written with WebAssembly)
- A **MQTT** subscriber and publisher (written with WebAssembly)
- A **CLI**, you can simply execute a WASM function in a terminal

> - **Capsule** is developed with GoLang and thanks to the 💜 **[Wazero](https://github.com/tetratelabs/wazero)** project
> - The wasm modules are developed in GoLang and compiled with **[TinyGo](https://tinygo.org/)** 💜 (with the WASI specification)

## What does a **WASM function** look like with Capsule?

```golang
package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

func main() {

	hf.SetHandleHttp(Handle)
}

func Handle(request hf.Request) (response hf.Response, errResp error) {

	headersResp := map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}

	jsondoc := `{"message": "👋 Hello World 🌍"}`

	return hf.Response{Body: jsondoc, Headers: headersResp}, err
}
```

## What are the **added values** of Capsule?

Capsule brings superpowers to the WASM function modules with **host functions**. Thanks to these **host functions**, a **WASM function** can, for example, prints a message, reads files, writes to files, makes HTTP requests, ... See the [host functions section](host-functions-intro.md).


!!! info "Useful information for this project"
    - 🖐 Issues: [https://github.com/bots-garden/capsule/issues](https://github.com/bots-garden/capsule/issues)
    - 🚧 Milestones: [https://github.com/bots-garden/capsule/milestones](https://github.com/bots-garden/capsule/milestones)
    - 📦 Last release: `v0.3.0 🎄 [Christmas tree]`
    - 📦 Next release: `v0.3.1 🎅 [santa]` *🚧 in progress*
    - 📦 Releases: [https://github.com/bots-garden/capsule/releases](https://github.com/bots-garden/capsule/releases)
