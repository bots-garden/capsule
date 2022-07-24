package main

import (
	helpers "github.com/bots-garden/capsule/helpers/functions"
)

// main is required for TinyGo to compile to Wasm.
func main() {}

/*
add 2 uint32, return 1 uint32 
*/
//export add
func add(x uint32, y uint32) uint32 {
	// 🖐 a wasm module cannot print something
	//fmt.Println(x,y)
	res := x + y

	helpers.Log("hello world")

	return res
}
