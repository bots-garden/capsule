// Package main, the next generation of the Capsule project
package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"os/signal"

	"syscall"
	"time"

	"log"
	"os"

	"github.com/bots-garden/capsule-host-sdk"
	"github.com/bots-garden/capsule/capsule-http/handlers"
	"github.com/bots-garden/capsule/capsule-http/tools"
	"github.com/go-resty/resty/v2"

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
	faas            bool

	// faasToken?
}

func main() {
	
	//version := string(textVersion)
	version := tools.GetVersion()

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
	faasPtr := flag.Bool("faas", false, "TBD 🚧TBD 🚧/Only for FaaS mode")

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
		*faasPtr,
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

	// -----------------------------------
	// START: host functions
	// -----------------------------------
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
	// -----------------------------------
	// END: host functions
	// -----------------------------------

	// This closes everything this Runtime created.
	defer runtime.Close(ctx)

	handlers.StoreRuntime(runtime)

	// -----------------------------------
	// Load the WebAssembly module
	// -----------------------------------
	// wasmFile []byte is global to be accessible
	// from CallOnStart and CallOnStop

	wasmFile, err := LoadWasmFile(ctx, flags, runtime)
	if err != nil {
		os.Exit(1)
	}
	// Call only once OnStart wasm module method
	// Onstart is an exported function
	if wasmFile != nil {
		// with FaaS mode, the wasm file could be empty
		mod, err := handlers.GetModule(ctx, wasmFile)
		if err != nil {
			//TODO: display error message
			log.Println("❌ [OnStart] Error with the module instance", err)
			os.Exit(1)
		}
		capsule.CallOnStart(ctx, mod, wasmFile)
	}
	
	if flags.faas == true {
		// -----------------------------------
		// Start FaaS mode
		// and define routes
		// -----------------------------------
		err := StartFaasMode(app)
		if err != nil {
			os.Exit(1)
		}
	}

	// --------------------------------------------
	// Handler to call the WASM function
	// --------------------------------------------
	defineMainCapsuleProcessRoutes(app, flags)


	// --------------------------------------------
	// Start listening (HTTP Server)
	// --------------------------------------------
	go func() {

		err := HTTPListening(ctx, flags, version, app)
		if err != nil {
			os.Exit(1)
		}

	}()

	go func() {
		/*
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
			}
		}
		*/
		stopProcess := ShouldStopAfterDelay(flags)
		if stopProcess {
			stop()
		}

	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Call OnStop (when Ctr+C)
	if wasmFile != nil {
		// Call only once CallOnStop wasm module method
		mod, err := handlers.GetModule(ctx, wasmFile)
		if err != nil {
			//TODO: display error message
			log.Println("❌ [OnStop] Error with the module instance", err)
			os.Exit(1)
		}
		capsule.CallOnStop(ctx, mod, wasmFile)

	}

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()

	log.Println("💊 Capsule shutting down...", flags.wasm)

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if flags.faas == true {

		// flags.parentEndpoint, flags.moduleName and flags.moduleRevision are set only
		// if the process was triggered by another Capsule process with **capsctl**
		// (faas mode)
		log.Println("💊 Capsule stopped", flags.wasm, flags.parentEndpoint, flags.moduleName, flags.moduleRevision)

		// Telling the main process that I'm exiting...
		if flags.parentEndpoint != "" {
			log.Println("Ⓜ️ sending notification", flags.parentEndpoint, flags.moduleName, flags.moduleRevision)
			httpClient := resty.New()
			_, err := httpClient.R().EnableTrace().Get(flags.parentEndpoint + "/notify/" + flags.moduleName + "/" + flags.moduleRevision)
			if err != nil {
				log.Println("❌ Error while sending notification:", err)
			}
		}
	} else {
		log.Println("💊 Capsule stopped", flags.wasm)
	}

}
