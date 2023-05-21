// Package handlers : this one is to start a new Capsule HTTP process
package handlers

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var runningCapsules sync.Map

// StartCapsuleHTTP is a Go function that handles HTTP requests
// for starting a capsule.
//
// It takes in a pointer to a fiber.Ctx object.
// It returns an error object.
func StartCapsuleHTTP(c *fiber.Ctx) error {

	functionName := c.Params("function_name")
	functionRevision := c.Params("function_revision") // not sure if I will use it
	// TODO: store somewhere the processes that are running (or not)
	fmt.Println(functionName, functionRevision)
	
	return nil
}
