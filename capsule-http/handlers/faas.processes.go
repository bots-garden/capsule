// Package handlers : this one is to start a new Capsule HTTP process
package handlers

import (
	"log"

	"github.com/bots-garden/capsule/capsule-http/data"
	"github.com/gofiber/fiber/v2"
)


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
