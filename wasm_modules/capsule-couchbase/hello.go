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
    hf.Log("👋")
    res, err := hf.CouchBaseQuery("SELECT * FROM \\`wasm-data\\`.data.docs")

    if err != nil {
        hf.Log(err.Error())
    } else {
        hf.Log("" + res)
    }

    return res, nil

}
