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

	capsule.CacheSet("one", []byte("👋"))
	capsule.CacheSet("two", []byte("hello"))
	capsule.CacheSet("three", []byte("world"))
	capsule.CacheSet("four", []byte("🌍"))

	word1, _ := capsule.CacheGet("one")
	word2, _ := capsule.CacheGet("two")
	word3, _ := capsule.CacheGet("three")
	word4, _ := capsule.CacheGet("four")

	capsule.Print("📝: " + string(word1) + " " + string(word2) + " " + string(word3) + " " + string(word4))
	
	capsule.CacheDel("three")

	// TODO: add other filters
	keys, err := capsule.CacheKeys("*")
	if err != nil {
		capsule.Print("🔴 " + err.Error())
	}
	for _, key := range keys {
		capsule.Print("🔑 " + key)
	}

	return []byte("👋 Hello"), nil

}
