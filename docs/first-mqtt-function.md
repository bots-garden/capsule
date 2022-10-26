# Developer Guide

## First MQTT function
> 🖐🚧 The MQTT integration with **Capsule** is a work in progress and the functions are subject to change

MQTT is a standard for IOT message.

> - About MQTT: https://mqtt.org/

### Requirements

#### MQTT Server

You need to install and run a MQTT server. To do so, go to the `./mqtt` directory of this project and run the docker-compose file

### Run **Capsule** as a MQTT subscriber:

```bash
capsule \
   -wasm=../wasm_modules/capsule-mqtt-subscriber/hello.wasm \
   -mode=mqtt \
   -mqttsrv=127.0.0.1:1883 \
   -topic=topic/sensor0 \
   -clientId=sensor
```
> - use the "MQTT mode": `-mode=mqtt`
> - define the MQTT topic: `-topic=<topic_name>`
> - define the MQTT clientId: `-clientId=<clientId>`
> - define the address of the MQTT server: `-mqttsrv=<mqtt_server:port>`

### MQTT function

A **Capsule** MQTT function is a subscription to a subject. **Capsule** is listening on a topic and execute a function every time a message is posted on the subject:

```golang
package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

func main() {
	hf.OnMqttMessage(Handle) // define the triggered function when a message "arrives" on the topic
}

// at every message on the subject channel, the `Handle` function is executed
func Handle(params []string) {
	// send a message to another subject
	_, err := hf.MqttPublish("topic/reply", "it's a wasm module here")

	if err != nil {
		hf.Log("😡 ouch something bad is happening")
		hf.Log(err.Error())
	}
}
```


### Capsule MQTT publisher
> Publish MQTT messages from capsule

You can use a **WASM Capsule module** to MQTT messages, even if **Capsule** is not started in "mqtt" mode, for example from a **WASM CLI Capsule module**:

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
	_, err1stMsg := hf.MqttConnectPublish("127.0.0.1:1883", "sensor", "topic/sensor1", "🖐 Hello from WASM with MQTT 💜")
	_, err2ndMsg := hf.MqttConnectPublish("127.0.0.1:1883", "sensor", "topic/sensor2", "👋 Hello World 🌍")

    if err1stMsg != nil {
        errs = append(errs, err1stMsg.Error())
    }
    if err2ndMsg != nil {
        errs = append(errs, err2ndMsg.Error())
    }

    return "MQTT Rocks!", errors.New(strings.Join(errs, "|"))
}
```
> In this use case, you need to define the MQTT server and create a connection
