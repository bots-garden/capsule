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
    _, _ = hf.RedisSet("bob1", "Bob One")
    _, _ = hf.RedisSet("jane", "Jane Doe")
    _, _ = hf.RedisSet("bob2", "Bob Two")
    _, _ = hf.RedisSet("bob3", "Bob Three")
    _, _ = hf.RedisSet("john", "John Doe")
    _, _ = hf.RedisSet("bob4", "Bob Four")
    _, _ = hf.RedisSet("bob5", "Bob Five")

    legion, err := hf.RedisKeys("bob*")
    if err != nil {
        hf.Log(err.Error())
    }

    for _, bob := range legion {
        hf.Log(bob)
    }

    return "we are legion", err

}
