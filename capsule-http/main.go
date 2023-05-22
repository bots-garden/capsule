// Package main, the next generation of the Capsule project
package main

import (
	"context"
	"flag"
	"fmt"
	"os/signal"

	//"strings"
	"syscall"
	"time"

	"log"
	"os"

	"github.com/bots-garden/capsule-host-sdk"
	"github.com/bots-garden/capsule/capsule-http/handlers"
	"github.com/bots-garden/capsule/capsule-http/tools"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	// go get -u github.com/ansrivas/fiberprometheus/v2
)

// CapsuleFlags handles params for the capsule-http command
type CapsuleFlags struct {
	wasm            string // wasm file location
	httpPort        string
	url             string // to download the wasm file
	authHeaderName  string // if needed for authentication
	authHeaderValue string // if needed for authentication
	crt             string // https (certificate)
	key             string // https (key)
	registry        string // url to the registry
	version         bool
}

func main() {
	version := "v0.3.6 🫐 [blueberries]"
	args := os.Args[1:]

	if len(args) == 0 {
		log.Println("Capsule needs some args to start.")
		os.Exit(0)
	}
	// Capsule flags
	wasmFilePathPtr := flag.String("wasm", "", "wasm module file path")
	httpPortPtr := flag.String("httpPort", "", "http port")
	wasmFileURLPtr := flag.String("url", "", "url for downloading wasm module file")
	authHeaderNamePtr := flag.String("authHeaderName", "", "header authentication for downloading wasm module file")
	authHeaderValuePtr := flag.String("authHeaderValue", "", "header authentication value for downloading wasm module file")
	registryPtr := flag.String("registry", "", "url of the wasm registry")
	crtPtr := flag.String("crt", "", "certificate")
	keyPtr := flag.String("key", "", "key")
	versionPtr := flag.Bool("version", false, "prints capsule CLI current version")

	flag.Parse()

	if *versionPtr {
		fmt.Println(version)
		os.Exit(0)
	}

	flags := CapsuleFlags{
		*wasmFilePathPtr,
		*httpPortPtr,
		*wasmFileURLPtr,
		*authHeaderNamePtr,
		*authHeaderValuePtr,
		*crtPtr,
		*keyPtr,
		*registryPtr,
		*versionPtr,
	}

	// Create context that listens for the interrupt signal from the OS.
	// This context will be used for function calls.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	handlers.StoreContext(ctx)

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		//DisableKeepalive:      true,
		//Concurrency:           100000,
	})

	// Create a new WebAssembly Runtime.
	runtime := capsule.GetRuntime(ctx)

	// START: host functions
	// Get the builder and load the default host functions
	builder := capsule.GetBuilder(runtime)

	// * Add your host functions here
	// 🏠
	// * End of of you hostfunction

	// Instantiate builder and default host functions
	_, err := builder.Instantiate(ctx)
	if err != nil {
		log.Println("❌ Error with env module and host function(s):", err)
		os.Exit(1)
	}
	// END: host functions

	// This closes everything this Runtime created.
	defer runtime.Close(ctx)

	handlers.StoreRuntime(runtime)

	// -----------------------------------
	// Load the WebAssembly module
	// -----------------------------------
	wasmFile, err := tools.GetWasmFile(flags.wasm, flags.url, flags.authHeaderName, flags.authHeaderValue)
	if err != nil {
		log.Println("❌ Error while loading the wasm file", err)
		os.Exit(1)
	}
	handlers.StoreWasmFile(wasmFile)

	// -----------------------------------
	// Prometheus
	// -----------------------------------
	//prometheus := fiberprometheus.New("capsule-http:"+httpPort+"|"+version+"("+flags.wasm+")")
	prometheus := fiberprometheus.New("capsule")

	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	// ----------------------------------------
	// Handler to launch a new Capsule process
	// and create a revision for a function
	//
	// TODO: protect this route
	// ----------------------------------------
	app.All("/functions/start", handlers.StartNewCapsuleHTTP)


	// ----------------------------------------
	// Handler to the revision of an external
	// function
	// ----------------------------------------
	app.All("/functions/call/:function_name/:function_revision", handlers.CallExternalFunction)



	// -----------------------------------
	// Handler to call the WASM function
	// -----------------------------------
	// TODO: protect routes
	// TODO: externalise the handler
	// TODO: create helpers to simplify the code
	app.All("/", handlers.CallWasmFunction)

	// -----------------------------------
	// Start listening
	// -----------------------------------
	go func() {

		var httpPort string

		if flags.httpPort == "" {
			httpPort = tools.GetHTTPPort()
		} else {
			httpPort = flags.httpPort
		}

		if flags.crt != "" {
			// certs/capsule.local.crt
			// certs/capsule.local.key
			log.Println("💊 Capsule", version, "http server is listening on:", httpPort, "🔐🌍")

			app.ListenTLS(":"+httpPort, flags.crt, flags.key)

		} else {
			log.Println("💊 Capsule", version, "http server is listening on:", httpPort, "🌍")

			if tools.GetEnv("NGROK_AUTH_TOKEN", "") != "" {
				tools.ActivateNgrokTunnel(ctx, app)
			}

			app.Listen(":" + httpPort)

		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("💊 Capsule shutting down gracefully...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("💊 Capsule exiting...")
}
