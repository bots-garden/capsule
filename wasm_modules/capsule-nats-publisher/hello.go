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

	// Publish and wait for an answer
	res, err2ndMsg := hf.NatsConnectRequest("nats.devsecops.fun:4222", "notify", "👋 Hello World 🌍", 1)

	hf.Log("👋 -> " + res)

	msg, err := hf.NatsConnectPublish("nats.devsecops.fun:4222", "ping", "🖐 Hello from WASM with Nats 💜")
	if err != nil {
		hf.Log("🔴" + err.Error())
	} else {
		hf.Log("🔵" + msg)
	}

	if err1stMsg != nil {
		errs = append(errs, err1stMsg.Error())
	}
	if err2ndMsg != nil {
		errs = append(errs, err2ndMsg.Error())
	}

	return "NATS Rocks!", errors.New(strings.Join(errs, "|"))
}
