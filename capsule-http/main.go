// Package main, the next generation of the Capsule project
package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"os/signal"
	"strconv"

	"syscall"
	"time"

	"log"
	"os"

	"github.com/bots-garden/capsule-host-sdk"
	"github.com/bots-garden/capsule/capsule-http/handlers"
	"github.com/bots-garden/capsule/capsule-http/tools"
	"github.com/go-resty/resty/v2"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
)

// CapsuleFlags handles params for the capsule-http command
type CapsuleFlags struct {
	wasm            string // wasm file location
	httpPort        string
	stopAfter       string // stop after a delay if not used
	url             string // to download the wasm file
	authHeaderName  string // if needed for authentication
	authHeaderValue string // if needed for authentication
	crt             string // https (certificate)
	key             string // https (key)
	registry        string // url to the registry
	version         bool
	parentEndpoint  string // url to the parent endpoint (use by faas mode / main capsule process)
	moduleName      string // functionName/revision (use by faas mode only)
	moduleRevision  string // functionName/revision (use by faas mode only)

	// faasToken?
}

//go:embed description.txt
var textVersion []byte

func main() {

	version := string(textVersion)
	args := os.Args[1:]

	if len(args) == 0 {
		log.Println("Capsule needs some args to start.")
		os.Exit(0)
	}
	// Capsule flags
	wasmFilePathPtr := flag.String("wasm", "", "wasm module file path")
	httpPortPtr := flag.String("httpPort", "", "http port")
	stopAfterPtr := flag.String("stopAfter", "", "stop after n seconds if not used")
	wasmFileURLPtr := flag.String("url", "", "url for downloading wasm module file")
	authHeaderNamePtr := flag.String("authHeaderName", "", "header authentication for downloading wasm module file")
	authHeaderValuePtr := flag.String("authHeaderValue", "", "header authentication value for downloading wasm module file")
	registryPtr := flag.String("registry", "", "url of the wasm registry")
	crtPtr := flag.String("crt", "", "certificate")
	keyPtr := flag.String("key", "", "key")
	versionPtr := flag.Bool("version", false, "prints capsule CLI current version")
	parentEndpointPtr := flag.String("parentEndpoint", "", "TBD 🚧/Only for FaaS mode")
	moduleNamePtr := flag.String("moduleName", "", "TBD 🚧TBD 🚧/Only for FaaS mode")
	moduleRevisionPtr := flag.String("moduleRevision", "", "TBD 🚧TBD 🚧/Only for FaaS mode")

	flag.Parse()

	if *versionPtr {
		fmt.Println(version)
		os.Exit(0)
	}

	flags := CapsuleFlags{
		*wasmFilePathPtr,
		*httpPortPtr,
		*stopAfterPtr,
		*wasmFileURLPtr,
		*authHeaderNamePtr,
		*authHeaderValuePtr,
		*crtPtr,
		*keyPtr,
		*registryPtr,
		*versionPtr,
		*parentEndpointPtr,
		*moduleNamePtr,
		*moduleRevisionPtr,
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
		log.Println("❌📝 Error while loading the wasm file", err)
		os.Exit(1)
	}
	handlers.StoreWasmFile(wasmFile)

	// --------------------------------------------
	// Prometheus
	// --------------------------------------------
	// ! this is experimental and subject to change
	//prometheus := fiberprometheus.New("capsule-http:"+httpPort+"|"+version+"("+flags.wasm+")")
	prometheus := fiberprometheus.New("capsule")

	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	// TODO: protect these routes

	// --------------------------------------------
	// Handler to launch a new Capsule process
	// and create a revision for a function
	// --------------------------------------------
	app.Post("/functions/start", handlers.StartNewCapsuleHTTPProcess)

	// Start a new Capsule HTTP process, the shutdown after a delay
	//app.Post("/functions/start/shutdown", handlers.StartNewCapsuleHTTPProcessThenShutdownItAfterDelay)

	// Get the list of processes
	app.Get("/functions/processes", handlers.GetListOfCapsuleHTTPProcesses)

	// TODO: do it with index too?
	// Duplicate a process
	app.Put("/functions/duplicate/:function_name/:function_revision/:new_function_revision", handlers.DuplicateExternalFunction)

	// Stop a process
	app.Delete("/functions/drop/:function_name", handlers.StopAndKillCapsuleHTTPProcess)
	app.Delete("/functions/drop/:function_name/:function_revision", handlers.StopAndKillCapsuleHTTPProcess)
	app.Delete("/functions/drop/:function_name/:function_revision/:function_index", handlers.StopAndKillCapsuleHTTPProcess)

	// --------------------------------------------
	// Handler to the revision of an external
	// function
	// --------------------------------------------
	app.All("/functions/:function_name", handlers.CallExternalFunction)
	app.All("/functions/:function_name/:function_revision", handlers.CallExternalFunction)
	app.All("/functions/:function_name/:function_revision/:function_index", handlers.CallExternalFunction)

	// --------------------------------------------
	// Handler to notify the main capsule process
	// --------------------------------------------
	app.All("/notify/:function_name/:function_revision", handlers.NotifiedMainCapsuleHTTPProcess)
	app.All("/notify/:function_name/:function_revision/:function_index", handlers.NotifiedMainCapsuleHTTPProcess)

	// --------------------------------------------
	// Handler to call the WASM function
	// --------------------------------------------
	app.All("/", handlers.CallWasmFunction)

	// --------------------------------------------
	// Start listening
	// --------------------------------------------
	go func() {

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

			app.ListenTLS(":"+httpPort, flags.crt, flags.key)

		} else {
			log.Println("💊 Capsule", version, "http server is listening on:", httpPort, "🌍")

			if tools.GetEnv("NGROK_AUTH_TOKEN", "") != "" {
				tools.ActivateNgrokTunnel(ctx, app)
			}

			app.Listen(":" + httpPort)

		}
	}()

	go func() {
		// Set a value for the last call
		if flags.stopAfter == "" {
			return
		}
		duration, _ := strconv.ParseFloat(flags.stopAfter, 64)
		handlers.SetLastCall(time.Now())
		for {
			time.Sleep(1 * time.Second)
			if time.Since(handlers.GetLastCall()).Seconds() >= duration {
				stop()
				//log.Println("👋 Bye!")
			}
			//else {
			//	log.Println("🟢 Last call since:", time.Since(handlers.GetLastCall()).Seconds())
			//}
		}

	}()

	// It's only for debugging
	/*
	go func() {
		for {
			time.Sleep(1 * time.Second)
			processes := data.GetAllCapsuleProcessRecords()
			if len(processes) > 0 { // I'm the main process
				for _, p := range processes {
					fmt.Println("📳 ->", p.FunctionName, p.FunctionRevision, p.Index, p.Description, p.StatusDescription)
				}
			}
		}
	}()
	*/

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()

	log.Println("💊 Capsule shutting down...", flags.wasm)

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// flags.parentEndpoint, flags.moduleName and flags.moduleRevision are set only
	// if the process was triggered by another Capsule process with **capsctl**
	// (faas mode)
	log.Println("💊 Capsule stopped", flags.wasm, flags.parentEndpoint, flags.moduleName, flags.moduleRevision)

	// Telling the main process that I'm exiting...
	if flags.parentEndpoint != "" {
		log.Println("Ⓜ️ sending notification", flags.parentEndpoint, flags.moduleName, flags.moduleRevision)
		httpClient := resty.New()
		_, err := httpClient.R().EnableTrace().Get(flags.parentEndpoint+"/notify/"+flags.moduleName+"/"+flags.moduleRevision)
		if err != nil {
			log.Println("❌ Error while sending notification:", err)
		}
	}

}
