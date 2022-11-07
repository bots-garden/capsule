module github.com/bots-garden/capsule/wasm_modules/capsule-http-get

go 1.18

replace github.com/bots-garden/capsule/capsulemodule => ../../capsulemodule

replace github.com/bots-garden/capsule/commons => ../../commons

require github.com/bots-garden/capsule/capsulemodule v0.2.9

require github.com/bots-garden/capsule/commons v0.2.9 // indirect
