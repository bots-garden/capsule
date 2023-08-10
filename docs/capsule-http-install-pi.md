# Run a FaaS on a Raspberry PI

> I did this on a Pi3A+ with the Raspberry PI OS Lite 64-bit

## Install Capsule HTTP

```bash
# connect to the PI
ssh k33g@capsulezero.local

## Install Capsule HTTP
VERSION="v0.4.2" OS="linux" ARCH="arm64"
wget -O capsule-http https://github.com/bots-garden/capsule/releases/download/${VERSION}/capsule-http-${VERSION}-${OS}-${ARCH}
chmod +x capsule-http
./capsule-http --version
```

Or you can copy the appropriate Capsule HTTP binary from your computer to the RPI:

```bash
scp capsule-http-v0.4.1-linux-arm64 k33g@capsulezero.local:./
```

## Start Capsule HTTP FaaS mode

```bash
ssh k33g@capsulezero.local -f "./capsule-http --httpPort=8080 --faas=true"
```

Try: `curl http://capsulezero.local:8080`, you should get `Capsule [HTTP] v0.4.2 ⛱️ [beach umbrella][faas]`


## Deploy some functions

```bash
# copy the functions to the RPI
cd capsule-http/tests/faas
scp -r ./functions k33g@capsulezero.local:./
```

> requirement, install **CapsCtl**: [capsule-http-faas-mode](capsule-http-faas-mode.md)

```bash
# start a function using capsctl
export CAPSULE_MAIN_PROCESS_URL="http://capsulezero.local:8080" 
capsctl \
  --cmd=start \
  --stopAfter=10 \
  --name=hey \
  --revision=default \
  --wasm=./functions/hey/hey.wasm
```


```bash
# Call the function
curl -X POST http://capsulezero.local:8080/functions/hey/default \
-H 'Content-Type: text/plain; charset=utf-8' \
-d "Bob Morane"
```

## Stop Capsule HTTP FaaS mode

```bash
ssh k33g@capsulezero.local -f "pkill capsule-http"
```
