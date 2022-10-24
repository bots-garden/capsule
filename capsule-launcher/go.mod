module github.com/bots-garden/capsule/capsule-launcher

go 1.18

replace (
	github.com/bots-garden/capsule/commons => ../commons
	github.com/bots-garden/capsule/mqttconn => ../mqttconn
	github.com/bots-garden/capsule/natsconn => ../natsconn
)

require (
	github.com/bots-garden/capsule/commons v0.0.0-00010101000000-000000000000
	github.com/bots-garden/capsule/mqttconn v0.0.0-00010101000000-000000000000
	github.com/bots-garden/capsule/natsconn v0.0.0-00010101000000-000000000000
	github.com/couchbase/gocb/v2 v2.6.0
	github.com/eclipse/paho.mqtt.golang v1.4.1
	github.com/go-redis/redis/v9 v9.0.0-rc.1
	github.com/go-resty/resty/v2 v2.7.0
	github.com/gofiber/fiber/v2 v2.39.0
	github.com/google/uuid v1.3.0
	github.com/nats-io/nats.go v1.18.0
	github.com/shirou/gopsutil/v3 v3.22.9
	github.com/tetratelabs/wazero v1.0.0-pre.2
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/couchbase/gocbcore/v10 v10.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/klauspost/compress v1.15.0 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/nats-io/jwt/v2 v2.3.0 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.40.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	golang.org/x/crypto v0.0.0-20220829220503-c86fa9a7ed90 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab // indirect
	golang.org/x/time v0.1.0 // indirect
)
