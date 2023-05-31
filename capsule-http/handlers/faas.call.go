// Package handlers : this one is to start a new Capsule HTTP process
package handlers

import (
	"log"
	"strings"
	"time"

	"github.com/bots-garden/capsule/capsule-http/data"
	"github.com/bots-garden/capsule/capsule-http/tools"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)


// CallExternalFunction calls the WASM module of another Capsule HTTP process
// ! this a work in progress
// c *fiber.Ctx: a pointer to a fiber context object that contains information about the http request.
// error: returns an error if the external function call fails.
func CallExternalFunction(c *fiber.Ctx) error {
	// register the last call
	SetLastCall(time.Now())

	/*
		app.All("/functions/:function_name", handlers.CallExternalFunction)
		app.All("/functions/:function_name/:function_revision", handlers.CallExternalFunction)
		app.All("/functions/:function_name/:function_revision/:function_index", handlers.CallExternalFunction)
	*/

	functionName := c.Params("function_name")
	functionRevision := c.Params("function_revision")
	functionIndex := c.Params("function_index") // ! default index is 0

	if functionRevision == "" {
		functionRevision = "default"
	}
	if functionIndex == "" {
		functionIndex = "0" // ! default index is 0
	}

	// the unique key to identify a Capsule Process
	key := functionName + "/" + functionRevision + "/" + functionIndex

	//log.Println("🤔 key:", key) //!DEBUG

	process, err := data.GetCapsuleProcessRecord(key)

	// TODO try another index if one does not exist
	// TODO or restart the process

	if err != nil {
		log.Println("🔴 Error when calling the external Capsule process:", err.Error())
		return c.Send([]byte(err.Error()))
	}

	bodyRequest := c.Body()
	headersRequest := c.GetReqHeaders()
	httpClient := resty.New()

	for key, value := range headersRequest {
		httpClient.SetHeader(key, value)
	}

	capsuleDomain := tools.GetEnv("CAPSULE_DOMAIN", c.Protocol()+"://"+c.IP()) // ! temporary solution... Or not
	//log.Println("🤔 capsuleDomain:", capsuleDomain) //!DEBUG
	
	capsuleURI := capsuleDomain + ":" + process.HTTPPort
	//log.Println("🤔 capsuleURI:", capsuleURI) //!DEBUG
	
	strBodyRequest := string(bodyRequest)

	restyHeadersToFiberHeaders := func(resp *resty.Response) {
		for key, value := range resp.Header() {
			c.Response().Header.Set(key, value[0])
		}
	}

	retryToConnect := func(freq func() (*resty.Response, error)) (*resty.Response, error) {
		// Retry the HTTP request
		var resp *resty.Response
		// Wait until the process is running/ready
		for i := 1; i < 5; i++ {
			time.Sleep(time.Millisecond * 500) // ? parameter or fixed?
			resp, err = freq()
			if err == nil {
				log.Println("🟢 Successfully restarted the Capsule process")
				break
			}
		}
		return resp, err
	}

	switch what := c.Method(); what {
	case "GET":
		getRequest := func() (*resty.Response, error) {
			resp, err := httpClient.R().EnableTrace().Get(capsuleURI)
			restyHeadersToFiberHeaders(resp)
			c.Status(resp.StatusCode())
			return resp, err
		}
		resp, err := getRequest()

		if err != nil { // That means that the Capsule process is not running
			if strings.Contains(err.Error(), "connection refused") {
				// Try to restart the process (in case it already exists in the process list)
				errStart := ReStartCapsuleHTTPProcess(c)
				if errStart != nil {
					return c.Send([]byte(err.Error()))
				}
				// Retry the HTTP request
				// Wait until the process is running/ready
				resp, err := retryToConnect(getRequest)

				if err != nil {
					// Probably the procesess never existed before
					return c.Send([]byte(err.Error()))
				}
				// Else return the response
				return c.Send([]byte(resp.String()))
			}

			return c.Send([]byte(err.Error()))
		}

		return c.Send([]byte(resp.String()))

	case "POST":
		//log.Println("🤔 POST REQUEST") //!DEBUG

		postRequest := func() (*resty.Response, error) {
			resp, err := httpClient.R().EnableTrace().SetBody(strBodyRequest).Post(capsuleURI)
			restyHeadersToFiberHeaders(resp)
			c.Status(resp.StatusCode())
			return resp, err
		}
		resp, err := postRequest()

		if err != nil { // That means that the Capsule process is not running
			if strings.Contains(err.Error(), "connection refused") {
				// Try to restart the process (in case it already exists in the process list)
				errStart := ReStartCapsuleHTTPProcess(c)
				if errStart != nil {
					return c.Send([]byte(err.Error()))
				}
				// Retry the HTTP request
				// Wait until the process is running/ready
				resp, err := retryToConnect(postRequest)

				if err != nil {
					// Probably the procesess never existed before
					return c.Send([]byte(err.Error()))
				}
				// Else return the response
				return c.Send([]byte(resp.String()))

			}
			return c.Send([]byte(err.Error()))
		}

		return c.Send([]byte(resp.String()))

	case "PUT":
		putRequest := func() (*resty.Response, error) {
			resp, err := httpClient.R().EnableTrace().SetBody(strBodyRequest).Put(capsuleURI)
			restyHeadersToFiberHeaders(resp)
			c.Status(resp.StatusCode())
			return resp, err
		}
		resp, err := putRequest()

		if err != nil { // That means that the Capsule process is not running
			if strings.Contains(err.Error(), "connection refused") {
				// Try to restart the process (in case it already exists in the process list)
				errStart := ReStartCapsuleHTTPProcess(c)
				if errStart != nil {
					return c.Send([]byte(err.Error()))
				}
				// Retry the HTTP request
				// Wait until the process is running/ready
				resp, err := retryToConnect(putRequest)

				if err != nil {
					// Probably the procesess never existed before
					return c.Send([]byte(err.Error()))
				}
				// Else return the response
				return c.Send([]byte(resp.String()))

			}
			return c.Send([]byte(err.Error()))
		}

		return c.Send([]byte(resp.String()))

	case "DELETE":
		deleteRequest := func() (*resty.Response, error) {
			resp, err := httpClient.R().EnableTrace().Delete(capsuleURI)
			restyHeadersToFiberHeaders(resp)
			c.Status(resp.StatusCode())
			return resp, err
		}
		resp, err := deleteRequest()

		if err != nil { // That means that the Capsule process is not running
			if strings.Contains(err.Error(), "connection refused") {
				// Try to restart the process (in case it already exists in the process list)
				errStart := ReStartCapsuleHTTPProcess(c)
				if errStart != nil {
					return c.Send([]byte(err.Error()))
				}
				// Retry the HTTP request
				// Wait until the process is running/ready
				resp, err := retryToConnect(deleteRequest)

				if err != nil {
					// Probably the procesess never existed before
					return c.Send([]byte(err.Error()))
				}
				// Else return the response
				return c.Send([]byte(resp.String()))

			}
			return c.Send([]byte(err.Error()))
		}

		return c.Send([]byte(resp.String()))

	default:
		return c.Send([]byte("method not supported"))
	}

}
