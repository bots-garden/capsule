// Package handlers : this one is to start a new Capsule HTTP process
package handlers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
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

	// Read the body of the request
	capsuleTask := data.CapsuleTask{}
	jsonPostPayloadErr := c.BodyParser(&capsuleTask)
	if jsonPostPayloadErr != nil {
		return jsonPostPayloadErr
	}

	// Get a new HTTP Port
	httpPort := tools.GetNewHTTPPort()

	// Update the arguments before starting a new capsule process
	capsuleTask.Args = append(capsuleTask.Args, "-httpPort="+httpPort)

	// ! this a work in progress 🚧
	fmt.Println("🔷", capsuleTask.Args)
	fmt.Println("🔷", capsuleTask.Env)

	// ? or use an environment variable?
	if capsuleTask.Path == "" {
		// Default value
		capsuleTask.Path = "capsule-http" //! had to be installed
	}

	cmd := &exec.Cmd{
		Path:   capsuleTask.Path,
		Args:   capsuleTask.Args,
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}
	newEnv := append(os.Environ(), capsuleTask.Env...)
	cmd.Env = newEnv

	/* Example:
	   "path": "./services/capsule/capsule-http",
	   "args": [
	     "",
	     "-wasm=./services/capsule/hello-world.wasm",
	     "-httpPort=59746"
	   ],
	*/
	err := cmd.Start()

	if err != nil {
		log.Println("🔴 Error when starting a new Capsule process:", err.Error())
		c.Status(fiber.StatusInternalServerError)
		return c.Send([]byte(err.Error()))
	}

	// Create a new record of the Capsule Process
	// TODO: Save the process (to be implemented)
	capsuleRecord := data.CapsuleProcess{
		FunctionName:     capsuleTask.FunctionName,
		FunctionRevision: capsuleTask.FunctionRevision,
		HTTPPort:         httpPort,
		Description:      capsuleTask.Description,
		CurrentStatus:    data.Started,
		CreatedAt:        time.Now(),
		StartedAt:        time.Now(),
		FinishedAt:       time.Now(),
		CancelledAt:      time.Now(),
		FailedAt:         time.Now(),
		CheckedAt:        time.Now(),
		Pid:              cmd.Process.Pid,
		Path:             capsuleTask.Path,
		Args:             capsuleTask.Args,
		Env:              newEnv,
		Cmd:              cmd,
	}
	idOfTheProcess := data.CreateCapsuleProcessRecord(capsuleRecord)

	c.Status(fiber.StatusOK)
	return c.Send([]byte(idOfTheProcess))

}

// GetListOfCapsuleHTTPProcesses retrieves the list of external Capsule processes
// and sends it as a JSON response.
//
// c *fiber.Ctx: A pointer to the fiber context.
// error: An error indicating if there was a problem with the retrieval of the data.
// error: An error indicating if there was a problem with sending the JSON response.
func GetListOfCapsuleHTTPProcesses(c *fiber.Ctx) error {

	// app.Get("/functions/processes", handlers.GetListOfCapsuleHTTPProcesses)

	jsonProcesses, err := data.GetJSONCapsuleProcesses()
	if err != nil {
		log.Println("🔴 Error when getting the list of external Capsule processes:", err.Error())
		c.Status(fiber.StatusInternalServerError)
		return c.Send([]byte(err.Error()))
	}

	c.Status(fiber.StatusOK)
	return c.Send(jsonProcesses)
}

// StopCapsuleHTTPProcess stops the HTTP process for the given Fiber context.
//
// c: a pointer to a fiber.Ctx object representing the context of the HTTP request.
// error: an error object that indicates whether an error occurred during the function execution.
func StopCapsuleHTTPProcess(c *fiber.Ctx) error {

	/*
		app.Delete("/functions/stop/:function_name", handlers.StopCapsuleHTTPProcess)
		app.Delete("/functions/stop/:function_name/:function_revision", handlers.StopCapsuleHTTPProcess)
		app.Delete("/functions/stop/:function_name/:function_revision/:function_index", handlers.StopCapsuleHTTPProcess)
	*/

	// linked to a function + revision + index
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
	
	process, err := data.GetCapsuleProcessRecord(key)

	if err != nil {
		log.Println("🔴 Error when calling the external Capsule process:", err.Error())
		c.Status(fiber.StatusInternalServerError)
		return c.Send([]byte(err.Error()))
	}

	errProcessToKill := process.Cmd.Process.Kill()
	if errProcessToKill != nil {
		process.CurrentStatus = data.Stucked
		data.SetCapsuleProcessRecord(process)

		log.Println("🔴 Error when killing the external Capsule process:", err.Error())
		c.Status(fiber.StatusInternalServerError)
		return c.Send([]byte(errProcessToKill.Error()))
	}

	process.CurrentStatus = data.Killed
	deletedKey := data.SetCapsuleProcessRecord(process)

	//? should I delete the record?
	//data.DeleteCapsuleProcessRecord(key)

	c.Status(fiber.StatusOK)
	return c.Send([]byte(deletedKey))
}

// CallExternalFunction is a Go function that handles calls to an external function.
// ! this a work in progress
// c *fiber.Ctx: a pointer to a fiber context object that contains information about the http request.
// error: returns an error if the external function call fails.
func CallExternalFunction(c *fiber.Ctx) error {

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

	//fmt.Println("🌍", key)

	process, err := data.GetCapsuleProcessRecord(key)

	// TODO try another index if one does not exist

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
		return c.Send([]byte("method not supported"))
	}

}

// 	app.All("/functions/call/:function_name/:function_revision", handlers.CallExternalFunction)


// DuplicateExternalFunction duplicates a given function and returns its new function name.
//
// c: a fiber context that contains information about the request.
// Returns an error in case of an error.
//
// ! this a work in progress
// The handler duplicate an existing function record with a different revision name
// With the same process (it's a pointer to the same process).
func DuplicateExternalFunction(c *fiber.Ctx) error {

	functionName := c.Params("function_name")
	functionRevision := c.Params("function_revision")
	functionIndex := "0"
	//functionIndex := c.Params("function_index") // ! default index is 0
	// TODO: check if other indexes exist and do it for every index
	newFunctionRevision := c.Params("new_function_revision")

	processToCopy, err := data.GetCapsuleProcessRecord(functionName + "/" + functionRevision + "/" + functionIndex)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.Send([]byte(err.Error()))
	}
	index, err := strconv.Atoi(functionIndex)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.Send([]byte(err.Error()))
	}

	// Create a new record with the same process (PID) but with different revision
	newProcess := data.DuplicateProcess(functionName, newFunctionRevision, index, processToCopy)

	c.Status(fiber.StatusOK)
	return c.Send([]byte(newProcess.FunctionName))

}

// "/functions/duplicate/hello-world/default/saved"
// app.Put("/functions/duplicate/:function_name/:function_revision/:new_function_revision", handlers.DuplicateExternalFunction)

