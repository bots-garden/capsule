package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

// the subject is defined when launching the capsule launcher
/*
capsule \
   -wasm=../wasm_modules/capsule-nats-subscriber/hello.wasm \
   -mode=nats \
   -natssrv=nats.devsecops.fun:4222 \
   -subject=notify
*/
func main() {
	hf.OnNatsMessage(Handle)
}

func Handle(params []string) {

	hf.Log("🟣👋 on subject: " + hf.NatsGetSubject() + ", 🎉 message " + params[0])

	// see: https://docs.nats.io/using-nats/developer/receiving/reply
	_, _ = hf.NatsReply("Hey I'm the other subscriber", 10)
	//_, _ = hf.NatsReply("Hola It's me again")

}

//export OnLoad
func OnLoad() {
	hf.Log("🙂 Hello from NATS subscriber")
	hf.Log(hf.GetHostInformation())
	hf.Log("👂Listening on: " + hf.NatsGetSubject())
	hf.Log("👋 NATS server: " + hf.NatsGetServer())

}

//export OnExit
func OnExit() {
	hf.Log("👋🤗 have a nice day")

	//hf.Log("Exit Error: " + hf.GetExitError())
	//hf.Log("Exit Code: " + hf.GetExitCode())

}
