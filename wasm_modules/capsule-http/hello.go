package main

import (
	hf "github.com/bots-garden/capsule/helpers/functions"
)

// main is required.
func main() {
	hf.SetHandle(Handle)
}

/*
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json' \
  -d '{"message": "Golang 💚 wasm"}'
*/

func Handle(param string) string {
	hf.Log("1️⃣ parameter is: " + param)
	ret := hf.Http("http://google.com", "GET", []string{"one", "two"}, "tada")
	return ret
}

// ? HandleJson, Handle<>, ...
