# 🐳 Capsule HTTP Docker image

!!! info "Capsule HTTP Docker images v0.3.6 🫐 [blueberries]"
    - `botsgarden/capsule-http-linux-arm64:0.3.6`
    - `botsgarden/capsule-http-linux-amd64:0.3.6`
    - `botsgarden/capsule-http-darwin-arm64:0.3.6`
    - `botsgarden/capsule-http-darwin-amd64:0.3.6`

> https://hub.docker.com/repositories/botsgarden

## How to use it

```bash
GOOS="linux" 
GOARCH="arm64"
IMAGE_TAG="0.3.6"
IMAGE_NAME="botsgarden/capsule-http-${GOOS}-${GOARCH}"

docker run \
  -p 8080:8080 \
  -v $(pwd)/functions/hello-world:/app --rm ${IMAGE_NAME}:${IMAGE_TAG} \
  /capsule-http \
    --wasm=./app/hello-world.wasm \
    --httpPort=8080
```
