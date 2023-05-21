package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/bots-garden/capsule-host-sdk"
	"github.com/bots-garden/capsule-host-sdk/models"
	"github.com/gofiber/fiber/v2"
	"github.com/tetratelabs/wazero"
)

var wasmFile []byte
var ctx context.Context
var runtime wazero.Runtime

func StoreWasmFile(loadedWasmFile []byte)  {
	wasmFile = loadedWasmFile
}

func StoreContext(capsuleContext context.Context) {
	ctx = capsuleContext
}

func StoreRuntime(capsuleRuntime wazero.Runtime) {
	runtime = capsuleRuntime
}

// CallWasmFunction is a function that handles the execution of a WebAssembly function.
//
// c is a pointer to a Fiber context.
// error is the error returned by the function.
func CallWasmFunction(c *fiber.Ctx) error {
	mod, err := runtime.Instantiate(ctx, wasmFile)
		if err != nil {
			log.Println("❌ Error with the module instance", err)
			c.Status(http.StatusInternalServerError) // .🤔
			return c.SendString(err.Error())
		}

		// Get the reference to the WebAssembly function: "callHandle"
		//! callHandle is exported by the Capsule plugin
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

		responseFromWasmGuest, err := capsule.Result(responseBuffer)
		if err != nil {
			log.Println("❌ Error when getting the Result", err)
			c.Status(http.StatusInternalServerError) // .🤔
			return c.SendString(err.Error())
		}

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
}
