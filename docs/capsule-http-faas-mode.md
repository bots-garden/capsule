# 🚀 Capsule HTTP FaaS mode

> - Released since Capsule HTTP `v0.3.7 🥦 [broccoli]`
> - This is work in progress 🚧

A Capsule HTTP server can start/spawn other Capsule HTTP server processes.

## Requirements

### Install the last version of Capsule HTTP

```bash
VERSION="v0.3.8" OS="linux" ARCH="arm64"
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
VERSION="v0.3.8" OS="linux" ARCH="arm64"
wget -O capsctl https://github.com/bots-garden/capsule/releases/download/${VERSION}/capsctl-${VERSION}-${OS}-${ARCH}
chmod +x capsctl
sudo cp capsctl  /usr/local/bin/capsctl
rm capsctl
capsctl --version
```

## Start Capsule HTTP FaaS mode

```bash
CAPSULE_DOMAIN="http://localhost" \
CAPSULE_FAAS_TOKEN="ILOVEPANDAS" \
capsule-http \
--wasm=./functions/index-page/index-page.wasm \
--httpPort=8080 \
--faas=true
```

You should get an output like this:
```
2023/05/29 15:12:18 🚀 faas mode activated!
2023/05/29 15:12:18 📦 wasm module loaded: ./functions/index-page/index-page.wasm
2023/05/29 15:12:18 💊 Capsule [HTTP] v0.3.8 🥬 [leafy greens]
 http server is listening on: 8080 🌍
```

> In a future version, the wasm file won't be mandatory anymore.

## Launch another Capsule HTTP server

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
    --path="/usr/local/bin/capsule-http" \
    --wasm=./functions/hello-green/hello-green.wasm
```
> - `--stopAfter=10` this will stop the Capsule HTTP server process after 10 seconds
> - `--stopAfter` is not mandatory (then the Capsule HTTP server process will never stop)
> - if the process is stopped, the Capsule HTTP server will be restarted at every call
> - `--path` means you can use various version of Capsule HTTP

**Now you can use this URL `http://localhost:8080/functions/hello/green` to call the hello green function**

### Call the hello green function

```bash
curl -X POST http://localhost:8080/functions/hello/green \
-H 'Content-Type: text/plain; charset=utf-8' \
-d "Bob Morane"
```

## Launch another Capsule HTTP server process

```bash
export CAPSULE_FAAS_TOKEN="ILOVEPANDAS"
export CAPSULE_INSTALL_PATH="/usr/local/bin/capsule-http"

capsctl \
    --cmd=start \
    --stopAfter=10 \
    --name=hello \
    --revision=blue \
    --description="this the hello module, blue revision" \
    --env='["MESSAGE=🔵","GREETING=🎉"]'\
    --path="/usr/local/bin/capsule-http" \
    --wasm=./functions/hello-blue/hello-blue.wasm
```

**Now you can use this URL `http://localhost:8080/functions/hello/blue` to call the hello blue function**

### Call the hello blue function

```bash
curl -X POST http://localhost:8080/functions/hello/blue \
-H 'Content-Type: text/plain; charset=utf-8' \
-d "Bob Morane"
```

## Stop and remove a running Capsule HTTP server process

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
