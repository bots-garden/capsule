package main

import (
	helpers "github.com/bots-garden/capsule/helpers/functions"
)

// main is required for TinyGo to compile to Wasm.
func main() {
	helpers.Log("🚀 ignition...")
	hostInformation := helpers.GetHostInformation()
	helpers.Log("👋 message from the wasm module: " + hostInformation)

	helpers.Log(helpers.Ping("✊ knock knock from the wasm module"))
}

//export add
func add(x uint32, y uint32) uint32 {
	// 🖐 a wasm module cannot print something
	//fmt.Println(x,y)
	res := x + y

	helpers.Log("👋 hello world 🌍")

	return res
}

//export helloWorld
func helloWorld() (strPtrPosSize uint64) {
	strPtrPos, size := helpers.GetStringPtrPositionAndSize("👋 hello world, I'm very happy to meet you, I love what you are doing my friend")
	return helpers.PackPtrPositionAndSize(strPtrPos, size)
}

//export sayHello
func sayHello(strPtrPos, size uint32) (strPtrPosSize uint64) {
	name := helpers.GetStringParam(strPtrPos, size)
	pos, length := helpers.GetStringPtrPositionAndSize("👋 hello " + name)

	return helpers.PackPtrPositionAndSize(pos, length)
}


// TODO: how to simplify ?
// TODO: here: create an non exportable function
// TODO: helpers/function: an exportable function that call this on
