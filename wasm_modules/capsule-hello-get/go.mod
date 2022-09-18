module github.com/bots-garden/capsule/wasm_modules/capsule-hello-get

go 1.18

replace github.com/bots-garden/capsule/capsulemodule => ../../capsulemodule
replace github.com/bots-garden/capsule/commons => ../../commons

require github.com/bots-garden/capsule/capsulemodule v0.0.0-20220905061317-2c8d99d2bdcd

require github.com/bots-garden/capsule/commons v0.0.0-20220905061317-2c8d99d2bdcd // indirect
