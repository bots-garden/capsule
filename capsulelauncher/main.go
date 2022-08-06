package main

import (
	"flag"
	reverse_proxy "github.com/bots-garden/capsule/capsulelauncher/services/reverse-proxy"
	"log"
	"os"

	capsulecli "github.com/bots-garden/capsule/capsulelauncher/services/cli"
	capsulehttp "github.com/bots-garden/capsule/capsulelauncher/services/http"
	"github.com/go-resty/resty/v2"
)

type CapsuleFlags struct {
	mode     string
	param    string
	wasm     string
	httpPort string
	url      string
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

	flag.Parse()

	flags := CapsuleFlags{
		*capsuleModePtr,
		*stringParamPtr,
		*wasmFilePathPtr,
		*httpPortPtr,
		*wasmFileUrlPtr,
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
		capsulehttp.Serve(flags.httpPort, getWasmFile())
	case "cli":
		/*
			for idx, arg := range flag.Args() {
				fmt.Println(idx, "==>", arg)
			}
		*/
		capsulecli.Execute(flag.Args(), getWasmFile())
	case "reverse-proxy":
		reverse_proxy.Serve(flags.httpPort)
	default:
		log.Panicln("🔴 bad mode", *capsuleModePtr)
	}

}
