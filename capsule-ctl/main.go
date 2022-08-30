package main

import (
	"flag"
	"fmt"
	"github.com/bots-garden/capsule/capsule-ctl/registry"
	"github.com/bots-garden/capsule/capsule-ctl/reverseproxy"
	"github.com/bots-garden/capsule/capsule-ctl/revisions"
	"github.com/bots-garden/capsule/capsule-ctl/worker"
	"github.com/bots-garden/capsule/commons"
	"os"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

type CapsuleCtlFlag struct {
	WasmRegistryUrl        string // wasm modules registry related
	WasmModuleOrganization string // wasm modules registry related
	WasmModuleName         string // wasm modules registry related
	WasmModuleTag          string // wasm modules registry related
	//WasmRegistryToken      string // wasm modules registry related
	WasmModuleFile string // wasm modules registry related
	WasmModuleInfo string // wasm modules registry related
	FunctionName   string
	RevisionName   string
	DownloadUrl    string
	EnvVariables   string
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("😮 no args. Type caps help or caps --help")
		os.Exit(0)
	}

	/*
	   You need to use a header with this key: CAPSULE_WORKER_ADMIN_TOKEN
	*/
	adminWorkerToken := GetEnv("CAPSULE_WORKER_ADMIN_TOKEN", "")

	workerUrl := GetEnv("CAPSULE_WORKER_URL", "")

	/*
	   You need to use a header with this key: CAPSULE_REVERSE_PROXY_ADMIN_TOKEN
	*/
	adminReverseProxyToken := GetEnv("CAPSULE_REVERSE_PROXY_ADMIN_TOKEN", "") // right now, not used

	/*
	   You need to use a header with this key: CAPSULE_REGISTRY_ADMIN_TOKEN
	*/
	registryAdminToken := GetEnv("CAPSULE_REGISTRY_ADMIN_TOKEN", "")

	reverseProxyUrl := GetEnv("CAPSULE_REVERSE_PROXY_URL", "")
	backend := GetEnv("CAPSULE_BACKEND", "")

	commands := map[string]string{
		"publish":       "publish a wasm module to the capsule registry",
		"deploy":        "deploy a function's revision",
		"downscale":     "downscale a function's revision",
		"un-deploy":     "undeploy a function's revision",
		"set-default":   "set the default revision of a function (and remove the previous one if it exists)",
		"unset-default": "remove the default revision of a function",
		"worker":        "display information about the worker",
		"reverse-proxy": "display information about the reverse-proxy",
		"help":          "",
		"version":       "get the capsulectl version"}

	if _, ok := commands[args[0]]; ok {
		//fmt.Println("🤖", "command:", args[0], ":", value)
		mainCommand := args[0]

		capsuleCtlFlagSet := flag.NewFlagSet("", flag.ExitOnError)

		// Where to download the wasm module
		wasmRegistryUrlPtr := capsuleCtlFlagSet.String("registryUrl", "", "wasm module registry url")
		//wasmRegistryTokenPtr := capsuleCtlFlagSet.String("registryToken", "", "wasm registry token")
		wasmModuleFilePtr := capsuleCtlFlagSet.String("wasmFile", "", "wasm module location")
		wasmModuleInfoPtr := capsuleCtlFlagSet.String("wasmInfo", "", "wasm module information when publishing to the registry")
		wasmModuleNamePtr := capsuleCtlFlagSet.String("wasmName", "", "wasm module name for publication")
		wasmModuleTagPtr := capsuleCtlFlagSet.String("wasmTag", "", "wasm module tag for publication")
		wasmModuleOrganizationPtr := capsuleCtlFlagSet.String("wasmOrg", "", "Organization for publication of the module")

		functionNamePtr := capsuleCtlFlagSet.String("function", "", "function name")
		revisionNamePtr := capsuleCtlFlagSet.String("revision", "", "revision name")

		downloadUrlPtr := capsuleCtlFlagSet.String("downloadUrl", "", "where to download the wasm module")
		envVariablesPtr := capsuleCtlFlagSet.String("envVariables", "{}", "environment variables for the module execution")

		err := capsuleCtlFlagSet.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("😡", err)
			os.Exit(1)
		}

		flags := CapsuleCtlFlag{
			RevisionName: *revisionNamePtr,
			FunctionName: *functionNamePtr,
			DownloadUrl:  *downloadUrlPtr,
			EnvVariables: *envVariablesPtr,
			//WasmRegistryToken:      *wasmRegistryTokenPtr,
			WasmRegistryUrl:        *wasmRegistryUrlPtr,
			WasmModuleFile:         *wasmModuleFilePtr,
			WasmModuleInfo:         *wasmModuleInfoPtr,
			WasmModuleName:         *wasmModuleNamePtr,
			WasmModuleTag:          *wasmModuleTagPtr,
			WasmModuleOrganization: *wasmModuleOrganizationPtr,
		}

		switch mainCommand {
		case "publish":
			/*
			   ./capsulectl publish \
			   -wasmFile=./hello/hello.wasm -wasmInfo=wip \
			   -wasmOrg=k33g -wasmName=hello -wasmTag=0.0.0 \
			   -registryUrl=http://localhost:4999
			*/
			registry.PublishToTheRegistry(
				flags.WasmModuleFile,
				flags.WasmModuleInfo,
				flags.WasmModuleOrganization,
				flags.WasmModuleName,
				flags.WasmModuleTag,
				flags.WasmRegistryUrl,
				registryAdminToken)
			//flags.WasmRegistryToken

		case "deploy":
			/*
			   CAPSULE_WORKER_URL="http://localhost:9999" ./capsulectl deploy \
			   -function=hello \
			   -revision=blue \
			   -downloadUrl=http://localhost:4999/k33g/hello/0.0.0/hello.wasm \
			   -envVariables={"MESSAGE": "Revision 🔵","TOKEN": "👩‍🔧🧑‍🔧👨‍🔧"}
			*/
			revisions.DeployFunctionRevision(
				flags.FunctionName,
				flags.RevisionName,
				flags.DownloadUrl,
				flags.EnvVariables,
				workerUrl,
				adminWorkerToken)

		case "un-deploy":
			/*
			   CAPSULE_WORKER_URL="http://localhost:9999" ./capsulectl un-deploy \
			   -function=hello \
			   -revision=blue
			*/
			revisions.UnDeployRevision(
				flags.FunctionName,
				flags.RevisionName,
				workerUrl,
				adminWorkerToken)

		case "set-default":
			/*
			   CAPSULE_WORKER_URL="http://localhost:9999" ./capsulectl set-default \
			   -function=hello \
			   -revision=blue
			*/
			revisions.SetDefaultRevision(
				flags.FunctionName,
				flags.RevisionName,
				workerUrl,
				adminWorkerToken)

		case "unset-default":
			/*
			   CAPSULE_WORKER_URL="http://localhost:9999" ./capsulectl unset-default \
			   -function=hello
			*/
			revisions.UnSetDefaultRevision(
				flags.FunctionName,
				workerUrl,
				adminWorkerToken)

		case "downscale":
			//TODO: check if default revision exist for this revision
			// remove url from default too
			/*
			   CAPSULE_WORKER_URL="http://localhost:9999" ./capsulectl downscale \
			   -function=hello \
			   -revision=orange
			*/

			revisions.DownscaleRevision(
				flags.FunctionName,
				flags.RevisionName,
				workerUrl,
				adminWorkerToken)

		case "worker":
			worker.WorkerInfo(workerUrl, adminWorkerToken, backend)

		case "reverse-proxy":
			reverseproxy.ReverseProxyInfo(reverseProxyUrl, adminReverseProxyToken, backend)

		case "version":
			fmt.Println(commons.CapsuleVersion())
			os.Exit(0)

		case "help":
			//TODO: add help for the flags
			for cmd, text := range commands {
				if cmd != "help" {
					fmt.Println("-", cmd, ":", text)
				}
			}

		default:
			//log.Panicln("😡", "Houston, we have a problem")
			fmt.Println("😡", "Houston, we have a problem")
			os.Exit(1)
		}

	} else {
		fmt.Println("😡", args[0], "is not a recognized command")
		os.Exit(1)
	}

}
