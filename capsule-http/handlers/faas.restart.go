// Package handlers : this one is to start a new Capsule HTTP process
package handlers

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/bots-garden/capsule/capsule-http/data"
	"github.com/gofiber/fiber/v2"
)


// ReStartCapsuleHTTPProcess restarts an existing stopped process
// Called by CallExternalFunction
func ReStartCapsuleHTTPProcess(c *fiber.Ctx) error {
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
		log.Println("🔴 Error when getting the existing Capsule process:", err.Error())
		c.Status(fiber.StatusInternalServerError)
		return c.Send([]byte(err.Error()))
	}

	cmd := &exec.Cmd{
		Path:   process.Path,
		Args:   process.Args,
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}
	newEnv := append(os.Environ(), process.Env...)
	cmd.Env = newEnv

	err = cmd.Start()

	if err != nil {
		log.Println("🔴 Error when starting a new Capsule process:", err.Error())
		c.Status(fiber.StatusInternalServerError)
		return c.Send([]byte(err.Error()))
	}

	process.StartedAt = time.Now()

	process.CurrentStatus = data.Started
	process.StatusDescription = data.GetStatusLabel(data.Started)

	// Save(Update)
	idOfTheProcess := data.SetCapsuleProcessRecord(process)

	c.Status(fiber.StatusOK)
	return c.Send([]byte(idOfTheProcess))
}

