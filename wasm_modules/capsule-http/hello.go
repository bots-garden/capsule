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

//TODO: handle error return
func Handle(param string) string {

	hf.Log("1️⃣ parameter is: " + param)

	headers := map[string]string{"Accept": "application/json", "Content-Type": " text/html; charset=UTF-8"}

	ret, err := hf.Http("http://google.com", "GET", headers, "tada")
	if err != nil {
		hf.Log(err.Error())
	}
	// SetHeader("Accept", "application/json")

	_, err2 := hf.Http("http://google.com", "XXX", headers, "tada")
	if err2 != nil {
		hf.Log("this is an error: " + err2.Error())
	}

	return ret
}

// ? HandleJson, Handle<>, ...
