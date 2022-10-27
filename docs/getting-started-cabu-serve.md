# 🚀 Getting Started


## Serve the **hello-world** function

Before serving the function, you need to install **Capsule**: see the [install section](install.md)

### Serve the function

```bash
capsule \
  -wasm=./hello-world.wasm \
  -mode=http \
  -httpPort=8080
```
> Reach [http://localhost:8080](http://localhost:8080) with your browser

### Serve the function with the **Capsule Docker image**

```bash
docker run \
  -p 8080:8080 \
  -v $(pwd):/app --rm k33g/capsule-launcher:0.2.8 \
  /capsule \
  -wasm=./app/hello-world.wasm \
  -mode=http \
  -httpPort=8080
```

> The **Capsule Docker image** is an external project: [https://github.com/bots-garden/capsule-docker-image](https://github.com/bots-garden/capsule-docker-image)
