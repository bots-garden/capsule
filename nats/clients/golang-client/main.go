package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
	nc, err := nats.Connect("nats.devsecops.fun:4222")
	// nats.DefaultURL
	if err != nil {
		fmt.Println(err.Error())
	}
	defer nc.Close()

	err = nc.Publish("ping", []byte("😍 Hello World"))

	err = nc.Publish("ping", []byte("😁😁 Hello World"))

	if err != nil {
		fmt.Println(err.Error())
	}

}
