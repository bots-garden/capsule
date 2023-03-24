module github.com/bots-garden/capsule/wasm_modules/capsule-hello-post

go 1.18

replace github.com/bots-garden/capsule/capsulemodule => ../../capsulemodule

replace github.com/bots-garden/capsule/commons => ../../commons

require (
	github.com/bots-garden/capsule/capsulemodule v0.3.1
	github.com/valyala/fastjson v1.6.4
)

require github.com/bots-garden/capsule/commons v0.3.1 // indirect
