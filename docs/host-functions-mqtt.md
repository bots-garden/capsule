# Host functions

## MQTT

> You must use the `"mqtt"` mode of **Capsule** as the MQTT connection is defined at the start of **Capsule** and shared with the WASM module:

```bash
capsule \
   -wasm=../wasm_modules/capsule-mqtt-subscriber/hello.wasm \
   -mode=mqtt \
   -mqttsrv=127.0.0.1:1883 \
   -topic=topic/sensor0 \
   -clientId=sensor_id1
```

### MQTT Handler as a Subscriber
> 🖐 you have to call `hf.OnMqttMessage(Handle)` from the `main` function.

```golang
package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

// the topic is defined when launching the capsule launcher
func main() {
	hf.OnMqttMessage(Handle)
}

func Handle(params []string) {
    message := params[0]
	hf.Log("👋 you get a message on topic " + hf.MqttGetTopic() + ": " + message)

	// we use the connection of the launcher (capsule)
	_, err := hf.MqttPublish("topic/reply", "it's a wasm module here")

	if err != nil {
		hf.Log("😡 ouch something bad is happening")
		hf.Log(err.Error())
	}
}
```
