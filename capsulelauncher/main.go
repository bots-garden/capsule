package main

import (
	"flag"
	capsulehttp_next "github.com/bots-garden/capsule/capsulelauncher/services/http"
	reverse_proxy "github.com/bots-garden/capsule/capsulelauncher/services/reverse-proxy"
	"log"
	"os"

	capsulecli "github.com/bots-garden/capsule/capsulelauncher/services/cli"
	"github.com/go-resty/resty/v2"
)

type CapsuleFlags struct {
	mode     string
	param    string
	wasm     string
	httpPort string
	url      string
	config   string
	crt      string
	key      string
}

func main() {

	//flags
	/*
		go run main.go \
			-wasm=./wasm_modules/capsulelauncher-function-template/hello.wasm \
			-mode=cli \
			-param="👋 hello world 🌍"

		go run main.go \
			-wasm=./wasm_modules/capsulelauncher-function-template/hello.wasm \
			-mode=http \
			-httpPort=7070
	*/
	capsuleModePtr := flag.String("mode", "http", "default mode is http else: cli")
	stringParamPtr := flag.String("param", "hello world", "string parameter for the cli mode")
	wasmFilePathPtr := flag.String("wasm", "", "wasm module file path")
	httpPortPtr := flag.String("httpPort", "8080", "http port")

	wasmFileUrlPtr := flag.String("url", "", "url for downloading wasm module file")
	configPtr := flag.String("config", "", "config file (reverse proxy)")

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
	}
	//fmt.Println(flags)

	//argsWithProg := os.Args
	//args := os.Args[1:]

	//wasmModuleFilePath := flags.wasm

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

	//var envVariables = make(map[string]string)
	// https://gobyexample.com/environment-variables
	// just provide the reading access with: os.Getenv("")
	// provide some other env var for httpport wasmurl?

	switch what := flags.mode; what {
	case "http":
		capsulehttp_next.Serve(flags.httpPort, getWasmFile(), flags.crt, flags.key)
	case "cli":
		/*
			for idx, arg := range flag.Args() {
				fmt.Println(idx, "==>", arg)
			}
		*/
		capsulecli.Execute(flag.Args(), getWasmFile())
	case "reverse-proxy":
		reverse_proxy.Serve(flags.httpPort, flags.config, flags.crt, flags.key)
	default:
		log.Panicln("🔴 bad mode", *capsuleModePtr)
	}

}
