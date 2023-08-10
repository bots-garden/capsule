module github.com/bots-garden/capsule/capsule-cli

go 1.20

require github.com/tetratelabs/wazero v1.4.0 // indirect

require github.com/bots-garden/capsule-host-sdk v0.0.7

require github.com/bots-garden/capsule/capsule-http v0.0.0-20230712075637-624af48d1327

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-resty/resty/v2 v2.7.0 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/gofiber/fiber/v2 v2.48.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/inconshreveable/log15 v3.0.0-testing.5+incompatible // indirect
	github.com/inconshreveable/log15/v3 v3.0.0-testing.5 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/redis/go-redis/v9 v9.0.5 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.48.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.ngrok.com/muxado/v2 v2.0.0 // indirect
	golang.ngrok.com/ngrok v1.3.1 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/term v0.11.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace github.com/bots-garden/capsule-host-sdk => ../../capsule-host-sdk

replace github.com/bots-garden/capsule-module-sdk => ../../capsule-module-sdk
