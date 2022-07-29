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
    //query := "INSERT INTO `wasm-data`.data.docs (KEY, VALUE) VALUES (\"key100\", { \"type\" : \"info\", \"name\" : \"this is an info\" });"
    query := "SELECT * FROM `wasm-data`.data.docs"
    res, err := hf.CouchBaseQuery(query)

    if err != nil {
        hf.Log(err.Error())
    } else {
        hf.Log("" + res)
    }

    return res, nil

}
