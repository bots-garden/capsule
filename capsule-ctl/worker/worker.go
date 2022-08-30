package worker

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
)

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
		os.Exit(1)

	} else {
		fmt.Println(resp)
		os.Exit(0)
	}

}
