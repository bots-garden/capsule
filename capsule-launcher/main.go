package main

import (
    "flag"
    "fmt"
    capsule_cli "github.com/bots-garden/capsule/capsulelauncher/services/cli"
    capsule_http "github.com/bots-garden/capsule/capsulelauncher/services/http"
    "github.com/bots-garden/capsule/commons"
    "github.com/go-resty/resty/v2"
    "log"
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

        registryPtr := flag.String("registry", "", "url of the wasm registry")

        crtPtr := flag.String("crt", "", "certificate")
        keyPtr := flag.String("key", "", "key")

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
        default:
            log.Panicln("🔴 bad mode", *capsuleModePtr)
        }
    }

}
