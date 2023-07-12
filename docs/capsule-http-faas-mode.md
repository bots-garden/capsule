# 🚀 Capsule HTTP FaaS mode

> - Released since Capsule HTTP `v0.3.7 🥦 [broccoli]`
> - This is work in progress 🚧

A Capsule HTTP server can start/spawn other Capsule HTTP server processes.

## Requirements

### Install the last version of Capsule HTTP

```bash
VERSION="v0.4.1" OS="linux" ARCH="arm64"
wget -O capsule-http https://github.com/bots-garden/capsule/releases/download/${VERSION}/capsule-http-${VERSION}-${OS}-${ARCH}
chmod +x capsule-http
sudo cp capsule-http  /usr/local/bin/capsule-http
rm capsule-http
capsule-http --version
```
> Set the appropriate OS, ARCH and VERSION

### Install the last version of CapsCtl

**CapsCtl** is a CLI to send commands to the Capsule HTTP server when it is unning in **FaaS** mode.

```bash
VERSION="v0.4.1" OS="linux" ARCH="arm64"
wget -O capsctl https://github.com/bots-garden/capsule/releases/download/${VERSION}/capsctl-${VERSION}-${OS}-${ARCH}
chmod +x capsctl
sudo cp capsctl  /usr/local/bin/capsctl
rm capsctl
capsctl --version
```

## Start Capsule HTTP FaaS mode

```bash
CAPSULE_FAAS_TOKEN="ILOVEPANDAS" \
capsule-http \
--wasm=./functions/index-page/index-page.wasm \
--httpPort=8080 \
--faas=true
```

> **Remarks:**: if you use SSL certificates, use these options:
> - `--crt=faas.capsule.foundation.crt`
> - `--key=faas.capsule.foundation.key`


You should get an output like this:
```
2023/05/29 15:12:18 🚀 faas mode activated!
2023/05/29 15:12:18 📦 wasm module loaded: ./functions/index-page/index-page.wasm
2023/05/29 15:12:18 💊 Capsule [HTTP] v0.4.1 🫑 [pepper]
 http server is listening on: 8080 🌍
```

> **Remarks:**
> - the wasm file (`--wasm`) is optional (a default message is served if not specified)
> - `CAPSULE_FAAS_TOKEN` is used to authenticate the `capsctl` CLI

## Start a function

With the FaaS mode activated, you can start functions. It' like running another **Capsule HTTP** processes (one wasm function == one Capsule HTTP process).

```bash
export CAPSULE_FAAS_TOKEN="ILOVEPANDAS"
# the main Capsule HTTP process is listeninhg on port 8080
export CAPSULE_MAIN_PROCESS_URL="http://localhost:8080" 

capsctl \
    --cmd=start \
    --stopAfter=10 \
    --name=hello \
    --revision=green \
    --description="this the hello module, green revision" \
    --env='["MESSAGE=🟢","GREETING=🤗"]' \
    --wasm=./functions/hello-green/hello-green.wasm
```
> - `--stopAfter=10` this will stop the Capsule HTTP server process after 10 seconds ()optional
> - `--stopAfter` is not mandatory (then the Capsule HTTP server process will never stop)
> - if the process is stopped, the Capsule HTTP server will be restarted at next call
> - `--description=` is optional
> - `--env='["MESSAGE=🟢","GREETING=🤗"]'` allows to pass environment variables to the function (optional)
> - `--wasm`: where to find the wasm file

**Now you can use this URL `http://localhost:8080/functions/hello/green` to call the hello green function**

### Call the hello green function

```bash
curl -X POST http://localhost:8080/functions/hello/green \
-H 'Content-Type: text/plain; charset=utf-8' \
-d "Bob Morane"
```

### Default revision

If you don't specify a revision, the default revision is called **default**, then you can call the function like this:

```bash
curl -X POST http://localhost:8080/functions/hello \
-H 'Content-Type: text/plain; charset=utf-8' \
-d "Bob Morane"
```

Or like this:

```bash
curl -X POST http://localhost:8080/functions/hello/default \
-H 'Content-Type: text/plain; charset=utf-8' \
-d "Bob Morane"
```

> 👋 the revision concept is useful to handle several version of a wasm module/function.

## Launch another function

```bash
export CAPSULE_FAAS_TOKEN="ILOVEPANDAS"
export CAPSULE_MAIN_PROCESS_URL="http://localhost:8080" 

capsctl \
    --cmd=start \
    --name=hello \
    --revision=blue \
    --wasm=./functions/hello-blue/hello-blue.wasm
```

**Now you can use this URL `http://localhost:8080/functions/hello/blue` to call the hello blue function**

### Call the hello blue function

```bash
curl -X POST http://localhost:8080/functions/hello/blue \
-H 'Content-Type: text/plain; charset=utf-8' \
-d "Bob Morane"
```

## Drop: stop and remove a running function

```bash
export CAPSULE_FAAS_TOKEN="ILOVEPANDAS"
export CAPSULE_MAIN_PROCESS_URL="http://localhost:8080"

capsctl \
    --cmd=drop \
    --name=hello \
    --revision=blue
```

## Duplicate a running Capsule HTTP server process with a new revision name

```bash
export CAPSULE_FAAS_TOKEN="ILOVEPANDAS"
export CAPSULE_MAIN_PROCESS_URL="http://localhost:8080" 

capsctl \
    --cmd=duplicate \
    --name=hello \
    --revision=green \
    --newRevision=saved_green
```
> It remains the same process (same `PID`), but you can access the function with this URL: `http://localhost:8080/functions/hello/saved_green` and this URL: `http://localhost:8080/functions/hello/green`

### Call the hello green function

```bash
curl -X POST http://localhost:8080/functions/hello/saved_green \
-H 'Content-Type: text/plain; charset=utf-8' \
-d "Bob Morane"

curl -X POST http://localhost:8080/functions/hello/green \
-H 'Content-Type: text/plain; charset=utf-8' \
-d "Bob Morane"
```

## Download the wasm module before starting the function

You can specify to the Capsule HTTP process with the `--url` option, where to download the wasm file and where to save it before starting with the `--wasm` option:

```bash
export CAPSULE_FAAS_TOKEN="ILOVEPANDAS"
export CAPSULE_MAIN_PROCESS_URL="http://localhost:8080"

capsctl \
    --cmd=start \
    --name=hello \
    --revision=0.0.1 \
    --wasm= ./store/hello.0.0.1.wasm \
    --url=http://wasm.files.com/hello/0.0.1/hello.0.0.1.wasm
```

### Authentication of the downlad

If you need to provide an authentication token, you can use these options:

```bash
--authHeaderName="PRIVATE-TOKEN" \
--authHeaderValue="${GITLAB_WASM_TOKEN}" \
```

