package hostfunctions

import (
	"context"
	"fmt"
	"log"

	//"os"
	//"strconv"

	"github.com/bots-garden/capsule/capsulelauncher/commons"
	"github.com/bots-garden/capsule/capsulelauncher/hostfunctions/memory"
	"github.com/tetratelabs/wazero/api"

	"github.com/couchbase/gocb/v2"
)

var couchBaseBucket *gocb.Bucket
var couchBaseCluster *gocb.Cluster


//TODO: implement certificate management

func InitCouchBaseCluster() {
	if couchBaseCluster == nil {
        fmt.Println("✅ cluster")
		//bucketName := commons.GetEnv("COUCHBASE_BUCKET", "wasm-data")
		username := commons.GetEnv("COUCHBASE_USER", "admin")
		password := commons.GetEnv("COUCHBASE_PWD", "ilovepandas")
        clusterUrl := commons.GetEnv("COUCHBASE_CLUSTER", "couchbases://127.0.0.1")

		// Initialize the Connection
		cluster, err := gocb.Connect(clusterUrl, gocb.ClusterOptions{
			Authenticator: gocb.PasswordAuthenticator{
				Username: username,
				Password: password,
			},
            SecurityConfig: gocb.SecurityConfig{
                //TLSRootCAs: roots,
                // WARNING: Do not set this to true in production, only use this for testing!
                TLSSkipVerify: true,
            },
		})
		if err != nil {
			log.Fatal(err)
		} else {
            fmt.Println("✅ cluster connected")
        }
        couchBaseCluster = cluster
        //couchBaseBucket := cluster.Bucket(bucketName)
        /*
        err = couchBaseBucket.WaitUntilReady(5*time.Second, nil)
        if err != nil {
            log.Fatal(err)
        }
        */
	}

}
func getCouchBaseBucket() *gocb.Bucket {
	return couchBaseBucket
}
func getCouchBaseCluster() *gocb.Cluster {
	return couchBaseCluster
}
/*

SELECT * FROM `wasm-data`.data.docs

script1='CREATE SCOPE `wasm-data`.data'
script2='CREATE COLLECTION `wasm-data`.data.docs'
script3='INSERT INTO `wasm-data`.data.docs (KEY, VALUE) VALUES ("key1", { "type" : "info", "name" : "this is an info" });'
script4='INSERT INTO `wasm-data`.data.docs (KEY, VALUE) VALUES ("key2", { "type" : "info", "name" : "this is another info" });'
script5='INSERT INTO `wasm-data`.data.docs (KEY, VALUE) VALUES ("key3", { "type" : "message", "name" : "👋 hello world 🌍" });'
script6='INSERT INTO `wasm-data`.data.docs (KEY, VALUE) VALUES ("key4", { "type" : "message", "name" : "👋 greetings 🎉" });'
*/

// CouchBaseQuery :
// The wasm module will call this function like this:
// `func CouchBaseQuery(query string) (string, error)`
func CouchBaseQuery(ctx context.Context, module api.Module, queryOffset, queryByteCount, retBuffPtrPos, retBuffSize uint32) {

	InitCouchBaseCluster()
	//=========================================================
	// Read arguments values of the function call
	// get strings from the wasm module function (from memory)
	//=========================================================
	queryStr := memory.ReadStringFromMemory(ctx, module, queryOffset, queryByteCount)

    fmt.Println("✅ query:", queryStr)

	//==[👋 Implementation: Start]=============================
	// Perform a N1QL Query
	queryResults, err := getCouchBaseCluster().Query(queryStr, nil)
    fmt.Println("✅ queryResults:", queryResults)
	var stringResultFromHost = ""

    var records []string
    var record interface{}

	for queryResults.Next() {
		err := queryResults.Row(&record)
		if err != nil {
			panic(err)
		}
        fmt.Println("=>", record)
		records = append(records, record.(string))
	}

	if err != nil {
		stringResultFromHost = commons.CreateStringError(err.Error(), 0)
		// if code 0 don't display code in the error message
	} else {
		stringResultFromHost = commons.CreateStringFromSlice(records, "|")
	}
	//==[👋 Implementation: End]===============================

	// Write the new string stringResultFromHost to the "shared memory"
	memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)
}




