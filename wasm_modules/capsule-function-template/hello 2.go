package main

import (
	helpers "github.com/bots-garden/capsule/helpers/functions"
)

// main is required.
func main() {

	/*
		helpers.Log("🚀 ignition...")
		hostInformation := helpers.GetHostInformation()
		helpers.Log("👋 message from the wasm module: " + hostInformation)

		helpers.Log(helpers.Ping("✊ knock knock from the wasm module"))
	*/

	helpers.SetHandle(Handle)
}

/*
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json' \
  -d '{"message": "Golang 💚 wasm"}'
*/

func Handle(param string) string {
	helpers.Log("1️⃣ parameter is: " + param)
	ret := "👋 you sent me this: " + param
	return ret
}

// ? HandleJson, Handle<>, ...
