// Package main, the next generation of the Capsule project
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os/signal"
	"strings"

	//"strings"
	"syscall"
	"time"

	"log"
	"net/http"
	"os"

	"github.com/bots-garden/capsule-host-sdk"
	"github.com/bots-garden/capsule-host-sdk/models"
	"github.com/bots-garden/capsule/capsule-http/tools"

	"github.com/gofiber/fiber/v2"
	//"github.com/minio/minio-go/v7"
	//"github.com/minio/minio-go/v7/pkg/credentials"
)

// CapsuleFlags handles params for the capsule-http command
type CapsuleFlags struct {
	wasm     string // wasm file location
	httpPort string
	url      string // to download the wasm file
	crt      string // https (certificate)
	key      string // https (key)
	registry string // url to the registry
	version  bool
}

func main() {
	version := "v0.3.5 🍓 [strawberry]"
	args := os.Args[1:]

	if len(args) == 0 {
		log.Println("Capsule needs some args to start.")
		os.Exit(0)
	}
	// Capsule flags
	wasmFilePathPtr := flag.String("wasm", "", "wasm module file path")
	httpPortPtr := flag.String("httpPort", "8080", "http port")
	wasmFileURLPtr := flag.String("url", "", "url for downloading wasm module file")
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
		*crtPtr,
		*keyPtr,
		*registryPtr,
		*versionPtr,
	}

	// Create context that listens for the interrupt signal from the OS.
	// This context will be used for function calls.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	//! --- THIS IS A TEST ---
	/*
		- https://min.io/docs/minio/linux/developers/go/API.html
		- https://github.com/minio/minio-go
		endpoint := "play.min.io"
		accessKeyID := "Q3AM3UQ867SPQQA43P2F"
		secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
		useSSL := true

		// Initialize minio client object.
		minioClient, errMinio := minio.New(endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: useSSL,
		})
		if errMinio != nil {
			log.Fatalln(errMinio)
		}

		log.Printf("%#v\n", minioClient) // minioClient is now setup
	*/

	//! --- THIS IS A TEST ---

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

	// Add your host functions here
	// 🏠
	// End of of you hostfunction

	// Instantiate builder and default host functions
	_, err := builder.Instantiate(ctx)
	if err != nil {
		log.Println("❌ Error with env module and host function(s):", err)
		os.Exit(1)
	}
	// END: host functions

	// This closes everything this Runtime created.
	defer runtime.Close(ctx)

	// Load the WebAssembly module
	wasmFile, err := tools.GetWasmFile(flags.wasm, flags.url)
	if err != nil {
		log.Println("❌ Error while loading the wasm file", err)
		os.Exit(1)
	}

	// TODO: POST and GET (eg: html) (even DELETE and PUT)
	// TODO: protect routes
	// TODO: systematic load tests
	app.All("/", func(c *fiber.Ctx) error {
		mod, err := runtime.Instantiate(ctx, wasmFile)
		if err != nil {
			log.Println("❌ Error with the module instance", err)
			c.Status(http.StatusInternalServerError) // .🤔
			return c.SendString(err.Error())
		}

		// Get the reference to the WebAssembly function: "callHandle"
		// callHandle is exported by the Capsule plugin
		handleFunction := capsule.GetHandleHTTP(mod)

		// build headers JSON string
		var headers []string
		for field, value := range c.GetReqHeaders() {
			headers = append(headers, `"`+field+`":"`+value+`"`)
		}
		headersStr := strings.Join(headers[:], ",")

		requestParam := models.Request{
			Body: string(c.Body()),
			//JSONBody: string(c.Body()), //! to use in the future
			//TextBody: string(c.Body()), //! to use in the future
			URI:     c.Request().URI().String(),
			Method:  c.Method(),
			Headers: headersStr,
		}

		JSONData, err := json.Marshal(requestParam)

		if err != nil {
			log.Println("❌ Error when reading the request parameter", err)
			c.Status(http.StatusInternalServerError) // .🤔
			return c.SendString(err.Error())
		}

		JSONDataPos, JSONDataSize, err := capsule.CopyDataToMemory(ctx, mod, JSONData)
		if err != nil {
			log.Println("❌ Error when copying data to memory", err)
			c.Status(http.StatusInternalServerError) // .🤔
			return c.SendString(err.Error())
		}

		// Now, we can call "callHandleHTTP"
		// the result type is []uint64
		result, err := handleFunction.Call(ctx,
			JSONDataPos, JSONDataSize)
		if err != nil {
			log.Println("❌ Error when calling callHandleHTTP", err)
			c.Status(http.StatusInternalServerError) // .🤔
			return c.SendString(err.Error())
		}

		responsePos, responseSize := capsule.UnPackPosSize(result[0])

		responseBuffer, err := capsule.ReadDataFromMemory(mod, responsePos, responseSize)
		if err != nil {
			log.Println("❌ Error when reading the memory", err)
			c.Status(http.StatusInternalServerError) // .🤔
			return c.SendString(err.Error())
		}

		// TODO: ReadStringFromMemory, ReadBytesFromMemory...

		responseFromWasmGuest, err := capsule.Result(responseBuffer)
		if err != nil {
			log.Println("❌ Error when getting the Result", err)
			c.Status(http.StatusInternalServerError) // .🤔
			return c.SendString(err.Error())
		}

		//fmt.Println("🚧", string(responseFromWasmGuest))
		// TODO: try unmarshaling with fastjson

		// unmarshal the response
		var response models.Response
		errMarshal := json.Unmarshal(responseFromWasmGuest, &response)
		if errMarshal != nil {
			log.Println("❌ Error when unmarshal the response", errMarshal)
			c.Status(http.StatusInternalServerError) // .🤔
			return c.SendString(errMarshal.Error())
		}

		c.Status(response.StatusCode)

		// set headers
		for key, value := range response.Headers {
			c.Set(key, value)
		}

		//fmt.Println("🟣 JSONBody", response.JSONBody)
		//fmt.Println("🟣 TextBody", response.TextBody)

		if len(response.TextBody) > 0 {
			// send text body
			return c.SendString(response.TextBody)
		}
		// send JSON body
		jsonStr, err := json.Marshal(response.JSONBody)
		if err != nil {
			log.Println("❌ Error when marshal the body", err)
			c.Status(http.StatusInternalServerError) // .🤔
			return c.SendString(errMarshal.Error())
		}

		return c.Send(jsonStr)

	})

	go func() {
		httpPort := flags.httpPort

		if flags.crt != "" {
			// certs/capsule.local.crt
			// certs/capsule.local.key
			log.Println("💊 Capsule", version, "http server is listening on:", httpPort, "🔐🌍")
			app.ListenTLS(":"+httpPort, flags.crt, flags.key)

		} else {
			log.Println("💊 Capsule", version, "http server is listening on:", httpPort, "🌍")
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
