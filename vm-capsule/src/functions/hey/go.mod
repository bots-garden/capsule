module github.com/bots-garden/capsule-faas-demo/hey

go 1.18

replace github.com/bots-garden/capsule/capsulemodule => ../../../../capsulemodule

require (
	github.com/bots-garden/capsule/capsulemodule v0.0.0-20220813055457-58259c3a7008
	github.com/tidwall/gjson v1.14.2
	github.com/tidwall/sjson v1.2.5
)

require (
	github.com/bots-garden/capsule/commons v0.0.0-20220821060842-d1dc9f030021 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
)
