package main

import (
	"flag"
	"fmt"
	"github.com/bots-garden/capsule/capsule-reverse-proxy/reverse-proxy"
	"github.com/bots-garden/capsule/commons"
	"os"
)

type CapsuleFlags struct {
	httpPort string
	config   string // config file for the reverse proxy
	crt      string // https (certificate)
	key      string // https (key)
	backend  string // backend for reverse proxy (and/or service discovery)

}

func main() {
	//argsWithProg := os.Args
	args := os.Args[1:]

	if args[0] == "version" {
		fmt.Println(commons.CapsuleVersion())
	} else {
		//flags
		httpPortPtr := flag.String("httpPort", "8080", "http port")

		configPtr := flag.String("config", "", "config file (reverse proxy)")
		backendPtr := flag.String("backend", "memory", "backend for reverse proxy, registration, discovery")

		crtPtr := flag.String("crt", "", "certificate")
		keyPtr := flag.String("key", "", "key")

		flag.Parse()

		flags := CapsuleFlags{
			*httpPortPtr,
			*configPtr,
			*crtPtr,
			*keyPtr,
			*backendPtr,
		}

		reverse_proxy.Serve(flags.httpPort, flags.config, flags.backend, flags.crt, flags.key)

	}

}
