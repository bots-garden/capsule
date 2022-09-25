package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

// the subject is defined when launching the capsule launcher
/*
capsule \
   -wasm=../wasm_modules/capsule-mqtt-subscriber/hello.wasm \
   -mode=mqtt \
   -mqttsrv=127.0.0.1:1883 \
   -topic=topic/sensor0 \
   -clientId=sensor
*/
func main() {
	hf.OnMqttMessage(Handle)
}

func Handle(params []string) {

	hf.Log("👋 on topic: " + hf.MqttGetTopic() + ", 🎉 message" + params[0])

	// we use the connection of the launcher (capsule)
	_, err := hf.MqttPublish("topic/reply", "it's a wasm module here")

	if err != nil {
		hf.Log("😡 ouch something bad is happening")
		hf.Log(err.Error())
	}
}

//export OnLoad
func OnLoad() {
	hf.Log("🙂 Hello from MQTT subscriber")
	hf.Log(hf.GetHostInformation())
	hf.Log("👂Listening on: " + hf.MqttGetTopic())
	hf.Log("👋 MQTT server: " + hf.MqttGetServer())

}

//export OnExit
func OnExit() {
	hf.Log("👋🤗 have a nice day")

	//hf.Log("Exit Error: " + hf.GetExitError())
	//hf.Log("Exit Code: " + hf.GetExitCode())
}
