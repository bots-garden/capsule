package main

import (
	"github.com/bots-garden/capsule/capsule-http/handlers"
	"github.com/gofiber/fiber/v2"
)

func defineMainCapsuleProcessRoutes(app *fiber.App, flags CapsuleFlags) {

	// TODO: Question: route protection or not?
	app.Get("/metrics", handlers.CallWasmFunctionMetrics)
	app.Get("/health", handlers.CallWasmFunctionHealthCheck)

	if flags.faas == true && flags.wasm == "" {
		// if "/*"
		// first: try a handler similar to handlers.CallExternalFunction (faas.call.go)
		// and function name is index.page
		app.All("/*", handlers.CallExternalIndexPageFunction)

	} else { 
		// "normal" mode or FaaS mode loading a wasm module in the same process
		app.All("/*", handlers.CallWasmFunction)
	}

}