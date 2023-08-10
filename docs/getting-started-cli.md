# 🚀 Getting Started

## Use the Capsule CLI

First, download the last version of the Capsule CLI for the appropriate OS & ARCH (and release version):

```bash
VERSION="v0.4.2" OS="linux" ARCH="arm64"
wget -O capsule https://github.com/bots-garden/capsule/releases/download/${VERSION}/capsule-${VERSION}-${OS}-${ARCH}
chmod +x capsule
```

## Run a WASM Capsule module

To run a WASM Capsule module you need to set 2 flags:

- `--wasm`: the path to the WASM module
- `--params`: the parameter to pass to the WASM module

```bash
./capsule \
--wasm=./functions/hello/hello.wasm \
--params="Hello World"
```

You can remotely download  the WASM module with the `--url` flag:
```bash
./capsule \
--url=http://localhost:5000/hello-world.wasm \
--wasm=./tmp/hello-world.wasm 
```

## Develop a WASM Capsule module

Have a look to these samples:

- [Capsule MDK documentation: first CLI module](https://bots-garden.github.io/capsule-module-sdk/first-cli-module/)
- [capsule-cli/functions](https://github.com/bots-garden/capsule/tree/main/capsule-cli/functions)
- [Samples of the Capsule MDK](https://github.com/bots-garden/capsule-module-sdk/tree/main/samples)

