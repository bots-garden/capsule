// Package main
package main

import (
	capsule "github.com/bots-garden/capsule-module-sdk"
)

func main() {
	capsule.SetHandle(Handle)
}

// Handle function
func Handle(params []byte) ([]byte, error) {

	capsule.RedisSet("one", []byte("👋"))
	capsule.RedisSet("two", []byte("hello"))
	capsule.RedisSet("three", []byte("world"))
	capsule.RedisSet("four", []byte("🌍"))

	word1, _ := capsule.RedisGet("one")
	word2, _ := capsule.RedisGet("two")
	word3, _ := capsule.RedisGet("three")
	word4, _ := capsule.RedisGet("four")

	capsule.Print("📝: " + string(word1) + " " + string(word2) + " " + string(word3) + " " + string(word4))
	
	capsule.CacheDel("three")

	// TODO: add other filters
	keys, err := capsule.RedisKeys("*")
	if err != nil {
		capsule.Print("🔴 " + err.Error())
	}
	for _, key := range keys {
		capsule.Print("🔑 " + key)
	}

	return []byte("👋 Hello"), nil

}
