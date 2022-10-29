# 🥘 CookBook

## Deploy a Capsule function on Fly.io

!!! info "What is Fly.io?"
    **Fly** is a platform for running full stack apps and databases.
    It's a very very easy way to deploy a container.

### Requirements

First, you need an account on **[Fly.io](https://fly.io/)**, then you will need to install some tools.

> The install commands I used were tested on macOs and Ubuntu. I'm using **[brew](https://brew.sh/)**, but there are severeal other ways to install all the needed tools.

- **[flyctl, the Fly.io CLI](https://fly.io/docs/hands-on/install-flyctl/)**:
  ```bash
  brew install superfly/tap/flyctl
  ```

Get your token from your Fly.io account, and set a `FLY_ACCESS_TOKEN` variable with the token's value.

## Create a new Capsule function

*hello.go*:
```golang
package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func main() {
	hf.SetHandleHttp(Handle)
}

func Handle(request hf.Request) (response hf.Response, errResp error) {

	name := gjson.Get(request.Body, "name")

	headersResp := map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}

	jsondoc := `{"message": ""}`
	jsondoc, _ = sjson.Set(jsondoc, "message", "👋 hello " + name.Str)

	return hf.Response{Body: jsondoc, Headers: headersResp}, nil
}

```

*go.mod*:
```text
module hello

go 1.18
```

## Dockerize the function

In the directory of the function, create a `Dockerfile`:

*Dockerfile*:
```dockerfile
FROM k33g/capsule-launcher:0.2.9
ADD hello.wasm ./
EXPOSE 8080
CMD ["/capsule", "-wasm=./hello.wasm", "-mode=http", "-httpPort=8080"]
```

## Build the function

Then build the wasm module with **[TinyGo](https://tinygo.org/)**
```bash
tinygo build -o hello.wasm -scheduler=none -target wasi ./hello.go
```

!!! info "If you don't want to install all the toolchain (Go, TinyGo)"
    you can install **[CaBu](getting-started-cabu-inst.md)** and compile the wasm module like this:
    ```bash
    cabu build . hello.go hello.wasm
    ```
    or use [multi-stage](https://docs.docker.com/build/building/multi-stage/) builds to first build the wasm function and then to create the smallest image as possible to serve the function. 👀 Look at [Capsule function on Civo](cookbook-fly.md).

## Build and push the Docker image

### Build the Docker image

Type the below commands to build the Docker image:
```bash
IMAGE_NAME="capsule-hello-demo"
IMAGE_TAG="0.0.0"
docker build -t ${IMAGE_NAME} .
```
!!! info "Test it"
    Run the below command:
    ```bash
    docker run -p 8080:8080 -it ${IMAGE_NAME}
    ```
    Then call the function:
    ```
    curl -X POST http://localhost:8080 \
    -H 'content-type: application/json' \
    -d '{"name": "Bob"}'
    ```
    You should get: `{"message":"👋 hello Bob"}`

### Push the Docker image to the Docker Hub

Type the below commands to publish the Docker image:
```bash
IMAGE_NAME="capsule-hello-demo"
IMAGE_TAG="0.0.0"
docker login -u ${DOCKER_USER} -p ${DOCKER_PWD}
docker tag ${IMAGE_NAME} ${DOCKER_USER}/${IMAGE_NAME}:${IMAGE_TAG}
docker push ${DOCKER_USER}/${IMAGE_NAME}:${IMAGE_TAG}
```

## 🚀 Deploy the function on Fly.io

Before the first deployment, you need to create the application on **Fly.io**:
```bash
# Create the application, only at the first deployment
APPLICATION_NAME="capsule-hello-demo"
flyctl apps create ${APPLICATION_NAME} --json
```

And then, run the below commands to deploy the function:

```bash
IMAGE_NAME="capsule-hello-demo"
IMAGE_TAG="0.0.0"
APPLICATION_NAME="capsule-hello-demo"

flyctl deploy \
  --app ${APPLICATION_NAME} \
  --image ${DOCKER_USER}/${IMAGE_NAME}:${IMAGE_TAG} \
  --verbose --json
```

**Output**:
```bash
==> Verifying app config
--> Verified app config
==> Building image
Searching for image 'k33g/capsule-hello-demo:0.0.3' remotely...
image found: img_ox20prmgxx3vj1zq
==> Creating release
--> release v2 created

--> You can detach the terminal anytime without stopping the deployment
==> Monitoring deployment

 1 desired, 1 placed, 0 healthy, 0 unhealthy [health checks: 1 total, 1 pa
 1 desired, 1 placed, 1 healthy, 0 unhealthy [health checks: 1 total, 1 pa
 1 desired, 1 placed, 1 healthy, 0 unhealthy [health checks: 1 total, 1 passing]
--> v0 deployed successfully
```

The function is deployed 🎉

## Test the function

Run the below commands:
```bash
APPLICATION_NAME="capsule-hello-demo"
# The function url follows the following form:
URL="https://${{APPLICATION_NAME}}.fly.dev"

curl -X POST ${URL} \
-H 'content-type: application/json' \
-d '{"name": "Bob"}'
```
You should get: `{"message":"👋 hello Bob"}`

🎉 **Fly.io** is an excellent option for deploying a container with ease.


