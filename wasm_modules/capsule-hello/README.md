# Capsule hello

This wasm module is used by the `http` mode

```bash
#!/bin/bash
export MESSAGE="💊 Capsule Rocks 🚀"
go run main.go \
   -wasm=../wasm_modules/capsule-hello/hello.wasm \
   -mode=http \
   -httpPort=7070
```
