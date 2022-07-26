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
	ret, err := hf.Http("http://google.com", "GET", []string{"one", "two"}, "tada")
  if err != nil {
    hf.Log(err.Error())
  }

  _, err2 := hf.Http("http://google.com", "XXX", []string{"one", "two"}, "tada")
  if err2 != nil {
    hf.Log("this is an error: " + err2.Error())
  }


	return ret
}

// ? HandleJson, Handle<>, ...
