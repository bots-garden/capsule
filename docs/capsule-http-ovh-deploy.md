# Deploy Capsule FaaS on OVH Cloud

For this recipe, I'm using an OVH Cloud compute instance (sponsored by [@titimoby](https://twitter.com/titimoby) DevRel at OVH) with the following specifications:

- Model: D2-4
- RAM: 4 GB
- Processor: 2 vCores

## SSH Connect to the OVH Cloud compute instance

```bash
ssh -i ./keys/capsule ubuntu@${OVH_IP}
```
> Remark: I'm using a ssh key to connect to the OVH Cloud instance and `OVH_IP` is the IP address of the OVH Cloud instance.

## Install the last version of Capsule HTTP on the OVH Cloud instance

```bash
VERSION="v0.4.1" OS="linux" ARCH="amd64"
wget -O capsule-http https://github.com/bots-garden/capsule/releases/download/${VERSION}/capsule-http-${VERSION}-${OS}-${ARCH}

chmod +x capsule-http
sudo cp capsule-http  /usr/local/bin/capsule-http
rm capsule-http
capsule-http --version
```

If you already have Capsule HTTP installed and running, you need to stop it:
```bash
# from your computer
pkill capsule-http
```

## Start Capsule HTTP on the OVH Cloud instance

```bash
export CAPSULE_FAAS_TOKEN="ILOVEPANDAS"
nohup capsule-http --httpPort=8888 \
--faas=true \
--crt=faas.capsule.foundation.crt \
--key=faas.capsule.foundation.key &> /dev/null &
```

> - `CAPSULE_FAAS_TOKEN` allows to protect the admin routes
> - I have a domain name linked to the instance `faas.capsule.foundation`
> - `--crt=faas.capsule.foundation.crt` and `--key=faas.capsule.foundation.key` allow to use HTTPS.
> - `--httpPort=8888`, the port of the Capsule HTTP service.

To be able to connect on https://faas.capsule.foundation/, you need to do the following:
```bash
sudo iptables -t nat -A PREROUTING -p tcp --dport 443 -j REDIRECT --to-port 8888
sudo apt-get install iptables-persistent
```

### Check the status of Capsule HTTP

If you open this URL https://faas.capsule.foundation/ on your browser, you should see the following message:

```bash
Capsule [HTTP] v0.4.1 🫑 [pepper][faas]
```

### Install the last version of CapsCtl

**CapsCtl** is a CLI to send commands to the Capsule HTTP server when it is running in **FaaS** mode.

```bash
VERSION="v0.4.1" OS="linux" ARCH="arm64"
wget -O capsctl https://github.com/bots-garden/capsule/releases/download/${VERSION}/capsctl-${VERSION}-${OS}-${ARCH}
chmod +x capsctl
sudo cp capsctl  /usr/local/bin/capsctl
rm capsctl
capsctl --version
```

## Add a start page to Capsule HTTP

We need to add a start page to Capsule HTTP by creating a new WASM function (module).

### `index.html`

Create a file named `index.html` in a directory:

```html
<html>
  <head>
    <meta charset="utf-8">
    <title>Wasm is fantastic 😍</title>
      <meta name="viewport" content="width=device-width, initial-scale=1">
      <style>
        .container { min-height: 100vh; display: flex; justify-content: center; align-items: center; text-align: center; }
        .title { font-family: "Source Sans Pro", "Helvetica Neue", Arial, sans-serif; display: block; font-weight: 300; font-size: 80px; color: #35495e; letter-spacing: 1px; }
        .subtitle { font-family: "Source Sans Pro", "Helvetica Neue", Arial, sans-serif; font-weight: 300; font-size: 32px; color: #526488; word-spacing: 5px; padding-bottom: 15px; }
        .links { padding-top: 15px; }
      </style>
  </head>

  <body>
    <section class="container">
      <div>
        <h1 class="title">👋 Hello World 🌍</h1>
        <h2 class="subtitle">Served with 💜 by Capsule 💊 [HTTP] v0.4.1 🫑 [pepper] </h2>
        <h2 class="subtitle">🎉 Hosted on OVH Cloud [🚀 Faas mode]</h2>
        <h2 class="subtitle">🥰 With the help of @titimoby</h2>
      </div>
    </section>
  </body>

</html>
```

### `main.go`

Into the same directory, create a file named `main.go` and add the following code:

```bash
go mod init index.page
touch main.go
```

> main.go
```golang
// Package main => serving an html resource
package main

import (
	_ "embed"
	capsule "github.com/bots-garden/capsule-module-sdk"
)

var (
	//go:embed index.html
	html []byte
)

func main() {

	capsule.SetHandleHTTP(func (param capsule.HTTPRequest) (capsule.HTTPResponse, error) {
		
		return capsule.HTTPResponse{
			TextBody: string(html),
			Headers: `{
				"Content-Type": "text/html; charset=utf-8",
				"Cache-Control": "no-cache",
				"X-Powered-By": "capsule-module-sdk"
			}`,
			StatusCode: 200,
		}, nil
	})
}
```

### Build the WASM module

Use the following command to build the WASM module:
```bash
go mod tidy
tinygo build -o index.page.wasm \
    -scheduler=none \
    --no-debug \
    -target wasi ./main.go 
```

### Deploy the WASM module on the OVH Cloud instance

From your computer, run the following command:
```bash
scp -i ./keys/capsule ./index-page/index.page.wasm ubuntu@${OVH_IP}:./
```

### Start the `index.page` function

```bash
export CAPSULE_FAAS_TOKEN="ILOVEPANDAS"
export CAPSULE_MAIN_PROCESS_URL="https://faas.capsule.foundation"

capsctl \
    --cmd=start \
    --name=index.page \
    --revision=default \
    --wasm=./index.page.wasm
```

You should get this output:
```bash
✅ index.page/default is started
ℹ️ url: https://faas.capsule.foundation/functions/index.page/default
```

Now you can call the function:
```bash
curl -X GET "https://faas.capsule.foundation/functions/index.page/default"
```

### `index.page` function name is mapped to the `/` root

If you open this URL: [https://faas.capsule.foundation](https://faas.capsule.foundation) in your browser, you will reach the `index.page` function.

It's a specific case of Capsule FaaS, and it's only possible because you name the function `index.page` when you started it with this command:

```bash
capsctl \
    --cmd=start \
    --name=index.page \
    --revision=default \
    --wasm=./index.page.wasm
```

## Create and Deploy a new function

### `main.go`

Into the a directory, create a file named `main.go` and add the following code:

```bash
go mod init hello
touch main.go
```

```golang
package main

import (
	capsule "github.com/bots-garden/capsule-module-sdk"
)

func main() {
	capsule.SetHandleHTTP(func (param capsule.HTTPRequest) (capsule.HTTPResponse, error) {

		return capsule.HTTPResponse{
			TextBody: "🍊 👋 Hey " + param.Body +" !",
			Headers: `{"Content-Type": "text/plain; charset=utf-8"}`,
			StatusCode: 200,
		}, nil
		
	})
}
```

### Build the WASM module

```bash
tinygo build -o hello.wasm \
    -scheduler=none \
    --no-debug \
    -target wasi ./main.go 
``` 

### Deploy the WASM module on the OVH Cloud instance

```bash
scp -i ./keys/capsule ./hello/hello.wasm ubuntu@${OVH_IP}:./
``` 

### Start the `hello` function with a revision named `orange`

```bash
capsctl \
    --cmd=start \
    --name=hello \
    --revision=orange \
    --wasm=./hello.wasm
``` 

You should get this output:
```bash
✅ hello/orange is started
ℹ️ url: https://faas.capsule.foundation/functions/hello/orange
``` 

Now you can call the function:
```bash
curl -X POST ${CAPSULE_MAIN_PROCESS_URL}/functions/hello/orange \
-H 'Content-Type: text/plain; charset=utf-8' \
-d 'Bob Morane'
```

You should get this output:
```
🍊 👋 Hey Bob Morane !
```

