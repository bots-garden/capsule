// Package handlers : this one is to start a new Capsule HTTP process
package handlers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/bots-garden/capsule/capsule-http/data"
	"github.com/bots-garden/capsule/capsule-http/tools"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

// StartNewCapsuleHTTP is a Go function that handles HTTP requests
// for starting a capsule.
// ! this a work in progress
// It takes in a pointer to a fiber.Ctx object.
// It returns an error object.
// TODO: protect this route with a middleware
func StartNewCapsuleHTTP(c *fiber.Ctx) error {
	//TODO check if the process exists

	/*
		type CapsuleTask struct {
			FunctionName     string   `json:"name"`
			FunctionRevision string   `json:"revision"`
			Description      string   `json:"description"`
			Path             string   `json:"path"`
			Args             []string `json:"args"`
			Env              []string `json:"env"`
		}
	*/

	capsuleTask := data.CapsuleTask{}
	jsonPostPayloadErr := c.BodyParser(&capsuleTask)
	if jsonPostPayloadErr != nil {
		return jsonPostPayloadErr
	}

	httpPort := tools.GetNewHTTPPort()

	capsuleTask.Args = append(capsuleTask.Args, "-httpPort="+httpPort)

	// TODO: store somewhere the processes that are running (or not)

	fmt.Println("🔷", capsuleTask.Args)
	fmt.Println("🔷", capsuleTask.Env)

	// try without the httport too

	if capsuleTask.Path == "" {
		capsuleTask.Path = "capsule-http" //! had to be installed
	}

	//? How to get the path of the current working directory
	//! default value is "capsule-http" (if installed)
	cmd := &exec.Cmd{
		Path:   capsuleTask.Path,
		Args:   capsuleTask.Args,
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}
	newEnv := append(os.Environ(), capsuleTask.Env...)
	cmd.Env = newEnv

	err := cmd.Start()

	if err != nil {
		log.Println("🚗", err.Error())
	} else {
		log.Println("🚙", cmd.Args)
	}

	//! Save the process
	capsuleRecord := data.CapsuleProcess{
		FunctionName:      capsuleTask.FunctionName,
		FunctionRevision:  capsuleTask.FunctionRevision,
		HTTPPort:          httpPort,
		Description:       capsuleTask.Description,
		CurrentStatus:     0,
		StatusDescription: "",
		CreatedAt:         time.Now(),
		StartedAt:         time.Now(),
		FinishedAt:        time.Now(),
		CancelledAt:       time.Now(),
		FailedAt:          time.Now(),
		CheckedAt:         time.Now(),
		Pid:               cmd.Process.Pid,
		Path:              capsuleTask.Path,
		Args:              capsuleTask.Args,
		Env:               newEnv,
		Cmd:               cmd,
	}
	idOfTheProcess := data.CreateCapsuleProcessRecord(capsuleRecord)

	fmt.Println("🤖 id of the process:", idOfTheProcess)
	return c.Send([]byte(idOfTheProcess))

}

/*
   "path": "./services/capsule/capsule-http",
   "args": [
     "",
     "-wasm=./services/capsule/hello-world.wasm",
     "-httpPort=59746"
   ],
*/

// CallExternalFunction is a Go function that handles calls to an external function.
// ! this a work in progress
// c *fiber.Ctx: a pointer to a fiber context object that contains information about the http request.
// error: returns an error if the external function call fails.
func CallExternalFunction(c *fiber.Ctx) error {

	functionName := c.Params("function_name")
	functionRevision := c.Params("function_revision")
	functionIndex := c.Params("function_index") // ! default index is 0

	if functionRevision == "" {
		functionRevision = "default"
	}
	if functionIndex == "" {
		functionIndex = "0"
	}

	key := functionName + "/" + functionRevision + "/" + functionIndex

	fmt.Println("🌍", key)

	process, err := data.GetCapsuleProcessRecord(key)

	if err != nil {
		//TODO handle error
		fmt.Println("😡", err.Error())
	}

	bodyRequest := c.Body()
	headersRequest := c.GetReqHeaders()
	httpClient := resty.New()

	for key, value := range headersRequest {
		httpClient.SetHeader(key, value)
	}

	// TODO: use an environment variable?
	capsuleDomain := "http://localhost"
	capsuleURI := capsuleDomain + ":" + process.HTTPPort
	strBodyRequest := string(bodyRequest)

	restyHeadersToFiberHeaders := func(resp *resty.Response) {
		for key, value := range resp.Header() {
			c.Response().Header.Set(key, value[0])
		}
	}

	switch what := c.Method(); what {
	case "GET":
		resp, err := httpClient.R().EnableTrace().Get(capsuleURI)

		restyHeadersToFiberHeaders(resp)
		c.Status(resp.StatusCode())

		if err != nil {
			return c.Send([]byte(err.Error()))
		}

		return c.Send([]byte(resp.String()))

	case "POST":
		resp, err := httpClient.R().EnableTrace().SetBody(strBodyRequest).Post(capsuleURI)
		//resp, err := httpClient.R().SetBody(string(bodyRequest)).Post(capsuleURI)

		restyHeadersToFiberHeaders(resp)
		c.Status(resp.StatusCode())

		if err != nil {
			return c.Send([]byte(err.Error()))
		}

		return c.Send([]byte(resp.String()))

	case "PUT":
		resp, err := httpClient.R().EnableTrace().SetBody(strBodyRequest).Put(capsuleURI)
		restyHeadersToFiberHeaders(resp)
		c.Status(resp.StatusCode())

		if err != nil {
			return c.Send([]byte(err.Error()))
		}

		return c.Send([]byte(resp.String()))

	case "DELETE":
		resp, err := httpClient.R().EnableTrace().Delete(capsuleURI)
		restyHeadersToFiberHeaders(resp)
		c.Status(resp.StatusCode())

		if err != nil {
			return c.Send([]byte(err.Error()))
		}

		return c.Send([]byte(resp.String()))

	default:
		return nil // TODO: return something else
	}

}

// 	app.All("/functions/call/:function_name/:function_revision", handlers.CallExternalFunction)
