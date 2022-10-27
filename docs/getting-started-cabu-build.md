# 🚀 Getting Started

### Build the **hello-world** function with **Cabu**

For building the WASM function, use the `cabu build` command:

```bash
cd hello-world
cabu build . hello-world.go hello-world.wasm
```

### Build the **hello-world** function with **TinyGo**

You can build the wasm module without **Cabu**. But you need to install **[Go](https://go.dev/doc/install)** and **[TinyGo](https://tinygo.org/getting-started/install/)**:

```bash
cd hello-world
go mod tidy
tinygo build -o hello-world.wasm -scheduler=none -target wasi ./hello-world.go
```

