module github.com/bots-garden/capsule/wasm_modules/registry-demo/capsule-hey

go 1.18

replace github.com/bots-garden/capsule/capsulemodule => ../../../capsulemodule
replace github.com/bots-garden/capsule/commons => ../../../commons


require (
	github.com/bots-garden/capsule/capsulemodule v0.0.0-20220815092415-964aa3e0fdc2
	github.com/tidwall/gjson v1.14.2
	github.com/tidwall/sjson v1.2.5
)

require (
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
)
