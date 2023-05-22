// Package handlers : this one is to start a new Capsule HTTP process
package handlers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/bots-garden/capsule/capsule-http/tools"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)



// StartNewCapsuleHTTP is a Go function that handles HTTP requests
// for starting a capsule.
// ! this a work in progress
// It takes in a pointer to a fiber.Ctx object.
// It returns an error object.
func StartNewCapsuleHTTP(c *fiber.Ctx) error {
	//TODO check if the process exists

	functionName := c.Params("function_name")
	functionRevision := c.Params("function_revision")
	httpPort := tools.GetHTTPPort()
	// TODO: store somewhere the processes that are running (or not)

	fmt.Println(functionName, functionRevision, httpPort)

	args := []string{
		"",
		"-wasm=./functions/hello-world/hello-world.wasm",
		"-httpPort=" + httpPort,
	}
	// try without the httport too

	//? How to get the path of the current working directory
	cmd := &exec.Cmd{
		Path:   "./capsule-http",
		Args:   args,
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}
	newEnv := append(os.Environ(), []string{"MESSAGE=Hello"}...)
	cmd.Env = newEnv

	err := cmd.Start()

	if err != nil {
		log.Println("🚗", err.Error())
	} else {
		log.Println("🚙", cmd.Args)
	}

	//TODO make an HTTP request to the function
	bodyRequest := c.Body()
	//headersRequest := c.GetReqHeaders()

	httpClient := resty.New()

	/*
		for key, value := range req.Headers {
			httpClient.SetHeader(key, value)
		}
	*/
	time.Sleep(1 * time.Second) // Wait for 1 seconds
	// TODO use health check before launch

	//? how to get the appropriate URI?
	//! eg: how to know we are on https?
	//TODO test request method
	resp, err := httpClient.R().EnableTrace().SetBody(string(bodyRequest)).Post("http://localhost:" + httpPort)



	if err != nil {
		return c.Send([]byte(err.Error()))
	}
	return c.Send([]byte(resp.String()))

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
	fmt.Println(functionName, functionRevision)
	return nil
}
