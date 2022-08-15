package main

import hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"

// main is required.
func main() {

	hf.Log("🚀 ignition...")

	hf.Log(hf.Ping("✊ knock knock from the wasm module"))
	hf.Log(hf.Ping("✊ knock knock from the wasm module"))

	hf.SetHandle(Handle)
}

func Handle(params []string) (string, error) {

	message, err := hf.GetEnv("MESSAGE")
	if err != nil {
		hf.Log(err.Error())
	} else {
		hf.Log("MESSAGE=" + message)
	}

	for _, param := range params {
		hf.Log("- parameter is: " + param)
	}

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

	ret := "👋 you sent me this: " + params[0]
	//return ret, errors.New("😡 ouch")
	return ret, nil
}

// ? HandleJson, Handle<>, ...
