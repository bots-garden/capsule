package main

import (
	"flag"
	"fmt"
	"github.com/bots-garden/capsule/capsule-registry/registry"
	"github.com/bots-garden/capsule/commons"
	"os"
)

type CapsuleFlags struct {
	httpPort string
	crt      string // https (certificate)
	key      string // https (key)
	files    string // root location of the wasm modules (for the registry)
}

func main() {
	//argsWithProg := os.Args
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("😮 no args. Type capsule-registry --help")
		os.Exit(0)
	}

	if args[0] == "version" {
		fmt.Println(commons.CapsuleVersion())
	} else {
		//flags
		httpPortPtr := flag.String("httpPort", "8080", "http port")
		filesPtr := flag.String("files", "", "root location of the wasm modules (for the registry)")

		crtPtr := flag.String("crt", "", "certificate")
		keyPtr := flag.String("key", "", "key")

		flag.Parse()

		flags := CapsuleFlags{
			*httpPortPtr,
			*crtPtr,
			*keyPtr,
			*filesPtr,
		}

		registry.Serve(flags.httpPort, flags.files, flags.crt, flags.key)

	}

}
