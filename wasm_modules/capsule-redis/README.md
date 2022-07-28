# Capsule Redis

This wasm module is used by the `cli` mode

```bash
#!/bin/bash
REDIS_ADDR="localhost:6379" REDIS_PWD="" go run main.go \
   -wasm=../wasm_modules/capsule-redis/hello.wasm \
   -mode=cli
```
