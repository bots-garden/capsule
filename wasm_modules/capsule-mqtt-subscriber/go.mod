module github.com/bots-garden/capsule/wasm_modules/capsule-nats-subscriber

go 1.18

replace github.com/bots-garden/capsule/capsulemodule => ../../capsulemodule

replace github.com/bots-garden/capsule/commons => ../../commons

require github.com/bots-garden/capsule/capsulemodule v0.0.0-20220918083243-1eea1d761338

require github.com/bots-garden/capsule/commons v0.2.2 // indirect
