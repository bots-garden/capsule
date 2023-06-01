// Package handlers : this one is to start a new Capsule HTTP process
package handlers

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/bots-garden/capsule/capsule-http/data"
	"github.com/bots-garden/capsule/capsule-http/tools"
	"github.com/gofiber/fiber/v2"
)

// StartNewCapsuleHTTPProcess is a Go function that handles HTTP requests
// for starting a capsule.
// ! this a work in progress
// It takes in a pointer to a fiber.Ctx object.
// It returns an error object.
func StartNewCapsuleHTTPProcess(c *fiber.Ctx) error {

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

	// TODO: test environment variables

	// ? or use an environment variable?
	if capsuleTask.Path == "" {
		// Default value
		capsuleTask.Path = "capsule-http" //! had to be installed
		// ! this does not work
		// ? how to get the path where I'm installed
	}

	//fmt.Println("🔷", capsuleTask.Path)

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
	capsuleRecord := data.CapsuleProcess{
		FunctionName:     capsuleTask.FunctionName,
		FunctionRevision: capsuleTask.FunctionRevision,
		HTTPPort:         httpPort,
		Description:      capsuleTask.Description,
		CurrentStatus:    data.Started,
		StatusDescription: data.GetStatusLabel(data.Started),
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
	// index is(will be) used with the scaling feature (it's a work in progress 🚧)
	idOfTheProcess, _ := data.CreateCapsuleProcessRecord(capsuleRecord)

	c.Status(fiber.StatusOK)
	return c.Send([]byte(idOfTheProcess))

}
