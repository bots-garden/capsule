# Capsule Redis

This wasm module is used by the `cli` mode. You need to start a Redis server (`redis-server`)

```bash
REDIS_ADDR="localhost:6379" REDIS_PWD="" ./capsule \
   -wasm=./hello.wasm \
   -mode=cli
```
