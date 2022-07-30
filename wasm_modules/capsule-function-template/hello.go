package main

// TinyGo wasm module
import (
    hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
    hf_console "github.com/bots-garden/capsule/capsulemodule/hostfunctions/console"

)

// main is required.
func main() {

    hf_console.Log("🚀 ignition...")
    hostInformation := hf.GetHostInformation()
    hf_console.Log("👋 message from the wasm module: " + hostInformation)
    hf_console.Log(hf.Ping("✊ knock knock from the wasm module"))
    hf_console.Log(hf.Ping("✊ knock knock from the wasm module"))

    hf.SetHandle(Handle)
}

func Handle(params []string) (string, error) {

    message, err := hf.GetEnv("MESSAGE")
    if err != nil {
        hf_console.Log(err.Error())
    } else {
        hf_console.Log("MESSAGE=" + message)
    }

    for _, param := range params {
        hf_console.Log("- parameter is: " + param)
    }

    txt, err := hf.ReadFile("about.txt")
    if err != nil {
        hf_console.Log(err.Error())
    }
    hf_console.Log(txt)

    newFile, err := hf.WriteFile("hello.txt", "👋 HELLO WORLD 🌍")
    if err != nil {
        hf_console.Log(err.Error())
    }
    hf_console.Log(newFile)

    ret := "👋 you sent me this: " + params[0]
    //return ret, errors.New("😡 ouch")
    return ret, nil
}

// ? HandleJson, Handle<>, ...
