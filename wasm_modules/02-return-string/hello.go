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

//export helloWorld
func helloWorld() uint64 {
	ptr, size := helpers.StringToPtr("👋 hello world, I'm very happy to meet you, I love what you are doing my friend")
	return helpers.PtrSizePair(ptr, size)
}

// how to simplify ?
