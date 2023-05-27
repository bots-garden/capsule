// Package main -> CLI to interact with Capsule HTTP
package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	//"github.com/go-resty/resty/v2"
)

// GetEnv returns the environment variable
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

/*
## List of commands

- start

*/

// StartFunctionParameters handles params for the start of a function (== start cmd)
type StartFunctionParameters struct {
	Name        string   `json:"name"`
	Revision    string   `json:"revision"`
	Description string   `json:"description"`
	Path        string   `json:"path"`
	Args        []string `json:"args"`
	Env         []string `json:"env"`
}

// CapsCtlFlags handles params for the CapsCtl CLI
type CapsCtlFlags struct {
	cmd         string
	name        string
	revision    string
	description string
	wasm        string // wasm file location
	stopAfter   string // stop after a delay if not used
	url         string // to download the wasm file
	path        string // capsuleHTTPPath
	env         string
	version     bool
}

//go:embed description.txt
var textVersion []byte


func main() {
	version := string(textVersion)
	args := os.Args[1:]

	if len(args) == 0 {
		log.Println("CapsCtl needs some args to start.")
		os.Exit(0)
	}

	// CapsCtl flags
	cmdPtr := flag.String("cmd", "", "CapsCtl command")
	namePtr := flag.String("name", "", "function/module name")
	revisionPtr := flag.String("revision", "", "function/module revision")
	description := flag.String("description", "", "function/module description")
	wasmFilePathPtr := flag.String("wasm", "", "wasm module file path")
	stopAfterPtr := flag.String("stopAfter", "", "stop after n seconds if not used")
	wasmFileURLPtr := flag.String("url", "", "url for downloading wasm module file")
	pathPtr := flag.String("path", "", "path of Capsule HTTP binary")
	envPrt := flag.String("env", "", `example:["MESSAGE='Hello'"]`)
	versionPtr := flag.Bool("version", false, "prints CapsCtl current version")

	flag.Parse()

	if *versionPtr {
		fmt.Println(version)
		os.Exit(0)
	}

	flags := CapsCtlFlags{
		*cmdPtr,
		*namePtr,
		*revisionPtr,
		*description,
		*wasmFilePathPtr,
		*stopAfterPtr,
		*wasmFileURLPtr,
		*pathPtr,
		*envPrt,
		*versionPtr,
	}

	httpClient := resty.New()
	capsuleURI := GetEnv("CAPSULE_MAIN_PROCESS_URL", "http://localhost:8080")
	httpClient.SetHeader("Content-Type", "application/json; charset=utf-8")

	switch what := flags.cmd; what {
	case "start":
		args := []string{"", "-wasm=" + flags.wasm, "-stopAfter=" + flags.stopAfter, "-url=" + flags.url}
		envStr := `{"env":` + flags.env + `}`

		var value map[string][]string

		json.Unmarshal([]byte(envStr), &value)

		bodyRequest := StartFunctionParameters{
			Name:        flags.name,
			Revision:    flags.revision,
			Description: flags.description,
			Path:        flags.path,
			Args:        args,
			Env:         value["env"], // "env": ["MESSAGE='hello world'"]

		}
		buff, _ := json.Marshal(bodyRequest)
		strBodyRequest := string(buff)

		_, err := httpClient.R().EnableTrace().SetBody(strBodyRequest).Post(capsuleURI + "/functions/start")

		if err != nil {
			fmt.Println("🔴 Error when starting a wasm module", flags.wasm, err)
		}
		fmt.Println("✅", flags.name+"/"+flags.revision, "is started")
		// resp.String(),
		fmt.Println("ℹ️", "url:", capsuleURI+"/functions/"+flags.name+"/"+flags.revision)

	default:
		fmt.Printf("🔴 Unknown command: %s\n", what)
	}

}
