package main

import (
	"log"
	"os"

	"github.com/bots-garden/capsule/capsule-http/handlers"
	"github.com/bots-garden/capsule/capsule-http/tools"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/skip"
)

// --------------------------------------------
// ! This is the FAAS mode of Capsule HTTP 🚀
// --------------------------------------------

var capsuleFaasToken string

func checkToken(c *fiber.Ctx) bool {
	predicate := c.Get("CAPSULE_FAAS_TOKEN") != capsuleFaasToken
	if predicate == true {
		log.Println("🔴🤚 FAAS mode activated, you need to set CAPSULE_FAAS_TOKEN!")
		//c.Status(fiber.StatusUnauthorized)
	}
	return predicate
}

// StartFaasMode -> start the FaaS mode of Capsule
func StartFaasMode(app *fiber.App) error {
	capsuleFaasToken = tools.GetEnv("CAPSULE_FAAS_TOKEN", "")
	ex, err := os.Executable()
	if err != nil {
		return err
	}

	// 0.3.9 handlers.SetMainCapsuleTaskPath(os.Args[0])
	// change for 0.4.0:
	handlers.SetMainCapsuleTaskPath(ex)

	log.Println("🚀 faas mode activated!", "["+ex+"]", handlers.GetMainCapsuleTaskPath())

	defineStartRoute(app)
	defineProcessListRoute(app)
	defineDuplicateProcessRoute(app)
	defineStopProcessRoutes(app)
	defineCallRevisionRoutes(app)
	defineHealthRoutes(app)
	defineNotifiyRoutes(app)

	return nil

}

func defineStartRoute(app *fiber.App) {
	// --------------------------------------------
	// Handler to launch a new Capsule process
	// and create a revision for a function
	// --------------------------------------------
	app.Post("/functions/start", skip.New(handlers.StartNewCapsuleHTTPProcess, checkToken))
}

func defineProcessListRoute(app *fiber.App) {
	// Get the list of processes
	app.Get("/functions/processes", skip.New(handlers.GetListOfCapsuleHTTPProcesses, checkToken))
}

func defineDuplicateProcessRoute(app *fiber.App) {
	// ???: do it with index too?
	// Duplicate a process
	app.Put("/functions/duplicate/:function_name/:function_revision/:new_function_revision", skip.New(handlers.DuplicateExternalFunction, checkToken))
}

func defineStopProcessRoutes(app *fiber.App) {
	// Stop a process
	app.Delete("/functions/drop/:function_name", skip.New(handlers.StopAndKillCapsuleHTTPProcess, checkToken))
	app.Delete("/functions/drop/:function_name/:function_revision", skip.New(handlers.StopAndKillCapsuleHTTPProcess, checkToken))
	app.Delete("/functions/drop/:function_name/:function_revision/:function_index", skip.New(handlers.StopAndKillCapsuleHTTPProcess, checkToken))
}

func defineCallRevisionRoutes(app *fiber.App) {
	// --------------------------------------------
	// Handler to call the revision of an external
	// function (module)
	// --------------------------------------------
	app.All("/functions/:function_name", handlers.CallExternalFunction)
	app.All("/functions/:function_name/:function_revision", handlers.CallExternalFunction)
	app.All("/functions/:function_name/:function_revision/:function_index", handlers.CallExternalFunction)

}

func defineHealthRoutes(app *fiber.App) {
	// --------------------------------------------
	// Handler to check the health of the
	// Capsule HTTP process
	// --------------------------------------------
	app.Get("/functions/health/:function_name", handlers.CallExternalFunctionHealthCheck)
	app.Get("/functions/health/:function_name/:function_revision", handlers.CallExternalFunctionHealthCheck)
	app.Get("/functions/health/:function_name/:function_revision/:function_index", handlers.CallExternalFunctionHealthCheck)
}

func defineNotifiyRoutes(app *fiber.App) {
	// --------------------------------------------
	// Handler to notify the main capsule process
	// --------------------------------------------
	app.All("/notify/:function_name/:function_revision", handlers.NotifiedMainCapsuleHTTPProcess)
	app.All("/notify/:function_name/:function_revision/:function_index", handlers.NotifiedMainCapsuleHTTPProcess)
}
