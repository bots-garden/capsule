module github.com/bots-garden/capsule/capsulemodule

go 1.18

replace github.com/bots-garden/capsule/commons => ../commons

require github.com/bots-garden/capsule/commons v0.0.0-20220918074717-33defaea8b7e

require (
	github.com/klauspost/compress v1.15.10 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/nats-io/jwt/v2 v2.3.0 // indirect
	github.com/nats-io/nats.go v1.17.0 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.0.0-20220829220503-c86fa9a7ed90 // indirect
	golang.org/x/time v0.0.0-20220722155302-e5dcc9cfc0b9 // indirect
)
