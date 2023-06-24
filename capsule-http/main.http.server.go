package main

import (
	"context"
	"log"

	"github.com/bots-garden/capsule/capsule-http/tools"
	"github.com/gofiber/fiber/v2"
)

// HTTPListening -> start the HTTP server
func HTTPListening(ctx context.Context, flags CapsuleFlags, version string, app *fiber.App) error {

	var httpPort string

	if flags.httpPort == "" {
		httpPort = tools.GetNewHTTPPort()
	} else {
		httpPort = flags.httpPort
	}

	log.Println("📦 wasm module loaded:", flags.wasm)

	if flags.crt != "" {
		// certs/capsule.local.crt
		// certs/capsule.local.key
		log.Println("💊 Capsule", version, "http server is listening on:", httpPort, "🔐🌍")

		err := app.ListenTLS(":"+httpPort, flags.crt, flags.key)
		if err != nil {
			return err
		}
		return nil

	} 

	log.Println("💊 Capsule", version, "http server is listening on:", httpPort, "🌍")

	if tools.GetEnv("NGROK_AUTH_TOKEN", "") != "" {
		tools.ActivateNgrokTunnel(ctx, app)
	}

	err := app.Listen(":" + httpPort)
	if err != nil {
		return err
	}
	return nil

}