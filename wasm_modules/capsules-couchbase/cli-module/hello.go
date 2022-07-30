package main

// TinyGo wasm module
import (
    hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
	"github.com/tidwall/gjson"
)

func main() {
    hf.SetHandle(Handle)
}

func Handle(params []string) (string, error) {
    bucketName, err := hf.GetEnv("COUCHBASE_BUCKET")
    query := "SELECT * FROM `" + bucketName +"`.data.docs"
    
    jsonStrArray, err := hf.CouchBaseQuery(query)

    if err != nil {
        hf.Log(err.Error())
    } else {
        jsonArray := gjson.Parse(jsonStrArray).Array()
        for _, jsonDoc := range jsonArray {
            hf.Log("📝: " + jsonDoc.String())
        }
    }
    return jsonStrArray, nil
}
/* insert a document:
    query := "INSERT INTO `wasm-data`.data.docs (KEY, VALUE) VALUES (\"key100\", { \"type\" : \"info\", \"name\" : \"this is an info\" });"
*/

