package main

import (
    "errors"
    hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
    "strings"
)

func main() {
    hf.SetHandle(Handle)
}

func Handle(params []string) (string, error) {
    var errs []string

    _, err1stMsg := hf.NatsPublish("ping", "🖐 Hello from WASM with Nats 💜")
    _, err2ndMsg := hf.NatsPublish("ping", "👋 Hello World 🌍")

    if err1stMsg != nil {
        errs = append(errs, err1stMsg.Error())
    }
    if err2ndMsg != nil {
        errs = append(errs, err2ndMsg.Error())
    }

    return "NATS Rocks!", errors.New(strings.Join(errs, "|"))
}
