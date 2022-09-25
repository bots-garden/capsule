module github.com/bots-garden/capsule/capsulemodule

go 1.18

replace github.com/bots-garden/capsule/capsulemodule => ../capsulemodule

replace github.com/bots-garden/capsule/commons => ../commons

replace (
	github.com/bots-garden/capsule v0.2.2 => ../
	github.com/bots-garden/capsule v0.2.3 => ../
	github.com/bots-garden/capsule v0.2.4 => ../
	github.com/bots-garden/capsule v0.2.5 => ../
)

replace (
	github.com/bots-garden/capsule/commons v0.2.2 => ../commons
	github.com/bots-garden/capsule/commons v0.2.3 => ../commons
	github.com/bots-garden/capsule/commons v0.2.4 => ../commons
	github.com/bots-garden/capsule/commons v0.2.5 => ../commons
)

require github.com/bots-garden/capsule/commons v0.2.2
