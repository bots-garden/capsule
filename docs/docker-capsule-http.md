# 🐳 Capsule HTTP Docker image

!!! info "Capsule HTTP Docker images v0.3.7 🥦 [broccoli]"
    - `botsgarden/capsule-http-linux-arm64:0.3.7`
    - `botsgarden/capsule-http-linux-amd64:0.3.7`
    - `botsgarden/capsule-http-darwin-arm64:0.3.7`
    - `botsgarden/capsule-http-darwin-amd64:0.3.7`

> https://hub.docker.com/repositories/botsgarden

## How to use it

```bash
GOOS="linux" 
GOARCH="arm64"
IMAGE_TAG="0.3.7"
IMAGE_NAME="botsgarden/capsule-http-${GOOS}-${GOARCH}"

docker run \
  -p 8080:8080 \
  -v $(pwd)/functions/hello-world:/app --rm ${IMAGE_NAME}:${IMAGE_TAG} \
  /capsule-http \
    --wasm=./app/hello-world.wasm \
    --httpPort=8080
```

## Dockerize Capsule HTTP and a WASM module

Create a new `Dockerfile`:

```dockerfile
FROM botsgarden/capsule-http-linux-arm64:0.3.7
COPY hello-world.wasm .
EXPOSE 8080
CMD ["/capsule-http", "--wasm=./hello-world.wasm", "--httpPort=8080"]
```

Build the image:

```bash
IMAGE_NAME="demo-capsule-http"
docker login -u ${DOCKER_USER} -p ${DOCKER_PWD}
docker build -t ${IMAGE_NAME} . 

docker images | grep ${IMAGE_NAME}
```

Run the container:

```bash
IMAGE_NAME="demo-capsule-http"
docker run \
  -p 8080:8080 \
  --rm ${IMAGE_NAME}
```

Call the service:

```bash
JSON_DATA='{"name":"Bob Morane","age":42}'
curl -X POST http://localhost:8080 \
  -H 'Content-Type: application/json; charset=utf-8' \
  -d "${JSON_DATA}"
```





