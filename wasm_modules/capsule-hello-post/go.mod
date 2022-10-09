module github.com/bots-garden/capsule/wasm_modules/capsule-hello-post

go 1.18

replace github.com/bots-garden/capsule/capsulemodule => ../../capsulemodule

replace github.com/bots-garden/capsule/commons => ../../commons

require (
	github.com/bots-garden/capsule/capsulemodule v0.2.7
	github.com/tidwall/gjson v1.14.3
	github.com/tidwall/sjson v1.2.5
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/bots-garden/capsule/commons v0.2.7 // indirect
	github.com/gofiber/fiber/v2 v2.38.1 // indirect
	github.com/klauspost/compress v1.15.11 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.40.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20221006211917-84dc82d7e875 // indirect
)
