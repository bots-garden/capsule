module github.com/bots-garden/capsule/wasm_modules/capsule-couchbase/cli-module

go 1.18

replace github.com/bots-garden/capsule/capsulemodule => ../../../capsulemodule

require (
	github.com/bots-garden/capsule/capsulemodule v0.0.0-20220815092415-964aa3e0fdc2
	github.com/tidwall/gjson v1.14.2
)

require (
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
)
