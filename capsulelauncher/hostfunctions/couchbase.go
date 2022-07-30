package hostfunctions

import (
	"context"
	"encoding/json"
	"log"
	"time"

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
        //fmt.Println("✅ initialize cluster connection")
		bucketName := commons.GetEnv("COUCHBASE_BUCKET", "wasm-data")
		username := commons.GetEnv("COUCHBASE_USER", "admin")
		password := commons.GetEnv("COUCHBASE_PWD", "ilovepandas")
        clusterUrl := commons.GetEnv("COUCHBASE_CLUSTER", "couchbase://127.0.0.1")

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
            //fmt.Println("😡 when connecting to the cluster", err.Error())
			log.Fatal(err)
		} else {
            couchBaseCluster = cluster
            //fmt.Println("✅ cluster connected", couchBaseCluster)
        }

        couchBaseBucket := cluster.Bucket(bucketName)

        err = couchBaseBucket.WaitUntilReady(5*time.Second, nil)
        if err != nil {
            //fmt.Println("😡 when connecting to the bucket:", bucketName , err.Error())
            log.Fatal(err)
        } else {
            //fmt.Println("✅ bucket connected", couchBaseBucket)
        }
	}

}
func getCouchBaseBucket() *gocb.Bucket {
	return couchBaseBucket
}
func getCouchBaseCluster() *gocb.Cluster {
	return couchBaseCluster
}

// CouchBaseQuery :
// The wasm module will call this function like this:
// `func CouchBaseQuery(query string) (string, error)`
func CouchBaseQuery(ctx context.Context, module api.Module, queryOffset, queryByteCount, retBuffPtrPos, retBuffSize uint32) {

	InitCouchBaseCluster()

    //gocb.SetLogger(gocb.VerboseStdioLogger())

	//=========================================================
	// Read arguments values of the function call
	// get strings from the wasm module function (from memory)
	//=========================================================
	queryStr := memory.ReadStringFromMemory(ctx, module, queryOffset, queryByteCount)

	//==[👋 Implementation: Start]=============================
	// Perform a N1QL Query
	queryResults, err := getCouchBaseCluster().Query(queryStr, nil)

	var stringResultFromHost = ""
    var records []map[string]interface {}
    var record interface{}

    if err != nil {
        log.Fatalf("%v", err)
    }

	for queryResults.Next() {
		err := queryResults.Row(&record)
		if err != nil {
			panic(err)
		}
		records = append(records, record.(map[string]interface {}))
	}
    jsonStringArray, _ := json.Marshal(records)

	if err != nil {
		stringResultFromHost = commons.CreateStringError(err.Error(), 0)
		// if code 0 don't display code in the error message
	} else {
		stringResultFromHost = string(jsonStringArray)
        if len(stringResultFromHost) == 0 {
            stringResultFromHost="empty"
        }
	}
	//==[👋 Implementation: End]===============================

	// Write the new string stringResultFromHost to the "shared memory"
	memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)
}




