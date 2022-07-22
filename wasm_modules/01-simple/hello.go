package main

import (
	"github.com/bots-garden/capsule/helpers/functions"
)

// main is required for TinyGo to compile to Wasm.
func main() {}

//export add
func add(x uint32, y uint32) uint32 {
	// 🖐 a wasm module cannot print something
	//fmt.Println(x,y)
	res := x + y

	helpers.Log("hello world")

	return res
}
