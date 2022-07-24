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

func Handle(param string) string {
	helpers.Log("1️⃣ parameter is: " + param)
	ret := "👋 you sent me this: " + param
	return ret
}

// ? HandleJson, Handle<>, ...

// TODO: how to simplify ?
// TODO: here: create an non exportable function
// TODO: helpers/function: an exportable function that call this on
