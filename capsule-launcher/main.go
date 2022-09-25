package main

import (
	"flag"
	"fmt"
	"github.com/bots-garden/capsule/capsule-launcher/services/cli"
	"github.com/bots-garden/capsule/capsule-launcher/services/http"
	capsulemqtt "github.com/bots-garden/capsule/capsule-launcher/services/mqtt"
	capsulenats "github.com/bots-garden/capsule/capsule-launcher/services/nats"
	"github.com/bots-garden/capsule/commons"
	"github.com/go-resty/resty/v2"
	"os"
)

type CapsuleFlags struct {
	mode     string // cli, http, reverse-proxy, registry
	param    string
	wasm     string // wasm file location
	httpPort string
	url      string // to download the wasm file
	crt      string // https (certificate)
	key      string // https (key)
	registry string // url to the registry
	natssrv  string // nats server
	subject  string // nats topic
	mqttsrv  string // mqtt server
	topic    string // mqtt topic
	clientId string // mqtt clientId
}

func main() {
	//argsWithProg := os.Args
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("😮 no args. Type capsule --help")
		os.Exit(0)
	}

	if args[0] == "version" {
		fmt.Println(commons.CapsuleVersion())
	} else {
		//flags
		capsuleModePtr := flag.String("mode", "http", "default mode is http else: cli")
		stringParamPtr := flag.String("param", "hello world", "string parameter for the cli mode")
		wasmFilePathPtr := flag.String("wasm", "", "wasm module file path")
		httpPortPtr := flag.String("httpPort", "8080", "http port")

		wasmFileUrlPtr := flag.String("url", "", "url for downloading wasm module file")

		registryPtr := flag.String("registry", "", "url of the wasm registry")

		crtPtr := flag.String("crt", "", "certificate")
		keyPtr := flag.String("key", "", "key")
		natssrvPtr := flag.String("natssrv", "", "nats server url")
		subjectPtr := flag.String("subject", "", "nats subject(topic)")

		mqttsrvPtr := flag.String("mqttsrv", "", "mqtt server url")
		topicPtr := flag.String("topic", "", "mqtt topic")
		clientIdPtr := flag.String("clientId", "", "mqtt client id")

		flag.Parse()

		flags := CapsuleFlags{
			*capsuleModePtr,
			*stringParamPtr,
			*wasmFilePathPtr,
			*httpPortPtr,
			*wasmFileUrlPtr,
			*crtPtr,
			*keyPtr,
			*registryPtr,
			*natssrvPtr,
			*subjectPtr,
			*mqttsrvPtr,
			*topicPtr,
			*clientIdPtr,
		}

		getWasmFile := func() []byte {
			var wasmFile []byte
			// 📂 Load from file and then Instantiate a WebAssembly module
			loadWasmFile := func(path string) []byte {
				wasmFileToLoad, errLoadWasmFile := os.ReadFile(path)

				if errLoadWasmFile != nil {
					//log.Panicln("🔴 Error while loading the wasm file:", errLoadWasmFile)
					fmt.Println("🔴 Error while loading the wasm file:", errLoadWasmFile)
					os.Exit(1)
				}
				return wasmFileToLoad
			}
			//TODO; add authentication with headers
			if len(flags.url) == 0 {
				wasmFile = loadWasmFile(flags.wasm)
			} else {
				client := resty.New()
				resp, errLoadWasmFileFromUrl := client.R().
					SetOutput(flags.wasm).
					Get(flags.url)

				fmt.Println("📥", "file to download", flags.url)
				fmt.Println("📦", "file saved as", flags.wasm)

				if resp.IsError() {
					fmt.Println("🔴 Error while downloading the wasm file:", "empty response")
					os.Exit(1)
				}

				if errLoadWasmFileFromUrl != nil {
					//log.Panicln("🔴 Error while downloading the wasm file:", errLoadWasmFileFromUrl)
					fmt.Println("🔴 Error while downloading the wasm file:", errLoadWasmFileFromUrl)
					os.Exit(1)
				} else {
					fmt.Println("🙂", "file downloaded", flags.wasm)
				}

				wasmFile = loadWasmFile(flags.wasm)
			}
			return wasmFile
		}

		switch what := flags.mode; what {
		case "http":
			capsulehttp.Serve(flags.httpPort, getWasmFile(), flags.crt, flags.key)
		case "cli":
			capsulecli.Execute(flag.Args(), getWasmFile())
		case "nats":
			capsulenats.Listen(flags.natssrv, flags.subject, getWasmFile())
		case "mqtt":
			capsulemqtt.Listen(flags.mqttsrv, flags.clientId, flags.topic, getWasmFile())
		default:
			fmt.Println("🔴 bad mode", *capsuleModePtr)
			//os.Exit(1)
		}
	}

}
