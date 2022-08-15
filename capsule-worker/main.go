package main

import (
	"flag"
	"fmt"
	"github.com/bots-garden/capsule/capsule-worker/worker"
	"github.com/bots-garden/capsule/commons"
	"os"
)

type CapsuleFlags struct {
	httpPort        string
	crt             string // https (certificate)
	key             string // https (key)
	backend         string // backend for reverse proxy (and/or service discovery)
	reverseProxy    string // url to the reverse proxy
	workerDomain    string // domain or ip address
	capsulePath     string // where is the capsule executable
	httpPortCounter int    // to attribute a http port to a running module
}

func main() {
	//argsWithProg := os.Args
	args := os.Args[1:]

	if args[0] == "version" {
		fmt.Println(commons.CapsuleVersion())
	} else {
		//flags
		httpPortPtr := flag.String("httpPort", "8080", "http port")

		backendPtr := flag.String("backend", "memory", "backend for reverse proxy, registration, discovery")

		reverseProxyPtr := flag.String("reverseProxy", "", "url of the reverse proxy")

		workerDomainPtr := flag.String("workerDomain", "localhost", "domain or ip address of the worker")

		capsulePathPtr := flag.String("capsulePath", "capsule", "path to capsule (it could be a cmd)")

		httpPortCounterPtr := flag.Int("httpPortCounter", 10000, "httpPortCounter is used for the wasm module server (incremented at every new function deployment)")

		crtPtr := flag.String("crt", "", "certificate")
		keyPtr := flag.String("key", "", "key")

		flag.Parse()

		flags := CapsuleFlags{
			*httpPortPtr,
			*crtPtr,
			*keyPtr,
			*backendPtr,
			*reverseProxyPtr,
			*workerDomainPtr,
			*capsulePathPtr,
			*httpPortCounterPtr,
		}

		worker.Serve(flags.httpPort, flags.capsulePath, flags.httpPortCounter, flags.reverseProxy, flags.workerDomain, flags.backend, flags.crt, flags.key)

	}

}
