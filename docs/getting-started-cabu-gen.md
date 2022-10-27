# 🚀 Getting Started

## Generate a new project function

**Cabu** can generate function project from templates:

```bash
# template name: `service-get`
# function project name `hello-world`
cabu generate service-get hello-world
```
> At the first launch, **Cabu** will pull a docker image with all the necessary resources to build the WASM function.
```bash
🐳 using k33g/capsule-builder:0.0.2
Unable to find image 'k33g/capsule-builder:0.0.2' locally
0.0.2: Pulling from k33g/capsule-builder
68c15fb212c3: Pull complete
28b965d0936e: Pull complete
f7ba6ae51b0b: Pull complete
Digest: sha256:47ebf274d7c378d1795f6c8a78d71c45e8368b33a7a3ba8e48ef131a08fd9ac4
Status: Downloaded newer image for k33g/capsule-builder:0.0.2
✅🙂 hello-world function generated
```

**Cabu** has generated the `hello-world` project:

```bash
.
├── hello-world
│  ├── go.mod
│  └── hello-world.go
```

With the following source code:
```golang
package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

func main() {
	hf.SetHandleHttp(Handle)
}

func Handle(req hf.Request) (resp hf.Response, errResp error) {

	headers := map[string]string{
		"Content-Type": "text/html; charset=utf-8",
	}
	resp = hf.Response{
		Body: "<h1>👋 hello world 🌍</h1>",
		Headers: headers,
	}

	return resp , nil
}
```
