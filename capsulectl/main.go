package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bots-garden/capsule/capsulectl/commons"
	"github.com/go-resty/resty/v2"
	"log"
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
	WasmRegistryToken      string // wasm modules registry related
	WasmModuleFile         string // wasm modules registry related
	WasmModuleInfo         string // wasm modules registry related
	FunctionName           string
	RevisionName           string
	DownloadUrl            string
	EnvVariables           string
}

/*
# Publish the wasm module to the registry
# 🖐 change the tag if you publish a new version
curl -X POST http://localhost:4999/upload/k33g/hello/0.0.0 \
  -F "file=@./hello/hello.wasm" \
  -F "info=hello function v0.0.0 from @k33g [GET]" \
  -H "Content-Type: multipart/form-data"

*/

func PublishToTheRegistry(wasmModuleFile, wasmModuleInfo, wasmModuleOrg, wasmModuleName, wasmModuleTag, wasmRegistryUrl, wasmRegistryToken string) {
	//TODO: make it wasmer.io compliant
	//fmt.Println(wasmModuleFile, wasmModuleInfo, wasmModuleOrg, wasmModuleName, wasmModuleTag, wasmRegistryUrl, wasmRegistryToken)

	fmt.Println("⏳", "[publishing to registry]", wasmModuleOrg, wasmModuleName, wasmModuleTag)

	client := resty.New()
	_, err := client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "multipart/form-data").
		SetFile("file", wasmModuleFile).
		SetMultipartFormData(map[string]string{
			"info": wasmModuleInfo,
		}).
		Post(wasmRegistryUrl + "/upload/" + wasmModuleOrg + "/" + wasmModuleName + "/" + wasmModuleTag)
	if err != nil {
		fmt.Println("😡", "[publishing to registry]", err)
	} else {
		fmt.Println("🙂", "[publishing to registry]", wasmModuleFile, "published")
	}

}

/*
curl -v -X POST \
http://localhost:9999/functions/deploy \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hello",
    "revision": "blue",
    "downloadUrl": "http://localhost:4999/k33g/hello/0.0.0/hello.wasm",
    "envVariables": {
        "MESSAGE": "Revision 🔵",
        "TOKEN": "👩‍🔧🧑‍🔧👨‍🔧"
    }
}
*/

func DeployFunctionRevision(functionName, revisionName, downloadUrl, envVariables, workerUrl, workerToken string) {
	fmt.Println("⏳", "[deploying to worker]", functionName, "/", revisionName)

	jsonEnvMapMap := make(map[string]interface{})
	jsonMapErr := json.Unmarshal([]byte(envVariables), &jsonEnvMapMap)
	if jsonMapErr != nil {
		fmt.Println("😡", "[(envVariables->map)deploying function revision]", jsonMapErr)
	}

	body := map[string]interface{}{
		"function":     functionName,
		"revision":     revisionName,
		"downloadUrl":  downloadUrl,
		"envVariables": jsonEnvMapMap,
	}

	bytesBody, jsonErr := json.Marshal(body)
	if jsonErr != nil {
		fmt.Println("😡", "[(body -> json)deploying function revision]", jsonErr)
	}

	jsonStringBody := string(bytesBody)

	client := resty.New()
	resp, err := client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json; charset=utf-8").
		SetHeader("CAPSULE_WORKER_ADMIN_TOKEN", workerToken).
		SetBody(jsonStringBody).
		Post(workerUrl + "/functions/deploy")

	if err != nil {
		fmt.Println("😡", "[when deploying to worker]", err)
	} else {
		fmt.Println("🙂", "[deployed to worker]", functionName, "/", revisionName)
		/*
		   {"code":"FUNCTION_DEPLOYED","function":"hello","localUrl":"http://localhost:10001","message":"Function deployed","remoteUrl":"http://localhost:8888/functions/hello/blue","revision":"blue"}
		*/
		jsonRespMap := make(map[string]interface{})
		jsonRespMapErr := json.Unmarshal([]byte(resp.String()), &jsonRespMap)
		if jsonRespMapErr != nil {
			fmt.Println("😡", "[(resp->map)deploying function revision]", jsonRespMapErr)
		}

		fmt.Println("🌍", "[serving]", jsonRespMap["remoteUrl"])
	}

}

/*
# Remove default revision if it exists
curl -v -X DELETE \
http://localhost:9999/functions/remove_default_revision \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hello"
}
EOF

# Now the green revision is the default revision
curl -v -X POST \
http://localhost:9999/functions/set_default_revision \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hello",
    "revision": "green"
}
EOF
*/

func SetDefaultRevision(functionName, revisionName, workerUrl, workerToken string) {
	fmt.Println("⏳", "[setting default revision]", functionName, "/", revisionName)

	unSetBody := map[string]interface{}{
		"function": functionName,
	}

	setBody := map[string]interface{}{
		"function": functionName,
		"revision": revisionName,
	}

	bytesUnSetBody, jsonErr := json.Marshal(unSetBody)
	if jsonErr != nil {
		fmt.Println("😡", "[(body -> json)unsetting default revision]", jsonErr)
	}

	bytesSetBody, jsonErr := json.Marshal(setBody)
	if jsonErr != nil {
		fmt.Println("😡", "[(body -> json)setting default revision]", jsonErr)
	}

	jsonStringUnSetBody := string(bytesUnSetBody)
	jsonStringSetBody := string(bytesSetBody)

	client := resty.New()

	resp, errUnSetDefault := client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json; charset=utf-8").
		SetHeader("CAPSULE_WORKER_ADMIN_TOKEN", workerToken).
		SetBody(jsonStringUnSetBody).
		Delete(workerUrl + "/functions/remove_default_revision")

	resp, errSetDefault := client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json; charset=utf-8").
		SetHeader("CAPSULE_WORKER_ADMIN_TOKEN", workerToken).
		SetBody(jsonStringSetBody).
		Post(workerUrl + "/functions/set_default_revision")

	if errUnSetDefault != nil {
		fmt.Println("😡", "[when unsetting the default revision]", errUnSetDefault)
	}

	if errSetDefault != nil {
		fmt.Println("😡", "[when setting the default revision]", errSetDefault)
	} else {
		fmt.Println("🙂", "[the default revision is set]->", functionName, "/", revisionName)

		jsonRespMap := make(map[string]interface{})
		jsonRespMapErr := json.Unmarshal([]byte(resp.String()), &jsonRespMap)
		if jsonRespMapErr != nil {
			fmt.Println("😡", "[(resp->map)setting the default revision]", jsonRespMapErr)
		}

		fmt.Println("🌍", "[serving]", jsonRespMap["url"])
		//fmt.Println("🌍", "[serving]", jsonRespMap)

	}

}

/*
# Remove default revision if it exists
curl -v -X DELETE \
http://localhost:9999/functions/remove_default_revision \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hello"
}
EOF
*/

func UnSetDefaultRevision(functionName, workerUrl, workerToken string) {
	fmt.Println("⏳", "[unsetting default revision]", functionName)

	unSetBody := map[string]interface{}{
		"function": functionName,
	}

	bytesUnSetBody, jsonErr := json.Marshal(unSetBody)
	if jsonErr != nil {
		fmt.Println("😡", "[(body -> json)unsetting default revision]", jsonErr)
	}

	jsonStringUnSetBody := string(bytesUnSetBody)

	client := resty.New()

	resp, errUnSetDefault := client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json; charset=utf-8").
		SetHeader("CAPSULE_WORKER_ADMIN_TOKEN", workerToken).
		SetBody(jsonStringUnSetBody).
		Delete(workerUrl + "/functions/remove_default_revision")

	if errUnSetDefault != nil {
		fmt.Println("😡", "[when unsetting the default revision]", errUnSetDefault)
	} else {
		fmt.Println("🙂", "[the default revision is unset]->", functionName)

		jsonRespMap := make(map[string]interface{})
		jsonRespMapErr := json.Unmarshal([]byte(resp.String()), &jsonRespMap)
		if jsonRespMapErr != nil {
			fmt.Println("😡", "[(resp->map)unsetting the default revision]", jsonRespMapErr)
		}

		//fmt.Println("🌍", "[serving]", jsonRespMap)

	}

}

/*
curl -v -X DELETE \
http://localhost:9999/functions/revisions/deployments \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hello",
    "revision": "blue"
}
EOF
*/

func UnDeployRevision(functionName, revisionName, workerUrl, workerToken string) {
	fmt.Println("⏳", "[un-deploying revision]", functionName, "/", revisionName)

	setBody := map[string]interface{}{
		"function": functionName,
		"revision": revisionName,
	}

	bytesSetBody, jsonErr := json.Marshal(setBody)
	if jsonErr != nil {
		fmt.Println("😡", "[(body -> json)un-deploying revision]", jsonErr)
	}

	jsonStringSetBody := string(bytesSetBody)

	client := resty.New()

	resp, errSetDefault := client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json; charset=utf-8").
		SetHeader("CAPSULE_WORKER_ADMIN_TOKEN", workerToken).
		SetBody(jsonStringSetBody).
		Delete(workerUrl + "/functions/revisions/deployments")

	if errSetDefault != nil {
		fmt.Println("😡", "[when un-deploying the revision]", errSetDefault)
	} else {
		fmt.Println("🙂", "[the revision is un-deployed (all processes killed)]->", functionName, "/", revisionName)

		jsonRespMap := make(map[string]interface{})
		jsonRespMapErr := json.Unmarshal([]byte(resp.String()), &jsonRespMap)
		if jsonRespMapErr != nil {
			fmt.Println("😡", "[(resp->map)un-deploying the revision]", jsonRespMapErr)
		}

		//fmt.Println("🌍", "[serving]", jsonRespMap["url"])
		//fmt.Println("🌍", "[serving]", jsonRespMap)

	}

}

/*
curl -v -X DELETE \
http://localhost:9999/functions/revisions/downscale \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hello",
    "revision": "orange"
}
EOF
*/

func DownscaleRevision(functionName, revisionName, workerUrl, workerToken string) {
	fmt.Println("⏳", "[downscaling revision]", functionName, "/", revisionName)

	setBody := map[string]interface{}{
		"function": functionName,
		"revision": revisionName,
	}

	bytesSetBody, jsonErr := json.Marshal(setBody)
	if jsonErr != nil {
		fmt.Println("😡", "[(body -> json)downscaling revision]", jsonErr)
	}

	jsonStringSetBody := string(bytesSetBody)

	client := resty.New()

	resp, errSetDefault := client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json; charset=utf-8").
		SetHeader("CAPSULE_WORKER_ADMIN_TOKEN", workerToken).
		SetBody(jsonStringSetBody).
		Delete(workerUrl + "/functions/revisions/downscale")

	if errSetDefault != nil {
		fmt.Println("😡", "[when downscaling the revision]", errSetDefault)
	} else {

		jsonRespMap := make(map[string]interface{})
		jsonRespMapErr := json.Unmarshal([]byte(resp.String()), &jsonRespMap)
		if jsonRespMapErr != nil {
			fmt.Println("😡", "[(resp->map)downscaling the revision]", jsonRespMapErr)
		}
		//fmt.Println("🌍", "[serving]", jsonRespMap)
		if jsonRespMap["code"] == "WASM_MODULE_DEPLOYMENT_NOT_REMOVED" {
			fmt.Println("😡", "[the revision is not downscaled]-> the revision needs at least one running wasm module")
		} else {
			fmt.Println("🙂", "[the revision is downscaled (one process killed)]->", functionName, "/", revisionName, "pid:", jsonRespMap["pid"])

		}

	}

}

func WorkerInfo(workerUrl, adminWorkerToken, backend string) {
	//TODO: change the route of the worker to taking account of the backend
	// curl http://localhost:9999/functions/list
	// fmt.Println(workerUrl, adminWorkerToken, backend)

	client := resty.New()

	resp, err := client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json; charset=utf-8").
		Get(workerUrl + "/functions/list")

	if err != nil {
		fmt.Println("😡", err)

	} else {
		fmt.Println(resp)
	}

}

func ReverseProxyInfo(reverseProxyUrl, adminReverseProxyToken, backend string) {
	//curl http://localhost:8888/memory/functions/list
	//fmt.Println(reverseProxyUrl, adminReverseProxyToken, backend)

	client := resty.New()

	resp, err := client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json; charset=utf-8").
		Get(reverseProxyUrl + "/" + backend + "/functions/list")

	if err != nil {
		fmt.Println("😡", err)
	} else {
		fmt.Println(resp)
	}

}

func main() {
	args := os.Args[1:]

	/*
	   You need to use a header with this key: CAPSULE_WORKER_ADMIN_TOKEN
	*/
	adminWorkerToken := GetEnv("CAPSULE_WORKER_ADMIN_TOKEN", "")

	workerUrl := GetEnv("CAPSULE_WORKER_URL", "")

	/*
	   You need to use a header with this key: CAPSULE_REVERSE_PROXY_ADMIN_TOKEN
	*/
	adminReverseProxyToken := GetEnv("CAPSULE_REVERSE_PROXY_ADMIN_TOKEN", "") // right now, not used

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
		wasmRegistryTokenPtr := capsuleCtlFlagSet.String("registryToken", "", "wasm registry token")
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
		}

		flags := CapsuleCtlFlag{
			RevisionName:           *revisionNamePtr,
			FunctionName:           *functionNamePtr,
			DownloadUrl:            *downloadUrlPtr,
			EnvVariables:           *envVariablesPtr,
			WasmRegistryToken:      *wasmRegistryTokenPtr,
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
			   -registryUrl=http://localhost:4999 \
			   -registryToken=nothing
			*/
			PublishToTheRegistry(
				flags.WasmModuleFile,
				flags.WasmModuleInfo,
				flags.WasmModuleOrganization,
				flags.WasmModuleName,
				flags.WasmModuleTag,
				flags.WasmRegistryUrl,
				flags.WasmRegistryToken)

		case "deploy":
			/*
			   CAPSULE_WORKER_URL="http://localhost:9999" ./capsulectl deploy \
			   -function=hello \
			   -revision=blue \
			   -downloadUrl=http://localhost:4999/k33g/hello/0.0.0/hello.wasm \
			   -envVariables={"MESSAGE": "Revision 🔵","TOKEN": "👩‍🔧🧑‍🔧👨‍🔧"}
			*/
			DeployFunctionRevision(
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
			UnDeployRevision(
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
			SetDefaultRevision(
				flags.FunctionName,
				flags.RevisionName,
				workerUrl,
				adminWorkerToken)

		case "unset-default":
			/*
			   CAPSULE_WORKER_URL="http://localhost:9999" ./capsulectl unset-default \
			   -function=hello
			*/
			UnSetDefaultRevision(
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

			DownscaleRevision(
				flags.FunctionName,
				flags.RevisionName,
				workerUrl,
				adminWorkerToken)

		case "worker":
			WorkerInfo(workerUrl, adminWorkerToken, backend)

		case "reverse-proxy":
			ReverseProxyInfo(reverseProxyUrl, adminReverseProxyToken, backend)

		case "version":
			fmt.Println(commons.CapsuleCtlVersion())

		case "help":
			//TODO: add help for the flags
			for cmd, text := range commands {
				if cmd != "help" {
					fmt.Println("-", cmd, ":", text)
				}
			}

		default:
			log.Panicln("😡", "Houston, we have a problem")
		}

	} else {
		fmt.Println("😡", args[0], "is not a recognized command")
	}

}
