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
	newRevision string
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
	newRevisionPtr := flag.String("newRevision", "", "used when duplicationg an existing revision")
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
		*newRevisionPtr,
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

	capsuleFaasToken := GetEnv("CAPSULE_FAAS_TOKEN", "")
	httpClient.SetHeader("CAPSULE_FAAS_TOKEN", capsuleFaasToken)
	

	switch what := flags.cmd; what {
	case "launch":
		// this is a work in progress
		// run a first capsule instance
		// is it useful ?
		// TODO: plan to add a flag "faas" to Capsule HTTP
	case "start":
		/* example:
		   capsctl \
		     --cmd=start \
		     --name=${FUNCTION_NAME} \
		     --revision=${WASM_VERSION} \
		     --description="hello green deployment" \
		     --path=${CAPSULE_INSTALL_PATH} \
		     --wasm=${WASM_FILES_LOCATION}/${FUNCTION_NAME}-${WASM_VERSION}/${WASM_MODULE}
		*/
		args := []string{
			"",
			"-wasm=" + flags.wasm,
			"-stopAfter=" + flags.stopAfter, // if the module is not used, it will be stopped after n seconds
			"-url=" + flags.url,
			"-parentEndpoint=" + capsuleURI,
			"-moduleName=" + flags.name,
			"-moduleRevision=" + flags.revision,
		}

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

		resp, err := httpClient.R().EnableTrace().SetBody(strBodyRequest).Post(capsuleURI + "/functions/start")

		if err != nil {
			fmt.Println("🔴 Error when starting a wasm module", flags.wasm, err)
			return
		}
		if resp.StatusCode() == 404 {
			fmt.Println("🤚", resp.String(), resp.Status())
			fmt.Println("🔴 It faas mode not activated or bad token")
			return
		}

		fmt.Println("✅", flags.name+"/"+flags.revision, "is started")
		
		fmt.Println("ℹ️", "url:", capsuleURI+"/functions/"+flags.name+"/"+flags.revision)

	case "drop":
		/* example:
		   capsctl \
		     --cmd=stop \
		     --name=${FUNCTION_NAME} \
		     --revision=${WASM_VERSION} \
		*/
		resp, err := httpClient.R().EnableTrace().Delete(capsuleURI + "/functions/drop/" + flags.name + "/" + flags.revision)

		if err != nil {
			fmt.Println("🔴 Error when dropping a wasm module", flags.wasm, err)
		}
		if resp.StatusCode() == 404 {
			fmt.Println("🤚", resp.String(), resp.Status())
			fmt.Println("🔴 It faas mode not activated or bad token")
			return
		}

		fmt.Println("✅", flags.name+"/"+flags.revision, "is dropped")
		fmt.Println("ℹ️", "url:", capsuleURI+"/functions/"+flags.name+"/"+flags.revision)

	case "duplicate": // fork?
		/* example:
		   capsctl \
		     --cmd=duplicate \
		     --name=${FUNCTION_NAME} \
		     --revision=${WASM_VERSION} \
			 --newRevision=saved_${WASM_VERSION}
		*/
		// !------------------------------------------------
		// ! It's the same process (from a PID perspective)
		// !------------------------------------------------
		// It's like creating a revision with the same module
		// curl -X PUT http://localhost:8080/functions/duplicate/hello-world/default/saved
		// then call the duplicate:
		/*
			JSON_DATA='{"name":"Bob Morane 🤣","age":42}'

			curl -X POST http://localhost:8080/functions/hello-world/saved \
					-H 'Content-Type: application/json; charset=utf-8' \
					-d "${JSON_DATA}"
		*/
		
		/* kind of 🔵🟢 deployment (to be tested)
			duplicate the current "defautl" revision to "saved" revision
			duplicate the "blue" revision to "default"
			drop the "saved" revision
		*/

		resp, err := httpClient.R().EnableTrace().
			Put(capsuleURI +
				"/functions/duplicate/" +
				flags.name + "/" +
				flags.revision + "/" +
				flags.newRevision)

		if err != nil {
			fmt.Println("🔴 Error when duplication a wasm module revision", flags.wasm, err)
		}

		if resp.StatusCode() == 404 {
			fmt.Println("🤚", resp.String(), resp.Status())
			fmt.Println("🔴 faas mode not activated or bad token")
			return
		}

		fmt.Println("✅", flags.name+"/"+flags.revision, "is duplicated to", flags.newRevision)
		fmt.Println("ℹ️", "url:", capsuleURI+"/functions/"+flags.name+"/"+flags.newRevision)

	case "scale":
		// this is a work in progress
	case "processes":
		// this is a work in progress

	case "publish":
		// this is a work in progress
		// publish a wasm file to a store (eg: GitLab registry, wapm.io, ...)

	case "call":
		fmt.Println("🙏 Please, use curl to call a function")

	default:
		fmt.Printf("🔴 Unknown command: %s\n", what)
	}

}
