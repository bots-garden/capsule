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

	capsule.Print("Environment variable → MESSAGE: " + capsule.GetEnv("MESSAGE"))

	err := capsule.WriteFile("./hello.txt", []byte("👋 Hello World! 🌍"))
	if err != nil {
		capsule.Print(err.Error())
	}

	
	data, err := capsule.ReadFile("./hello.txt")
	if err != nil {
		capsule.Print(err.Error())
	}
	capsule.Print("📝: " + string(data))
	

	return []byte("👋 Hello " + string(params)), nil

}
