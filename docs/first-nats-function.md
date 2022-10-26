# Developer Guide

## First Nats function
> 🖐🚧 The NAT integration with **Capsule** is a work in progress and the functions are subject to change

NATS is an open-source messaging system.

> - About NATS: https://nats.io/ and https://docs.nats.io/
> - Nats Overview: https://docs.nats.io/nats-concepts/overview

### Requirements

#### NATS Server

You need to install and run a NATS server: https://docs.nats.io/running-a-nats-service/introduction/installation.
Otherwise, I created a Virtual Machine for this; If you have installed [Multipass](https://multipass.run/), go to the `./nats/vm-nats` directory of this project. I created some scripts for my experiments:

- `create-vm.sh` *create the multipass VM, the settings of the VM are stored in the `vm.nats.config`*
- `01-install-nats-server.sh` *install the NATS server inside the VM*
- `02-start-nats-server.sh` *start the NATS server*
- `03-stop-nats-server.sh` *stop the NATS server*
- `stop-vm.sh` *stop the VM*
- `start-vm.sh` *start the VM*
- `destroy-vm.sh` *delete the VM*
- `shell-vm.sh` *SSH connect to the VM*

#### NATS Client

You need a NATS client to publish messages. You can find sample of Go and Node.js NATS clients in the `./nats/clients`.

### Run **Capsule** as a NATS subscriber:

```bash
capsule \
   -wasm=../wasm_modules/capsule-nats-subscriber/hello.wasm \
   -mode=nats \
   -natssrv=nats.devsecops.fun:4222 \
   -subject=ping
```
> - use the "NATS mode": `-mode=nats`
> - define the NATS subject: `-subject=<subject_name>`
> - define the address of the NATS server: `-natssrv=<nats_server:port>`

### NATS function

A **Capsule** NATS function is a subscription to a subject. **Capsule** is listening on a subject(like a MQTT topic) and execute a function every time a message is posted on the subject:

```golang
package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

func main() {
	hf.OnNatsMessage(Handle) // define the triggered function when a message "arrives" on the subject/topic
}

// at every message on the subject channel, the `Handle` function is executed
func Handle(params []string) {
	// send a message to another subject
	_, err := hf.NatsPublish("notify", "it's a wasm module here")

	if err != nil {
		hf.Log("😡 ouch something bad is happening")
		hf.Log(err.Error())
	}
}
```


### Capsule NATS publisher
> Publish NATS messages from capsule

You can use a **WASM Capsule module** to publish NATS messages, even if **Capsule** is not started in "nats" mode, for example from a **WASM CLI Capsule module**:

```golang
package main

import (
    "errors"
    hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
    "strings"
)

func main() {
    hf.SetHandle(Handle)
}

func Handle(params []string) (string, error) {
    var errs []string

    // a new connection is created at every call/publish
    _, err1stMsg := hf.NatsConnectPublish("nats.devsecops.fun:4222", "ping", "🖐 Hello from WASM with Nats 💜")
    _, err2ndMsg := hf.NatsConnectPublish("nats.devsecops.fun:4222", "notify", "👋 Hello World 🌍")

    if err1stMsg != nil {
        errs = append(errs, err1stMsg.Error())
    }
    if err2ndMsg != nil {
        errs = append(errs, err2ndMsg.Error())
    }

    return "NATS Rocks!", errors.New(strings.Join(errs, "|"))
}
```
> In this use case, you need to define the NATS server and create a connection

### Request and Reply

A NATS "publisher" can make a request to a NATS "subscriber" and wait for an answer

```golang
package main

import (
	"errors"
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
	"strings"
)

func main() {
	hf.SetHandle(Handle)
}

func Handle(params []string) (string, error) {

	// Publish and wait for an answer; 1 is the timeout in seconds
	res, err := hf.NatsConnectRequest("nats.devsecops.fun:4222", "notify", "👋 Hello World 🌍", 1)

	if err != nil {
		hf.Log("🔴" + err.Error())
	} else {
        // Display the answer
		hf.Log("🔵" + res)
	}

	return "NATS Rocks!", err
}
```

A NATS "subscriber" can reply to a request received on its subject

```golang
package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

func main() {
	hf.OnNatsMessage(Handle)
}

func Handle(params []string) {

	hf.Log("Message on subject: " + hf.NatsGetSubject() + ", 🎉 message: " + params[0])

	// reply to the message on the current subject; 10 is the timeout in seconds
	_, _ = hf.NatsReply("Hey! What's up", 10)

}
```
