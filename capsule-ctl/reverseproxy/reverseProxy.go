package reverseproxy

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
)

func ReverseProxyInfo(reverseProxyUrl, adminReverseProxyToken, backend string) {
	//curl http://localhost:8888/memory/functions/list
	//fmt.Println(reverseProxyUrl, adminReverseProxyToken, backend)

	client := resty.New()

	resp, err := client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json; charset=utf-8").
		Get(reverseProxyUrl + "/" + backend + "/functions/list")

	if err != nil {
		fmt.Println("😡", err)
		os.Exit(1)
	} else {
		fmt.Println(resp)
		os.Exit(0)
	}

}
