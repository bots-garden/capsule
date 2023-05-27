// Package handlers : this one is to start a new Capsule HTTP process
package handlers

import (
	"strconv"

	"github.com/bots-garden/capsule/capsule-http/data"
	"github.com/gofiber/fiber/v2"
)


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
	newProcess := data.DuplicateProcessWithNewRevision(functionName, newFunctionRevision, index, processToCopy)

	c.Status(fiber.StatusOK)
	return c.Send([]byte(newProcess.FunctionName))

}

// "/functions/duplicate/hello-world/default/saved"
// app.Put("/functions/duplicate/:function_name/:function_revision/:new_function_revision", handlers.DuplicateExternalFunction)
