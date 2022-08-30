package revisions

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
)

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
		os.Exit(1)
	}

	bytesSetBody, jsonErr := json.Marshal(setBody)
	if jsonErr != nil {
		fmt.Println("😡", "[(body -> json)setting default revision]", jsonErr)
		os.Exit(1)
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
		os.Exit(1)
	}

	if errSetDefault != nil {
		fmt.Println("😡", "[when setting the default revision]", errSetDefault)
		os.Exit(1)
	} else {

		jsonRespMap := make(map[string]interface{})
		jsonRespMapErr := json.Unmarshal([]byte(resp.String()), &jsonRespMap)
		if jsonRespMapErr != nil {
			fmt.Println("😡", "[(resp->map)setting the default revision]", jsonRespMapErr)
			os.Exit(1)
		}

		if jsonRespMap["code"] == "KO" {
			fmt.Println("😡", "[when setting the default revision]", jsonRespMap["message"])
			os.Exit(1)

		} else {
			fmt.Println("🙂", "[the default revision is set]->", functionName, "/", revisionName)
			fmt.Println("🌍", "[serving]", jsonRespMap["url"])
			os.Exit(0)

		}

		//fmt.Println("🌍", "[serving]", jsonRespMap["url"])
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
		os.Exit(1)
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
		os.Exit(1)
	} else {

		jsonRespMap := make(map[string]interface{})
		jsonRespMapErr := json.Unmarshal([]byte(resp.String()), &jsonRespMap)
		if jsonRespMapErr != nil {
			fmt.Println("😡", "[(resp->map)unsetting the default revision]", jsonRespMapErr)
			os.Exit(1)
		}

		if jsonRespMap["code"] == "KO" {
			fmt.Println("😡", "[when unsetting the default revision]", jsonRespMap["message"])
			os.Exit(1)

		} else {
			fmt.Println("🙂", "[the default revision is unset]->", functionName)
			os.Exit(0)
		}

		//fmt.Println("🌍", "[serving]", jsonRespMap)

	}

}
