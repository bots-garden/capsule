package main

// TinyGo wasm module
import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
	"github.com/tidwall/gjson"
)

func main() {
	hf.SetHandleHttp(Handle)
}

func Handle(bodyReq string, headersReq map[string]string) (response hf.Response, errResp error) {
	bucketName, err := hf.GetEnv("COUCHBASE_BUCKET")
	query := "SELECT * FROM `" + bucketName + "`.data.docs"

	jsonStrArray, err := hf.CouchBaseQuery(query)

	if err != nil {
		hf.Log(err.Error())
	} else {
		jsonArray := gjson.Parse(jsonStrArray).Array()
		for _, jsonDoc := range jsonArray {
			hf.Log("📝: " + jsonDoc.String())
		}
	}

	headersResp := map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}

	// return the first document
	//return gjson.Parse(jsonStrArray).Array()[0].String(), headersResp, err

	// return all documents
	return hf.Response{Body: jsonStrArray, Headers: headersResp}, err
}

/* insert a document:
   query := "INSERT INTO `wasm-data`.data.docs (KEY, VALUE) VALUES (\"key100\", { \"type\" : \"info\", \"name\" : \"this is an info\" });"
*/
