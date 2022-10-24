module github.com/bots-garden/capsule/wasm_modules/capsule-hello-post-before

go 1.18

replace github.com/bots-garden/capsule/capsulemodule => ../../capsulemodule

replace github.com/bots-garden/capsule/commons => ../../commons

require github.com/bots-garden/capsule/capsulemodule v0.0.0-00010101000000-000000000000

require github.com/bots-garden/capsule/commons v0.2.8 // indirect
