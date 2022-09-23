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
