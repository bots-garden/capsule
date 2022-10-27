# 🚀 Getting Started

## Serve the module frome an url

At start, **Capsule** can download the module function from an URL, store it to a directory of your choice, and then serve it.

```bash
capsule \
   -url=http://localhost:9090/hello-world/hello-world.wasm \
   -wasm=./tmp/hello-world.wasm \
   -mode=http \
   -httpPort=7070
```

> You can provide a wasm module through HTTP with any HTTP server:
> ```bash
> python3 -m http.server 9090
> ```
> or you can use a wasm registry, like https://wapm.io/

