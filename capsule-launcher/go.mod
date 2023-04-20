module github.com/bots-garden/capsule/capsule-launcher

go 1.20

replace (
	github.com/bots-garden/capsule/commons => ../commons
	github.com/bots-garden/capsule/mqttconn => ../mqttconn
	github.com/bots-garden/capsule/natsconn => ../natsconn
)

require (
	github.com/bots-garden/capsule/commons v0.3.1
	github.com/bots-garden/capsule/mqttconn v0.0.0-20230207102522-ad91242de694
	github.com/bots-garden/capsule/natsconn v0.0.0-20230207102522-ad91242de694
	github.com/couchbase/gocb/v2 v2.6.2
	github.com/eclipse/paho.mqtt.golang v1.4.2
	github.com/go-redis/redis/v9 v9.0.0-rc.2
	github.com/go-resty/resty/v2 v2.7.0
	github.com/gofiber/fiber/v2 v2.42.0
	github.com/google/uuid v1.3.0
	github.com/nats-io/nats.go v1.24.0
	github.com/tetratelabs/wazero v1.0.2
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/couchbase/gocbcore/v10 v10.2.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/klauspost/compress v1.16.3 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/nats-io/jwt/v2 v2.3.0 // indirect
	github.com/nats-io/nkeys v0.4.4 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/philhofer/fwd v1.1.2 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/savsgio/dictpool v0.0.0-20221023140959-7bf2e61cea94 // indirect
	github.com/savsgio/gotils v0.0.0-20230208104028-c358bd845dee // indirect
	github.com/tinylib/msgp v1.1.8 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.45.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/time v0.1.0 // indirect
)
