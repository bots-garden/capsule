// Package tools (ngrok part)
package tools

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

// ActivateNgrokTunnel creates a ngrok tunnel and writes its URL to a file.
//
// It receives a context and a pointer to a fiber.App. It returns nothing.
/*
This code creates an ngrok tunnel and writes the URL to a file named "ngrok.url"
in the same directory as the executable file.
The ActivateNgrokTunnel function takes in a context and a pointer to a Fiber application.
It uses the ngrok package to create the tunnel, and then saves the tunnel URL to a file.
Finally, it sets the Fiber application's listener to the created tunnel.
*/
func ActivateNgrokTunnel(ctx context.Context, app *fiber.App) {
	// Ngrok support: https://ngrok.com
	// https://ngrok.com/blog-post/ngrok-go

	var configHTTPEndPoint config.Tunnel
	var ngrokDomain = GetEnv("NGROK_DOMAIN", "")
	if ngrokDomain != "" {
		configHTTPEndPoint = config.HTTPEndpoint(config.WithDomain(ngrokDomain))
	} else {
		configHTTPEndPoint = config.HTTPEndpoint()
	}

	tun, err := ngrok.Listen(ctx,
		configHTTPEndPoint,
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		log.Println("❌ Error while creating tunnel:", err)
	}

	log.Println("👋 Ngrok tunnel created:", tun.URL())

	ex, err := os.Executable()
	if err != nil {
		log.Fatal("❌ Error after creating tunnel:", err)
	}
	exPath := filepath.Dir(ex)

	f, err := os.Create(exPath + "/ngrok.url")

	if err != nil {
		log.Fatal("❌ Error when creating ngrok.url:", err)

	}

	defer f.Close()

	_, errWrite := f.WriteString(tun.URL())

	if errWrite != nil {
		log.Fatal("❌ Error when writing ngrok.url:", errWrite)
	}

	log.Println("🤚 Ngrok URL:", exPath+"/ngrok.url")

	app.Listener(tun)
}
