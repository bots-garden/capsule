# Capsule Registry


## Use the Capsule Wasm modules registry
> 🚧 this is a work in progress

It's possible to download the wasm module from a remote location before serving it:

```bash
./capsule \
   -url=http://localhost:9090/hello.wasm \
   -wasm=./tmp/hello.wasm \
   -mode=http \
   -httpPort=8080
```

The **Capsule Registry** allows to upload and serve the wasm modules

### Start the wasm registry

```bash
./capsule-registry \
   -files="./wasm_modules" \
   -httpPort=4999
```
> The `-files` tag defines where the modules are uploaded

### Upload a wasm module

```bash
curl -X POST http://localhost:4999/upload/k33g/hola/0.0.0 \
  -F "file=@../with-proxy/capsule-hola/hola.wasm" \
  -F "info=hola function from @k33g" \
  -H "Content-Type: multipart/form-data"
```
> - The upload url is defined like this: `/upload/user_name_or_organization/module_name/tag`
> - The wasm module will be saved to `./wasm_modules/user_name_or_organization/module_name/tag/module_name.wasm`
> - This data `"info=hola function from @k33g"` will create a file `./wasm_modules/user_name_or_organization/module_name/tag/module_name.info`

Then you can download the module at http://localhost:4999/upload/k33g/hola/0.0.0/hello.wasm

### Download and start a wasm module

```bash
./capsule \
   -wasm=./tmp/hola.wasm \
   -url="http://localhost:4999/k33g/hola/0.0.0/hola.wasm" \
   -mode=http \
   -httpPort=7072
```

### Get information about a wasm module

```bash
curl http://localhost:4999/info/k33g/hola/0.0.0
```

### Get the list of all the wasm modules

```bash
curl http://localhost:4999/modules
```
