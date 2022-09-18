# `nats/listen.go`

This execution mode (`nats`) "transforms" **capsule** in a NATS subscriber.
```bash
capsule \
   -wasm=../wasm_modules/capsule-nats-subscriber/hello.wasm \
   -mode=nats \
   -natssrv=nats.devsecops.fun:4222 \
   -subject=ping
```

- **capsule** is listening on the **subject** "ping" (like a MQTT topic)
- every time a message is posted on "ping" **capsule** will run a wasm module and call the "`Handle`" function of the module

The `Listen` function is called by `../../main.go`:
- it stores/saves the subject string (then it will be available for the wasm module thanks to a host function)
- it creates the NATS connection
- it subscribes to the subject

> helpers to create + get the connection and to store the topic are located in `commons/nats.go`
