module github.com/bots-garden/capsule/wasm_modules/capsule-hello-get

go 1.18

replace github.com/bots-garden/capsule/capsulemodule => ../../capsulemodule

replace github.com/bots-garden/capsule/commons => ../../commons

require github.com/bots-garden/capsule/capsulemodule v0.0.0-20220903173120-dd90b15300ea

require github.com/bots-garden/capsule/commons v0.0.0-20220903173120-dd90b15300ea // indirect
