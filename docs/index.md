# Capsule Project: the nano wasm runners

!!! info "What's new?"
    - `v0.3.7 🥦 [broccoli]`: 🚀 **FaaS** mode (documentation in progress) + **NGrok** integration
    - `v0.3.6 🫐 [blueberries]`: Prometheus metrics + 🐳 Docker images
    - `v0.3.5 🍓 [strawberry]`: Update with HDK & MDK `v0.0.2`
    - `v0.3.4 🍋 [lemon]`: Capsule next generation (performances: x 10 🚀)
    - 🌍 Downloads: [https://github.com/bots-garden/capsule/releases/tag/v0.3.5](https://github.com/bots-garden/capsule/releases/tag/v0.3.5)
    - 🤚 With the previous version of the project, Capsule was an only one application to run as an HTTP server, a CLI, a NATS subscriber and publisher and a MQTT subscriber and publisher. In the future, we will reintroduce the capabilities of NATS and MQTT, but with separate runners.

## What is the **Capsule** project?

Capsule is a set of **WASM runners**. Right now, the Capsule project is composed of:

- Capsule **CLI**: to simply execute a **WebAssembly module** in a terminal
- Capsule **HTTP** server to serve a **WebAssembly module** like a micro service or a function.

> - **Capsule** applications are developed with GoLang and thanks to the 💜 **[Wazero](https://github.com/tetratelabs/wazero)** project. 
> - The wasm modules are developed in GoLang and compiled with **[TinyGo](https://tinygo.org/)** 💜 (with the WASI specification)

### Host DK & Module DK

- The **Capsule** applications are developed thanks to the [Capsule Host SDK (HDK)](https://bots-garden.github.io/capsule-host-sdk/)
- The **Capsule** modules executed by the The **Capsule** applications are developed thanks to the [Capsule Module SDK (MDK)](https://bots-garden.github.io/capsule-module-sdk/)

**🎉 That means, since now, it's possible to develop various runners thanks to the Capsule Host SDK**

> Tutorials are coming soon!

## What does a **WASM Capsule module** look like?

### WASM Module for the Capsule CLI
```golang
package main

import (
	capsule "github.com/bots-garden/capsule-module-sdk"
)

func main() {
	capsule.SetHandle(Handle)
}

// Handle function
func Handle(params []byte) ([]byte, error) {

	capsule.Print("Environment variable → MESSAGE: " + capsule.GetEnv("MESSAGE"))

	err := capsule.WriteFile("./hello.txt", []byte("👋 Hello World! 🌍"))
	if err != nil {
		capsule.Print(err.Error())
	}

	data, err := capsule.ReadFile("./hello.txt")
	if err != nil {
		capsule.Print(err.Error())
	}
	capsule.Print("📝: " + string(data))
	
	return []byte("👋 Hello " + string(params)), nil
}

```

### WASM Module for the Capsule HTTP server
```golang
// Package main
package main

import (
	"strconv"
	"github.com/bots-garden/capsule-module-sdk"
	"github.com/valyala/fastjson"
)

func main() {
	capsule.SetHandleHTTP(Handle)
}

// Handle function 
func Handle(param capsule.HTTPRequest) (capsule.HTTPResponse, error) {
	
	capsule.Print("📝: " + param.Body)
	capsule.Print("🔠: " + param.Method)
	capsule.Print("🌍: " + param.URI)
	capsule.Print("👒: " + param.Headers)
	
	var p fastjson.Parser
	v, err := p.Parse(param.Body)
	if err != nil {
		capsule.Log(err.Error())
	}
	message := string(v.GetStringBytes("name")) + " " + strconv.Itoa(v.GetInt("age"))
	capsule.Log(message)

	response := capsule.HTTPResponse{
		JSONBody: `{"message": "`+message+`", "things":{"emoji":"🐯"}}`,
		Headers: `{"Content-Type": "application/json; charset=utf-8"}`,
		StatusCode: 200,
	}

	return response, nil
}
```

## What are the **added values** of Capsule?

Capsule applications bring superpowers to the WASM modules with **host functions**. Thanks to these **host functions**, a **WASM function** can, for example, prints a message, reads files, writes to files, makes HTTP requests, ... See the [host functions section](host-functions-intro.md).


!!! info "Useful information for this project"
    - 🖐 Issues: [https://github.com/bots-garden/capsule/issues](https://github.com/bots-garden/capsule/issues)
    - 🚧 Milestones: [https://github.com/bots-garden/capsule/milestones](https://github.com/bots-garden/capsule/milestones)
    - 📦 Releases: [https://github.com/bots-garden/capsule/releases](https://github.com/bots-garden/capsule/releases)
