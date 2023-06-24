package handlers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/bots-garden/capsule-host-sdk"
	"github.com/bots-garden/capsule-host-sdk/models"
	"github.com/gofiber/fiber/v2"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

var wasmFile []byte
var ctx context.Context
var runtime wazero.Runtime

// We want to know when the last call was made (in seconds)
// for example to stop the process if it is too old
var lastCallAt time.Time

// GetLastCall returns the time of the last call.
//
// This function does not take any parameters.
// It returns a time.Time object.
func GetLastCall() time.Time {
	return lastCallAt
}

// SetLastCall sets the last call time to the given time.
//
// t time.Time
func SetLastCall(t time.Time) {
	//fmt.Println("🔵 SetLastCall", t)
	lastCallAt = t
}

// StoreWasmFile stores the given byte slice of a loaded WebAssembly file.
//
// loadedWasmFile is a byte slice containing the contents of the loaded WebAssembly file.
func StoreWasmFile(loadedWasmFile []byte) {
	wasmFile = loadedWasmFile
}

// StoreContext stores the given context for later use.
//
// capsuleContext: the context to be stored.
func StoreContext(capsuleContext context.Context) {
	ctx = capsuleContext
}

// StoreRuntime stores the given capsule runtime.
//
// capsuleRuntime: a wazero.Runtime object representing the runtime to be stored.
func StoreRuntime(capsuleRuntime wazero.Runtime) {
	runtime = capsuleRuntime
}

// store all your plugins in a normal Go hash map, protected by a Mutex
var m sync.Mutex
var modules = make(map[string](api.Module))

// GetModule from modules map
func GetModule(ctx context.Context, wasmFile []byte) (api.Module, error) {
	if module, ok := modules["code"]; ok {
		return module, nil
	}
	module, err := runtime.Instantiate(ctx, wasmFile)
	if err != nil {
		return nil, err
	}
	modules["code"] = module
	return module, nil
}

// CallWasmFunction is a function that handles the execution of a WebAssembly function.
//
// c is a pointer to a Fiber context.
// error is the error returned by the function.
func CallWasmFunction(c *fiber.Ctx) error {
	// register the last call
	SetLastCall(time.Now())

	 // get a lock on the mutex, ensuring no other goroutine has access to the plugin while its executing
	m.Lock()
	// don't forget to release the lock on the Mutex, sometimes its best to `defer m.Unlock()` right after yout get the lock
	defer m.Unlock()
	mod, err := GetModule(ctx, wasmFile)

	//mod, err := runtime.Instantiate(ctx, wasmFile)
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

	//! if TextBody contains "\n" or quotes there is an error (fix something in capsule-module-sdk/hande.http.go => callHandleHTTP)
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
		decodedStrAsByteSlice, _ := base64.StdEncoding.DecodeString(string(response.TextBody))

		// send text body
		//return c.SendString(response.TextBody)
		return c.SendString(string(decodedStrAsByteSlice))
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

func callExportedFunction(exportedFunc string, c *fiber.Ctx) error {
	 // get a lock on the mutex, ensuring no other goroutine has access to the plugin while its executing
	 m.Lock()
	 // don't forget to release the lock on the Mutex, sometimes its best to `defer m.Unlock()` right after yout get the lock
	 defer m.Unlock()
	 mod, err := GetModule(ctx, wasmFile)

	/*
	mod, err := runtime.Instantiate(ctx, wasmFile)
	if err != nil {
		log.Println("❌ Error with the module instance", err)
		c.Status(http.StatusInternalServerError) // .🤔
		return c.SendString(err.Error())
	}
	*/

	// Get the reference to the WebAssembly function: "OnHealthCheck" or "OnMetrics"
	//! this function is exported by the Capsule plugin
	wasmFunction := mod.ExportedFunction(exportedFunc)
	if wasmFunction == nil {
		log.Println("❌ The exported function " + exportedFunc + " does not exist")
		c.Status(http.StatusInternalServerError) // .🤔
		return c.SendString("❌ The exported function" + exportedFunc + "does not exist")
	}

	// Now, we can call the exported function
	// the result type is []uint64
	result, err := wasmFunction.Call(ctx)
	if err != nil {
		log.Println("❌ Error when calling "+exportedFunc, err)
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
		decodedStrAsByteSlice, _ := base64.StdEncoding.DecodeString(string(response.TextBody))

		return c.SendString(string(decodedStrAsByteSlice))
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

// CallWasmFunctionHealthCheck is a function that handles the execution of a WebAssembly function.
// Route: app.All("/health", handlers.CallWasmFunctionHealthCheck)
func CallWasmFunctionHealthCheck(c *fiber.Ctx) error {
	//TODO: Protect the route if Token ?
	return callExportedFunction("OnHealthCheck", c)

}

// CallWasmFunctionMetrics is a function that handles the execution of a WebAssembly function.
// Route: app.All("/metrics", handlers.CallWasmFunctionMetrics)
func CallWasmFunctionMetrics(c *fiber.Ctx) error {
	//TODO: Protect the route if Token ?
	return callExportedFunction("OnMetrics", c)

}
