# Capsule: the nano (wasm) functions runner

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

### Capsule brings superpowers to the WASM functions

Thanks to **host functions** provided by **Capsule**, the **WASM functions** can:

| **Description**  | **Host function**  |
|---|---|
| Print a message to the console | `hf.Log(string)` |
| Read files | `hf.ReadFile("about.txt")` |
| Write files | `hf.WriteFile("hello.txt", "👋 HELLO WORLD 🌍")`|
| **Read value of the environment variables** | `hf.GetEnv("MESSAGE")` |
| Make HTTP requests | `hf.Http("https://httpbin.org/post", "POST", headers, "👋 hello world 🌍")` |
| Use memory cache (set) | `hf.MemorySet("message", "🚀 hello is started")` |
|  | `hf.MemoryGet("message")` |
| Make Redis queries | `hf.RedisSet("greetings", "Hello World")` |
|  | `hf.RedisGet("greetings")` |
|  | `hf.RedisKeys("bob*")` |
| **Make CouchBase N1QL Query** | `jsonStringArray, err := hf.CouchBaseQuery(query)` |
| **Use Nats** | `hf.NatsPublish("subject", "hello")` |
| | `hf.NatsReply("it's a wasm module here", 10)` |
| | `hf.NatsGetSubject()` |
| | `hf.NatsGetServer()` |
| | `hf.NatsConnectPublish("nats.devsecops.fun:4222", "subject", "🖐 Hello from WASM with Nats 💜")` |
| | `hf.NatsConnectRequest("nats.devsecops.fun:4222", "subject", "👋 Hello World 🌍", 1)` |
| **Use MQTT** | `hf.MqttConnectPublish("127.0.0.1:1883", "sensor_id0", "topic", "👋 Hello World 🌍")` |
| | `hf.MqttGetTopic()` |
| | `hf.MqttPublish("topic", "it's a wasm module here")` |
| Manage Errors | *🖐 🚧 it's a work in progress* |
| | `hf.GetExitError()` |
| | `hf.GetExitCode()` |

## Information

| **Label**  | **Description**  |
|---|---|
| Issues        | [https://github.com/bots-garden/capsule/issues](https://github.com/bots-garden/capsule/issues)  |
| Last release  | `v0.2.8 🦤 [dodo]`  |
| Dev release   | `v0.2.9 🦜 [parrot]` *🚧 in progress*  |
| Releases      | [https://github.com/bots-garden/capsule/releases](https://github.com/bots-garden/capsule/releases) |

## What's new

`v0.2.8`: Capsule uses now [Fiber](https://github.com/gofiber/fiber) instead [Gin](https://github.com/gin-gonic/gin). The size of the Capsule Runner Docker image is now 16.8M!
