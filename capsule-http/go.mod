module github.com/bots-garden/capsule/capsule-http

go 1.20

require github.com/tetratelabs/wazero v1.1.0 // indirect

require (
	github.com/bots-garden/capsule-host-sdk v0.0.1
	github.com/gofiber/fiber/v2 v2.44.0
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/go-resty/resty/v2 v2.7.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/klauspost/compress v1.16.5 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/philhofer/fwd v1.1.2 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/savsgio/dictpool v0.0.0-20221023140959-7bf2e61cea94 // indirect
	github.com/savsgio/gotils v0.0.0-20230208104028-c358bd845dee // indirect
	github.com/tinylib/msgp v1.1.8 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.47.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
)

replace github.com/bots-garden/capsule-host-sdk => ../../capsule-host-sdk

//replace github.com/bots-garden/capsule-module-sdk => ../../capsule-module-sdk
