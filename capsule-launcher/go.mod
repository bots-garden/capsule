module github.com/bots-garden/capsule/capsule-launcher

go 1.18

replace (
	github.com/bots-garden/capsule/commons => ../commons
	github.com/bots-garden/capsule/mqttconn => ../mqttconn
	github.com/bots-garden/capsule/natsconn => ../natsconn
)

require (
	github.com/bots-garden/capsule/commons v0.2.6
	github.com/bots-garden/capsule/mqttconn v0.0.0-20221008083118-2753339a260f
	github.com/bots-garden/capsule/natsconn v0.0.0-20221008083118-2753339a260f
	github.com/couchbase/gocb/v2 v2.5.3
	github.com/eclipse/paho.mqtt.golang v1.4.1
	github.com/gin-gonic/gin v1.8.1
	github.com/go-redis/redis/v9 v9.0.0-beta.3
	github.com/go-resty/resty/v2 v2.7.0
	github.com/nats-io/nats.go v1.17.0
	github.com/shirou/gopsutil/v3 v3.22.9
	github.com/tetratelabs/wazero v1.0.0-pre.2
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/couchbase/gocbcore/v10 v10.1.5 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.1 // indirect
	github.com/goccy/go-json v0.9.11 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lufia/plan9stats v0.0.0-20220913051719-115f729f3c8c // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nats-io/jwt/v2 v2.3.0 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
	github.com/power-devops/perfstat v0.0.0-20220216144756-c35f1ee13d7c // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	golang.org/x/crypto v0.0.0-20221005025214-4161e89ecf1b // indirect
	golang.org/x/net v0.0.0-20221004154528-8021a29435af // indirect
	golang.org/x/sync v0.0.0-20220929204114-8fcdb60fdcc0 // indirect
	golang.org/x/sys v0.0.0-20221006211917-84dc82d7e875 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20220722155302-e5dcc9cfc0b9 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
