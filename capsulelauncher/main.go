package main

import (
	"flag"
	"fmt"
	capsule_cli "github.com/bots-garden/capsule/capsulelauncher/services/cli"
	capsule_http "github.com/bots-garden/capsule/capsulelauncher/services/http"
	"github.com/bots-garden/capsule/capsulelauncher/services/worker"
	"github.com/bots-garden/capsule/commons"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
)

type CapsuleFlags struct {
	mode            string // cli, http, reverse-proxy, registry
	param           string
	wasm            string // wasm file location
	httpPort        string
	url             string // to download the wasm file
	config          string // config file for the reverse proxy
	crt             string // https (certificate)
	key             string // https (key)
	backend         string // backend for reverse proxy (and/or service discovery)
	files           string // root location of the wasm modules (for the registry)
	registry        string // url to the registry
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
		capsuleModePtr := flag.String("mode", "http", "default mode is http else: cli")
		stringParamPtr := flag.String("param", "hello world", "string parameter for the cli mode")
		wasmFilePathPtr := flag.String("wasm", "", "wasm module file path")
		httpPortPtr := flag.String("httpPort", "8080", "http port")

		wasmFileUrlPtr := flag.String("url", "", "url for downloading wasm module file")
		configPtr := flag.String("config", "", "config file (reverse proxy)")
		backendPtr := flag.String("backend", "memory", "backend for reverse proxy, registration, discovery")

		filesPtr := flag.String("files", "", "root location of the wasm modules (for the registry)")

		registryPtr := flag.String("registry", "", "url of the wasm registry")
		reverseProxyPtr := flag.String("reverseProxy", "", "url of the reverse proxy")

		workerDomainPtr := flag.String("workerDomain", "localhost", "domain or ip address of the worker")

		capsulePathPtr := flag.String("capsulePath", "capsule", "path to capsule (it could be a cmd)")

		httpPortCounterPtr := flag.Int("httpPortCounter", 10000, "httpPortCounter is used for the wasm module server (incremented at every new function deployment)")

		crtPtr := flag.String("crt", "", "certificate")
		keyPtr := flag.String("key", "", "key")

		flag.Parse()

		flags := CapsuleFlags{
			*capsuleModePtr,
			*stringParamPtr,
			*wasmFilePathPtr,
			*httpPortPtr,
			*wasmFileUrlPtr,
			*configPtr,
			*crtPtr,
			*keyPtr,
			*backendPtr,
			*filesPtr,
			*registryPtr,
			*reverseProxyPtr,
			*workerDomainPtr,
			*capsulePathPtr,
			*httpPortCounterPtr,
		}

		getWasmFile := func() []byte {
			var wasmFile []byte
			// 📂 Load from file and then Instantiate a WebAssembly module
			loadWasmFile := func(path string) []byte {
				wasmFileToLoad, errLoadWasmFile := os.ReadFile(path)

				if errLoadWasmFile != nil {
					log.Panicln("🔴 Error while loading the wasm file:", errLoadWasmFile)
				}
				return wasmFileToLoad
			}
			//TODO; add authentication with headers
			if len(flags.url) == 0 {
				wasmFile = loadWasmFile(flags.wasm)
			} else {
				client := resty.New()
				_, errLoadWasmFileFromUrl := client.R().
					SetOutput(flags.wasm).
					Get(flags.url)

				if errLoadWasmFileFromUrl != nil {
					log.Panicln("🔴 Error while downloading the wasm file:", errLoadWasmFileFromUrl)
				}
				wasmFile = loadWasmFile(flags.wasm)
			}
			return wasmFile
		}

		switch what := flags.mode; what {
		case "http":
			capsule_http.Serve(flags.httpPort, getWasmFile(), flags.crt, flags.key)
		case "cli":
			capsule_cli.Execute(flag.Args(), getWasmFile())
		case "worker":
			worker.Serve(flags.httpPort, flags.capsulePath, flags.httpPortCounter, flags.reverseProxy, flags.workerDomain, flags.backend, flags.crt, flags.key)

		default:
			log.Panicln("🔴 bad mode", *capsuleModePtr)
		}
	}

}
