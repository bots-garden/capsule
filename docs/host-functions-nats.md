# Host functions

## Nats

*`NatsPublish(subject string, message string)`*: publish a message on a subject
```golang
_, err := hf.NatsPublish("notify", "it's a wasm module here")
```
> You must use the `"nats"` mode of **Capsule** as the NATS connection is defined at the start of **Capsule** and shared with the WASM module:

```bash
capsule \
 -wasm=../wasm_modules/capsule-nats-subscriber/hello.wasm \
 -mode=nats \
 -natssrv=nats.devsecops.fun:4222 \
 -subject=ping
```

*`NatsReply(message string, timeout uint32)`*: publish a message on the current subject and wait for an answer
```golang
_, err := hf.NatsReply("it's a wasm module here", 10)
```
> You must use the `"nats"` mode of **Capsule** as the NATS connection and the subscription are defined at the start of **Capsule** and shared with the WASM module.



*`NatsGetSubject()`*: get the subject listened by the **Capsule** launcher
```golang
hf.Log("👂Listening on: " + hf.NatsGetSubject())
```


*`NatsGetServer()`*: get the connected NATS server
```golang
hf.Log("👋 NATS server: " + hf.NatsGetServer())
```

*`NatsConnectPublish(server string, subject string, message string)`*: connect to a NATS server and send a message on a subject
```golang
_, err := hf.NatsConnectPublish("nats.devsecops.fun:4222", "ping", "🖐 Hello from WASM with Nats 💜")
```
> You can use this function with all the running modes of **Capsule**

*`NatsConnectPublish(server string, subject string, message string, timeout uint32)`*: connect to a NATS server and send a message on a subject
```golang
answer, err := hf.NatsConnectRequest("nats.devsecops.fun:4222", "notify", "👋 Hello World 🌍", 1)
```
> You can use this function with all the running modes of **Capsule**
