package main

// TinyGo wasm module
import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

// main is required.
func main() {

	hf.Log("🚀 ignition...")
	hostInformation := hf.GetHostInformation()
	hf.Log("👋 message from the wasm module: " + hostInformation)
	hf.Log(hf.Ping("✊ knock knock from the wasm module"))

	hf.SetHandle(Handle)
}

func Handle(param string) (string, error) {

	message, err := hf.GetEnv("MESSAGE")
	if err != nil {
		hf.Log(err.Error())
	} else {
		hf.Log("MESSAGE=" + message)
	}

	hf.Log("1️⃣ parameter is: " + param)

	txt, err := hf.ReadFile("about.txt")
	if err != nil {
		hf.Log(err.Error())
	}
	hf.Log(txt)

	newFile, err := hf.WriteFile("hello.txt", "👋 HELLO WORLD 🌍")
	if err != nil {
		hf.Log(err.Error())
	}
	hf.Log(newFile)

	ret := "👋 you sent me this: " + param
	//return ret, errors.New("😡 ouch")
	return ret, nil
}

// ? HandleJson, Handle<>, ...
