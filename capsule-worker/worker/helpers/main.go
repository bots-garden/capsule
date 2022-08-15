package helpers

import (
	"fmt"
	"github.com/bots-garden/capsule/capsule-worker/worker/models"
	"github.com/go-resty/resty/v2"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func isDefaultRevision(functionName, revisionName, reverseProxy, backend string) bool {

	client := resty.New()

	respDefaultRevision, errDefaultRevision := client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json; charset=utf-8").
		Get(reverseProxy + "/" + backend + "/functions/" + functionName + "/" + "default" + "/urls")

	if errDefaultRevision != nil {
		fmt.Println("😡errDefaultRevision", errDefaultRevision)

	}

	respCurrentRevision, errCurrentRevision := client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json; charset=utf-8").
		Get(reverseProxy + "/" + backend + "/functions/" + functionName + "/" + revisionName + "/urls")

	if errCurrentRevision != nil {
		fmt.Println("😡errCurrentRevision", errCurrentRevision)
	}

	if errDefaultRevision != nil || errCurrentRevision != nil {
		return false
	}
	if respCurrentRevision.String() == respDefaultRevision.String() {
		return true
	} else {
		return false
	}

}

// JsonFuncList returns a JsonString from the list of the deployed functions
func JsonFuncList(funcList map[string]models.Function, reverseProxy, backend string) string {

	// curl http://localhost:8888/memory/functions/hello/default/urls
	// curl http://localhost:8888/memory/functions/hello/:revision_name/urls

	jsonString := `{`

	for functionName, functionElement := range funcList {

		jsonString += `"` + functionName + `":{`

		fmt.Println("Function:", functionElement.Name, "key:", functionName)
		fmt.Println("  Revisions:")

		for revisionName, revisionElement := range functionElement.Revisions {

			defaultRevision := isDefaultRevision(functionName, revisionName, reverseProxy, backend)

			jsonString += `"` + revisionName + `":{`
			jsonString += `"wasmRegistryUrl":"` + revisionElement.WasmRegistryUrl + `",`
			jsonString += `"isDefaultRevision":"` + strconv.FormatBool(defaultRevision) + `",`
			jsonString += `"wasmModules":{`

			fmt.Println("    ->", revisionElement.Name, "key:", revisionName)
			fmt.Println("      - wasmRegistryUrl:", revisionElement.WasmRegistryUrl)
			fmt.Println("      - (running)wasmModules:")

			for idOfProcess, wasmModuleElement := range revisionElement.WasmModules {

				jsonString += `"` + strconv.Itoa(wasmModuleElement.Pid) + `":{`
				jsonString += `"localUrl":"` + wasmModuleElement.LocalUrl + `",`
				jsonString += `"remoteUrl":"` + wasmModuleElement.RemoteUrl + `",`

				fmt.Println("        ->", wasmModuleElement.Pid, "key:", idOfProcess)
				fmt.Println("          - localUrl", wasmModuleElement.LocalUrl)
				fmt.Println("          - remoteUrl", wasmModuleElement.RemoteUrl)

				// Environment Variables
				fmt.Println("          - envVariables", wasmModuleElement.EnvVariables)

				jsonString += `"envVariables":{`
				for varName, varValue := range wasmModuleElement.EnvVariables {
					jsonString += `"` + varName + `":"` + varValue + `",`
				}
				// remove the last ","
				jsonString = strings.TrimSuffix(jsonString, ",")

				jsonString += `}},` // end of running module
			}
			// remove the last ","
			jsonString = strings.TrimSuffix(jsonString, ",")

			jsonString += `}},` // end of revision
		}
		// remove the last ","
		jsonString = strings.TrimSuffix(jsonString, ",")

		jsonString += `},` // end of function
	}
	// remove the last ","
	jsonString = strings.TrimSuffix(jsonString, ",")

	jsonString += `}` // end of json string

	return jsonString
}

// GetOutboundIP Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func GetModuleServerUrl(workerDomain, moduleServerPort string, httpPortCounter int) string {
	var moduleServerUrl string

	if len(workerDomain) == 0 {
		// 🔎 discovering automatically the domain (or IP address) of the worker
		nodeName, err := os.Hostname()
		if err != nil {
			moduleServerUrl = moduleServerPort + "://" + GetOutboundIP().String() + ":" + strconv.Itoa(httpPortCounter)

		} else {
			moduleServerUrl = moduleServerPort + "://" + nodeName + ":" + strconv.Itoa(httpPortCounter)
		}
	} else {
		moduleServerUrl = moduleServerPort + "://" + workerDomain + ":" + strconv.Itoa(httpPortCounter)
	}
	return moduleServerUrl
}
