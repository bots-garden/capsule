package revisions

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
)

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
		os.Exit(1)
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
		os.Exit(1)
	} else {

		jsonRespMap := make(map[string]interface{})
		jsonRespMapErr := json.Unmarshal([]byte(resp.String()), &jsonRespMap)
		if jsonRespMapErr != nil {
			fmt.Println("😡", "[(resp->map)un-deploying the revision]", jsonRespMapErr)
			os.Exit(1)
		}

		if jsonRespMap["code"] == "KO" {
			fmt.Println("😡", "[when un-deploying the revision]", jsonRespMap["message"])
			os.Exit(1)

		} else {
			fmt.Println("🙂", "[the revision is un-deployed (all processes killed)]->", functionName, "/", revisionName)
			os.Exit(0)
		}

		//fmt.Println("🌍", "[serving]", jsonRespMap["url"])
		//fmt.Println("🌍", "[serving]", jsonRespMap)

	}

}
