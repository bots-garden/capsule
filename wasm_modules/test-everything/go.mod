module github.com/bots-garden/capsule/wasm_modules/capsule-hello-post

go 1.18

replace github.com/bots-garden/capsule/capsulemodule => ../../capsulemodule

replace github.com/bots-garden/capsule/commons => ../../commons

require (
	github.com/bots-garden/capsule/capsulemodule v0.2.9
	github.com/tidwall/gjson v1.14.3
	github.com/tidwall/sjson v1.2.5
)

require (
	github.com/bots-garden/capsule/commons v0.2.9 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
)
