// Package handlers : this one is to start a new Capsule HTTP process
package handlers

import (
	"log"

	"github.com/bots-garden/capsule/capsule-http/data"
	"github.com/gofiber/fiber/v2"
)

// NotifiedMainCapsuleHTTPProcess is triggered when
// a Capsule HTTP child process notifies
// the parent process (main Capsule HTTP process)
func NotifiedMainCapsuleHTTPProcess(c *fiber.Ctx) error {

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


	// TODO: handle the error
	process, err := data.GetCapsuleProcessRecord(key)

	if err != nil {
		log.Println("🔴 Error when getting the existing Capsule process:", err.Error())
		c.Status(fiber.StatusInternalServerError)
		return c.Send([]byte(err.Error()))
	}

	process.CurrentStatus = data.Stopped
	process.StatusDescription = data.GetStatusLabel(data.Stopped)

	// Save(Update)
	data.SetCapsuleProcessRecord(process)

	c.Status(fiber.StatusOK)
	return c.Send([]byte(key))
}
