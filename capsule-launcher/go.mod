module github.com/bots-garden/capsule/capsule-launcher

go 1.18

replace (
	github.com/bots-garden/capsule/commons => ../commons
	github.com/bots-garden/capsule/mqttconn => ../mqttconn
	github.com/bots-garden/capsule/natsconn => ../natsconn
)

require (
	github.com/bots-garden/capsule/commons v0.2.9
	github.com/bots-garden/capsule/mqttconn v0.0.0-20221027063202-a59dfbe32b65
	github.com/bots-garden/capsule/natsconn v0.0.0-20221027063202-a59dfbe32b65
	github.com/couchbase/gocb/v2 v2.6.0
	github.com/eclipse/paho.mqtt.golang v1.4.2
	github.com/go-redis/redis/v9 v9.0.0-rc.1
	github.com/go-resty/resty/v2 v2.7.0
	github.com/gofiber/fiber/v2 v2.39.0
	github.com/google/uuid v1.3.0
	github.com/nats-io/nats.go v1.18.0
	github.com/tetratelabs/wazero v1.0.0-pre.3
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/couchbase/gocbcore/v10 v10.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/klauspost/compress v1.15.12 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/nats-io/jwt/v2 v2.3.0 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/rivo/uniseg v0.4.2 // indirect
	github.com/stretchr/objx v0.4.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.41.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/crypto v0.1.0 // indirect
	golang.org/x/net v0.1.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/time v0.1.0 // indirect
)
