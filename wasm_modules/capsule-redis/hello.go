package main

// TinyGo wasm module
import (
    hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

// main is required.
func main() {

    hf.Log("🚀 ignition...")
    hf.SetHandle(Handle)
}

func Handle(params []string) (string, error) {

    // add a key, value
    res1, err := hf.RedisSet("001", "Hello World")
    if err != nil {
        hf.Log(err.Error())
    } else {
        hf.Log("" + res1)
    }

    // read the value
    res2, err := hf.RedisGet("001")
    if err != nil {
        hf.Log(err.Error())
    } else {
        hf.Log("🎉 value: " + res2)
    }

    return res2, nil

}

// ? HandleJson, Handle<>, ...
