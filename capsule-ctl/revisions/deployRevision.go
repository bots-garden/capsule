package revisions

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
)

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
		os.Exit(1)
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
		os.Exit(1)
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
		os.Exit(1)
	} else {
		/*
		   {"code":"FUNCTION_DEPLOYED","function":"hello","localUrl":"http://localhost:10001","message":"Function deployed","remoteUrl":"http://localhost:8888/functions/hello/blue","revision":"blue"}
		*/
		jsonRespMap := make(map[string]interface{})
		jsonRespMapErr := json.Unmarshal([]byte(resp.String()), &jsonRespMap)
		if jsonRespMapErr != nil {
			fmt.Println("😡", "[(resp->map)deploying function revision]", jsonRespMapErr)
			os.Exit(1)
		}

		if jsonRespMap["code"] == "KO" {
			fmt.Println("😡", "[when deploying to worker]", jsonRespMap["message"])
			os.Exit(1)

		} else {
			fmt.Println("🙂", "[deployed to worker]", functionName, "/", revisionName)
			fmt.Println("🌍", "[serving]", jsonRespMap["remoteUrl"])
			os.Exit(0)
		}

	}

}
