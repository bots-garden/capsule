package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

func main() {
	hf.OnNatsMessage(Handle)
}

func Handle(params []string) {
	hf.Log("🎉 " + params[0])
}

//export OnLoad
func OnLoad() {
	hf.Log("🙂 Hello from NATS subscriber")
}
