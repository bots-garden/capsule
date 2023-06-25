# Capsule Project: the nano wasm runners

!!! info "What's new?"
	- `v0.4.0 🌶️ [chili pepper]`: update of [HDK 0.0.4 then 0.0.5](https://github.com/bots-garden/capsule-host-sdk/releases/tag/v0.0.5), (🎉 performances: more than x 2 🚀). **capsule-http**: add of 2 endpoints (`/metrics`and `/health`) triggering the `OnMetrics` and `OnHealthCheck` functions of the WASM module. 
    - `v0.3.9 🥒 [cucumber]`: update of [HDK 0.0.3](https://github.com/bots-garden/capsule-host-sdk) with [Wazero 1.2.0](https://github.com/tetratelabs/wazero/releases/tag/v1.2.0) and [MDK 0.0.3](https://github.com/bots-garden/capsule-module-sdk) (encoding of the HTML string into JSON string, then it's easier to serve HTML)
    - `v0.3.8 🥬 [leafy greens]`: 🐛 fixes of the **FaaS** mode
    - `v0.3.7 🥦 [broccoli]`: 🚀 **FaaS** mode (documentation in progress) + **NGrok** integration
    - `v0.3.6 🫐 [blueberries]`: Prometheus metrics + 🐳 Docker images
    - `v0.3.5 🍓 [strawberry]`: Update with HDK & MDK `v0.0.2`
    - `v0.3.4 🍋 [lemon]`: Capsule next generation (performances: x 10 🚀)
    - 🌍 Downloads: [https://github.com/bots-garden/capsule/releases/tag/v0.3.8](https://github.com/bots-garden/capsule/releases/tag/v0.3.5)
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

#### Tutorials
> More tutorials are coming soon!

- [Capsule: the WASM runners project](https://k33g.hashnode.dev/capsule-the-wasm-runners-project): with this blog post I explain how to create WASM modules (with the **MDK**) for the Capsule CLI and the Capsule HTTP server, but too, how to create your Capsule application (with the **HDK**).

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
