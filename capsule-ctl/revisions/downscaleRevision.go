package revisions

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
)

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
		Delete(workerUrl + "/functions/revisions/downscale")

	if errSetDefault != nil {
		fmt.Println("😡", "[when downscaling the revision]", errSetDefault)
		os.Exit(1)
	} else {

		jsonRespMap := make(map[string]interface{})
		jsonRespMapErr := json.Unmarshal([]byte(resp.String()), &jsonRespMap)
		if jsonRespMapErr != nil {
			fmt.Println("😡", "[(resp->map)downscaling the revision]", jsonRespMapErr)
			os.Exit(1)
		}
		//fmt.Println("🌍", "[serving]", jsonRespMap)
		if jsonRespMap["code"] == "WASM_MODULE_DEPLOYMENT_NOT_REMOVED" {
			fmt.Println("😡", "[the revision is not downscaled]-> the revision needs at least one running wasm module")
			os.Exit(1)
		} else {
			fmt.Println("🙂", "[the revision is downscaled (one process killed)]->", functionName, "/", revisionName, "pid:", jsonRespMap["pid"])
			os.Exit(0)
		}

	}

}
