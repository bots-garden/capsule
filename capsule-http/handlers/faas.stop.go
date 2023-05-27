// Package handlers : this one is to start a new Capsule HTTP process
package handlers

import (
	"log"
	"os"

	"github.com/bots-garden/capsule/capsule-http/data"

	"github.com/gofiber/fiber/v2"
)


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

	//errProcessToKill := process.Cmd.Process.Kill()
	// Capsule shutting down gracefully
	errProcessToKill := process.Cmd.Process.Signal(os.Interrupt)

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
