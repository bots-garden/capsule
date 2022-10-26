# Developer Guide

## First CLI function

Create a `go.mod` file: (`go mod init cli-say-hello`)
```
module cli-say-hello

go 1.18
```

Install the Capsule dependencies:
```bash
go get github.com/bots-garden/capsule/capsulemodule/hostfunctions
```

Create a `hello.go` file:
```go
package main

import hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"

// main is required.
func main() {
	hf.SetHandle(Handle)
}

func Handle(params []string) (string, error) {
	var err error
	for _, param := range params {
		hf.Log("- parameter is: " + param)
	}

    ret := "The first parameter is: " + params[0]

    return ret, err // err = nil
}
```
> - `hf.SetHandle(Handle)` defines the called wasm function
> - `hf.Log(string)` prints a value

Build the wasm module:
```bash
tinygo build -o hello.wasm -scheduler=none --no-debug -target wasi ./hello.go
```

Execute the `Handle` function:
```bash
./capsule \
   -wasm=./hello.wasm \
   -mode=cli \
   "👋 hello world 🌍🎃" 1234 "Bob Morane"
```
> - `-wasm` flag: the path to the wasm file
> - `-mode` execution mode


*output:*
```bash
- parameter is: 👋 hello world 🌍🎃
- parameter is: 1234
- parameter is: Bob Morane
The first parameter is: 👋 hello world 🌍🎃
```
