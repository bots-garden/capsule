package main

import (
	//"errors"
	hf "github.com/bots-garden/capsule/wasmhostfunctions"
)

// main is required.
func main() {

	hf.Log("🚀 ignition...")
	hostInformation := hf.GetHostInformation()
	hf.Log("👋 message from the wasm module: " + hostInformation)
	hf.Log(hf.Ping("✊ knock knock from the wasm module"))

	hf.SetHandle(Handle)
}

/*
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json' \
  -d '{"message": "Golang 💚 wasm"}'
*/

func Handle(param string) (string, error) {
	hf.Log("1️⃣ parameter is: " + param)

    txt, err := hf.ReadFile("about.txt")
    if err != nil {
        hf.Log(err.Error())
    }

    hf.Log(txt)


	ret := "👋 you sent me this: " + param
	//return ret, errors.New("😡 ouch")
    return ret, nil
}

// ? HandleJson, Handle<>, ...
