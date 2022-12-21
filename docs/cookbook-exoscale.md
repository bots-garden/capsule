# 🥘 CookBook

## Deploy a Capsule function on Exoscale

!!! info "What is Exoscale?"
    **Exoscale** is a Swiss cloud provider, that offers a wide range of cloud services like:
    - Compute instances
    - Kubernetes clusters
    - Object storage
    - Databases as a service (PostgreSQL, Redis, Kafka, Opensearch, MySql)
    - DNS
    It has several datacenters in Europe, which is really a nice, not only this gives us the possibility to deploy your
    applications in the datacenter that is closest to your users, but we can also create a high availability setup.

### Requirements

To follow this cookbook, you need to have the following tools installed:

- Kubernetes CLI (kubectl)
- Exoscale CLI (exo)
- Cabu (Capsule Builder)

There are several ways to install these tools, but I will show you how to install them on a macOs machine
using [brew](https://brew.sh/).

- **[kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)**:
  ```bash
  brew install kubernetes-cli
  ```
- **[exo](https://community.exoscale.com/documentation/tools/exoscale-command-line-interface/)**:
  ```bash
  brew tap exoscale/tap
  brew install exoscale-cli
  ```
- **[cabu]()**:
  ```bash
  CAPSULE_BUILDER_VERSION="v0.0.4"
  wget -O - https://raw.githubusercontent.com/bots-garden/capsule-function-builder/${CAPSULE_BUILDER_VERSION}/install-capsule-builder.sh | bash
  ```

## Create an SKS cluster on Exoscale

As soon as you have created an account, you are able to configure the CLI tool.

```bash
exo config
```

The CLI will guide you in the initial configuration.

Now we can create a Kubernetes cluster on Exoscale, in this example we're going to use as CNI `Cilium`:

```bash
exo compute security-group create sks-security-group

# Open 30000 to 32767 TCP from all sources for NodePort and LoadBalancer services
exo compute security-group rule add sks-security-group \
    --description "NodePort services" \
    --protocol tcp \
    --network 0.0.0.0/0 \
    --port 30000-32767

# Open 10250 TCP with the security group as a source
exo compute security-group rule add sks-security-group \
    --description "SKS kubelet" \
    --protocol tcp \
    --port 10250 \
    --security-group sks-security-group


# Open PING (ICMP type 8 & code 0) with the security group as a source for health checks
exo compute security-group rule add sks-security-group \
    --description "Cilium (healthcheck)" \
    --protocol icmp \
    --icmp-type 8 \
    --icmp-code 0 \
    --security-group sks-security-group

# Open 8472 UDP with the security group as a source for VXLAN communication between nodes
exo compute security-group rule add sks-security-group \
    --description "Cilium (vxlan)" \
    --protocol udp \
    --port 8472 \
    --security-group sks-security-group

# 4240 TCP with the security group as a source for network connectivity health API (health-checks)
exo compute security-group rule add sks-security-group \
    --description "Cilium (healthcheck)" \
    --protocol tcp \
    --port 4240 \
    --security-group sks-security-group


CLUSTER_NAME="swiss-alps"
CLUSTER_NODES=1
CLUSTER_ZONE="de-fra-1"

# Create the cluster

exo compute sks create ${CLUSTER_NAME} \
    --zone ${CLUSTER_ZONE} \
    --cni cilium \
    --service-level pro \
    --nodepool-name swiss-alps-nodepool \
    --nodepool-size ${CLUSTER_NODES} \
    --nodepool-security-group sks-security-group

# Get the kubeconfig
exo compute sks kubeconfig ${CLUSTER_NAME} kube-admin \
    --zone ${CLUSTER_ZONE} \
    --group system:masters > ${CLUSTER_NAME}.kubeconfig

export KUBECONFIG=${CLUSTER_NAME}.kubeconfig
```

**Output**:

```bash
 ✔ Creating Security Group "sks-security-group"... 3s
┼──────────────────┼──────────────────────────────────────┼
│  SECURITY GROUP  │                                      │
┼──────────────────┼──────────────────────────────────────┼
│ ID               │ acdaaf51-c2f1-402b-87aa-6b9be44b0e6e │
│ Name             │ sks-security-group                   │
│ Description      │                                      │
│ Ingress Rules    │ -                                    │
│ Egress Rules     │ -                                    │
│ External Sources │ -                                    │
┼──────────────────┼──────────────────────────────────────┼
 ✔ Adding rule to Security Group "sks-security-group"... 3s
┼──────────────────┼──────────────────────────────────────────────────────────────────────────────────────────────┼
│  SECURITY GROUP  │                                                                                              │
┼──────────────────┼──────────────────────────────────────────────────────────────────────────────────────────────┼
│ ID               │ acdaaf51-c2f1-402b-87aa-6b9be44b0e6e                                                         │
│ Name             │ sks-security-group                                                                           │
│ Description      │                                                                                              │
│ Ingress Rules    │                                                                                              │
│                  │   3b45b19d-e488-41d6-b78b-1a2410d6b16e   NodePort services   TCP   0.0.0.0/0   30000-32767   │
│                  │                                                                                              │
│ Egress Rules     │ -                                                                                            │
│ External Sources │ -                                                                                            │
┼──────────────────┼──────────────────────────────────────────────────────────────────────────────────────────────┼
 ✔ Adding rule to Security Group "sks-security-group"... 3s
┼──────────────────┼──────────────────────────────────────────────────────────────────────────────────────────────────────────┼
│  SECURITY GROUP  │                                                                                                          │
┼──────────────────┼──────────────────────────────────────────────────────────────────────────────────────────────────────────┼
│ ID               │ acdaaf51-c2f1-402b-87aa-6b9be44b0e6e                                                                     │
│ Name             │ sks-security-group                                                                                       │
│ Description      │                                                                                                          │
│ Ingress Rules    │                                                                                                          │
│                  │   00ce6538-ea7d-40ba-bd7a-5bf3e30c82ad   SKS kubelet         TCP   SG:sks-security-group   10250         │
│                  │   3b45b19d-e488-41d6-b78b-1a2410d6b16e   NodePort services   TCP   0.0.0.0/0               30000-32767   │
│                  │                                                                                                          │
│ Egress Rules     │ -                                                                                                        │
│ External Sources │ -                                                                                                        │
┼──────────────────┼──────────────────────────────────────────────────────────────────────────────────────────────────────────┼
 ✔ Adding rule to Security Group "sks-security-group"... 3s
┼──────────────────┼─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼
│  SECURITY GROUP  │                                                                                                                     │
┼──────────────────┼─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼
│ ID               │ acdaaf51-c2f1-402b-87aa-6b9be44b0e6e                                                                                │
│ Name             │ sks-security-group                                                                                                  │
│ Description      │                                                                                                                     │
│ Ingress Rules    │                                                                                                                     │
│                  │   00ce6538-ea7d-40ba-bd7a-5bf3e30c82ad   SKS kubelet            TCP    SG:sks-security-group   10250                │
│                  │   3b45b19d-e488-41d6-b78b-1a2410d6b16e   NodePort services      TCP    0.0.0.0/0               30000-32767          │
│                  │   f79c6e84-f430-4a0c-b9e7-a778a86c6db6   Cilium (healthcheck)   ICMP   SG:sks-security-group   ICMP code:0 type:8   │
│                  │                                                                                                                     │
│ Egress Rules     │ -                                                                                                                   │
│ External Sources │ -                                                                                                                   │
┼──────────────────┼─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼
 ✔ Adding rule to Security Group "sks-security-group"... 3s
┼──────────────────┼─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼
│  SECURITY GROUP  │                                                                                                                     │
┼──────────────────┼─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼
│ ID               │ acdaaf51-c2f1-402b-87aa-6b9be44b0e6e                                                                                │
│ Name             │ sks-security-group                                                                                                  │
│ Description      │                                                                                                                     │
│ Ingress Rules    │                                                                                                                     │
│                  │   00ce6538-ea7d-40ba-bd7a-5bf3e30c82ad   SKS kubelet            TCP    SG:sks-security-group   10250                │
│                  │   3b45b19d-e488-41d6-b78b-1a2410d6b16e   NodePort services      TCP    0.0.0.0/0               30000-32767          │
│                  │   f79c6e84-f430-4a0c-b9e7-a778a86c6db6   Cilium (healthcheck)   ICMP   SG:sks-security-group   ICMP code:0 type:8   │
│                  │   770ff5d5-4d80-49dd-81c9-64c7135e656b   Cilium (vxlan)         UDP    SG:sks-security-group   8472                 │
│                  │                                                                                                                     │
│ Egress Rules     │ -                                                                                                                   │
│ External Sources │ -                                                                                                                   │
┼──────────────────┼─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼
 ✔ Adding rule to Security Group "sks-security-group"... 3s
┼──────────────────┼─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼
│  SECURITY GROUP  │                                                                                                                     │
┼──────────────────┼─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼
│ ID               │ acdaaf51-c2f1-402b-87aa-6b9be44b0e6e                                                                                │
│ Name             │ sks-security-group                                                                                                  │
│ Description      │                                                                                                                     │
│ Ingress Rules    │                                                                                                                     │
│                  │   00ce6538-ea7d-40ba-bd7a-5bf3e30c82ad   SKS kubelet            TCP    SG:sks-security-group   10250                │
│                  │   3b45b19d-e488-41d6-b78b-1a2410d6b16e   NodePort services      TCP    0.0.0.0/0               30000-32767          │
│                  │   00fd91cc-9a61-4a07-9dcf-7ba8f5f5bd03   Cilium (healthcheck)   TCP    SG:sks-security-group   4240                 │
│                  │   f79c6e84-f430-4a0c-b9e7-a778a86c6db6   Cilium (healthcheck)   ICMP   SG:sks-security-group   ICMP code:0 type:8   │
│                  │   770ff5d5-4d80-49dd-81c9-64c7135e656b   Cilium (vxlan)         UDP    SG:sks-security-group   8472                 │
│                  │                                                                                                                     │
│ Egress Rules     │ -                                                                                                                   │
│ External Sources │ -                                                                                                                   │
┼──────────────────┼─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼
 ✔ Creating SKS cluster "swiss-alps"... 1m12s
 ✔ Adding Nodepool "swiss-alps-nodepool"... 3s
┼───────────────┼──────────────────────────────────────────────────────────────────┼
│  SKS CLUSTER  │                                                                  │
┼───────────────┼──────────────────────────────────────────────────────────────────┼
│ ID            │ 29c1bf5b-5219-497f-a041-6e718ed535d4                             │
│ Name          │ swiss-alps                                                       │
│ Description   │                                                                  │
│ Zone          │ de-fra-1                                                         │
│ Creation Date │ 2022-12-21 19:31:12 +0000 UTC                                    │
│ Endpoint      │ https://xxx.sks-de-fra-1.exo.io │
│ Version       │ 1.26.0                                                           │
│ Service Level │ pro                                                              │
│ CNI           │ cilium                                                           │
│ Add-Ons       │ exoscale-cloud-controller                                        │
│               │ metrics-server                                                   │
│ State         │ running                                                          │
│ Labels        │ n/a                                                              │
│ Nodepools     │ 4bf4d5ab-fc39-4a42-a9f5-45a5887147ee | swiss-alps-nodepool       │
┼───────────────┼──────────────────────────────────────────────────────────────────┼
```

You can now use the `kubectl` command to interact with your cluster:

```bash
export KUBECONFIG=${CLUSTER_NAME}.kubeconfig
kubectl get nodes
```

And you should see something like this:

```bash
NAME               STATUS   ROLES    AGE   VERSION
pool-11cb8-djvbs   Ready    <none>   95s   v1.26.0
```

## Create a new Capsule function

We're going to use the `cabu` cli to create a new function.

```bash
cabu generate service-get hello-swiss-robot
```

You should see something like this:

```bash
Unable to find image 'k33g/capsule-builder:0.0.4' locally
0.0.4: Pulling from k33g/capsule-builder
ff12a8b5aa85: Pull complete
afe5a53d5555: Pull complete
aa0b3dfc8ee1: Pull complete
e5320b25a8f8: Pull complete
6b9621a83b69: Pull complete
a50259c3276d: Pull complete
0d14b5adb1b9: Pull complete
0e019d286d0c: Pull complete
Digest: sha256:17649d47bcd43df4c9738323d0a8dc68a47d90b8c64639394a801d7a943907f2
Status: Downloaded newer image for k33g/capsule-builder:0.0.4
WARNING: The requested image's platform (linux/amd64) does not match the detected host platform (linux/arm64/v8) and no specific platform was requested
✅🙂 hello-swiss-robot function generated
```

This generated a new folder `hello-swiss-robot` with a `go.mod` file and a `hello-swiss-robot.go` file.

Open the `hello-swiss-robot.go` file and replace the content with this:

```go
package main

import (
    hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
    "github.com/mattes/go-asciibot"
    "github.com/tidwall/gjson"
)

func main() {
    hf.SetHandleHttp(Handle)
}

func Handle(req hf.Request) (resp hf.Response, errResp error) {

    robot := gjson.Get(req.Body, "robot")

    headers := map[string]string{
        "Content-Type": "text/plain",
    }

    resp = hf.Response{
        Body:    asciibot.MustGenerate(robot.Str),
        Headers: headers,
    }

    return resp, nil
}
```

And modify the `go.mod` file to add this line:

```go
module hello-swiss-robot

go 1.18

require (
github.com/bots-garden/capsule/capsulemodule v0.3.0
github.com/mattes/go -asciibot v0.0.0-20190603170252-3fa6d766c482
github.com/tidwall/gjson v1.14.4
)
```

To build the function, we will need a `Dockerfile`. Create a new file named `Dockerfile` and add this content:

```dockerfile
FROM k33g/capsule-builder:0.0.4
COPY go.mod ./
COPY hello-swiss-robot.go ./
RUN  go get -u ./...; go mod tidy;
RUN tinygo build -o hello-swiss-robot.wasm -target wasi hello-swiss-robot.go

FROM k33g/capsule-launcher:0.2.9
COPY --from=0 /home/function/hello-swiss-robot.wasm ./
EXPOSE 8080
CMD ["/capsule", "-wasm=./hello-swiss-robot.wasm", "-mode=http", "-httpPort=8080"]
```

We can now build the function thanks to the `Dockerfile` and its multi-stage capabilities:

```bash
IMAGE_NAME="capsule-hello-swiss-robot"
IMAGE_TAG="1.0.0"

docker login ...
docker build -t ${IMAGE_NAME} .

docker images | grep ${IMAGE_NAME}
```

Finally, we push the image to the Docker Hub:

```bash
IMAGE_NAME="capsule-hello-swiss-robot"
IMAGE_TAG="1.0.0"

docker tag ${IMAGE_NAME} ${DOCKER_USER}/${IMAGE_NAME}:${IMAGE_TAG}
docker push ${DOCKER_USER}/${IMAGE_NAME}:${IMAGE_TAG}
```

### Test the function locally

We can now test the function locally:

```bash
IMAGE_NAME="capsule-hello-swiss-robot"
IMAGE_TAG="1.0.0"

docker run -p 8080:8080 ${DOCKER_USER}/${IMAGE_NAME}:${IMAGE_TAG}
💊 Capsule ( v0.2.9 🎄 [Christmas tree] ) http server is listening on: 8080 🌍
```

Now let's test the function, we pass a JSON object to the function, with a `robot` field. This field is a 5 digit
hexadecimal number, which will be used to generate the robot.

```bash
curl -X POST http://localhost:8080 \
-H 'content-type: application/json' \
-d '{"robot": "10333"}'
    \.===./
    | d b |
     \_O_/
    /| []|\
  ()/|___|\()
     // \\
    _\\ //_
```

### Deploy the function on the SKS cluster

We can now deploy the function on the SKS cluster. For this I will use the `kubectl` commands. You can also crate all the Kubernetes resources as YAML files and use `kubectl apply -f ...`.

```bash
IMAGE_NAME="capsule-hello-swiss-robot"
IMAGE_TAG="1.0.0"
kubectl create deployment hello-swiss-robot-deployment --image ${DOCKER_USER}/${IMAGE_NAME}:${IMAGE_TAG} --port 8080
kubectl expose deployment hello-swiss-robot-deployment --port=8080 --target-port=8080 --type=LoadBalancer
```

### Call the function

Let's get the IP address of the service using the `kubectl` command:

```bash
kubectl get service hello-swiss-robot-deployment -o jsonpath="{.status.loadBalancer.ingress[].ip}"
```

And we test the function, to display a new robot:

```bash
```bash
curl -X POST http://<IP_ADDRESS>:8080 \
-H 'content-type: application/json' \
-d '{"robot": "10233"}'
```

And you should see something like this:

```bash
    \.===./
    | d b |
     \_O_/
    /| []|\
  ()/|___|\()
    . \_/  .
   . .:::.. .
```

And try another roboter with another hexadecimal number:

```bash
curl -X POST http://<IP_ADDRESS>:8080 \
-H 'content-type: application/json' \
-d '{"robot": "f2b11"}'
   _ _,_,_ _
   \( p q )/
     \_=_/
    /|(\)|\
   d |___| b
     (_|_)
     (o|o)
```

And that's it! You successfully deployed a Capsule function on a Exoscale SKS cluster.

### Housekeeping

!!! info "How to delete the cluster?"
    ```bash
    exo compute sks nodepool delete ${CLUSTER_NAME} swiss-alps-nodepool -z ${CLUSTER_ZONE}
    exo compute sks delete ${CLUSTER_NAME}
    exo compute security-group delete sks-security-group
    ```
