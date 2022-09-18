module github.com/bots-garden/capsule/capsule-ctl

go 1.18

replace github.com/bots-garden/capsule/commons => ../commons

require (
	github.com/bots-garden/capsule/commons v0.0.0-00010101000000-000000000000
	github.com/go-resty/resty/v2 v2.7.0
)

require golang.org/x/net v0.0.0-20211029224645-99673261e6eb // indirect
