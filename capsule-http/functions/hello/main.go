// Package main
package main

import (
	"strconv"
	"strings"

	"github.com/bots-garden/capsule-module-sdk"
	"github.com/valyala/fastjson"
)

var counter int

func main() {
	
	counter = 0

	capsule.SetHandleHTTP(func(param capsule.HTTPRequest) (capsule.HTTPResponse, error) {
		
		counter++
		capsule.Print("🟢 Counter: " + strconv.Itoa(counter))

		capsule.Print("📝: " + param.Body)
		capsule.Print("🔠: " + param.Method)
		capsule.Print("🌍: " + param.URI)
		capsule.Print("👒: " + param.Headers)
	
		var p fastjson.Parser
		v, err := p.Parse(param.Body)
		if err != nil {
			capsule.Log(err.Error())
		}
		message := string(v.GetStringBytes("name")) + " " + strconv.Itoa(v.GetInt("age"))
		capsule.Log(message)
	
		response := capsule.HTTPResponse{
			JSONBody:   `{"message": "` + message + `", "things":{"emoji":"🐯"}}`,
			Headers:    `{"Content-Type": "application/json; charset=utf-8"}`,
			StatusCode: 200,
		}
	
		return response, nil
	})
}

// OnStart function
//export OnStart
func OnStart() {
	capsule.Print("🚗 OnStart")
	counter = 42
}

// OnStop function
//export OnStop
func OnStop() {
	capsule.Print("🚙 OnStop")
	capsule.Print("🟢 Counter: " + strconv.Itoa(counter))

}

// OnHealthCheck function
//export OnHealthCheck
func OnHealthCheck() uint64 {
	capsule.Print("⛑️ OnHealthCheck")

	response := capsule.HTTPResponse{
		JSONBody:   `{"message": "OK"}`,
		Headers:    `{"Content-Type": "application/json; charset=utf-8"}`,
		StatusCode: 200,
	}

	return capsule.Success([]byte(capsule.StringifyHTTPResponse(response)))
}

// OnMetrics function
//export OnMetrics
func OnMetrics() uint64 {
	capsule.Print("📊 OnMetrics")

	// Generate OpenText Prometheus metric
	counterMetrics := []string{
		"# HELP call counter",
		"# TYPE call_counter counter",
		"call_counter " + strconv.Itoa(counter)}

	response := capsule.HTTPResponse{
		TextBody:   strings.Join(counterMetrics, "\n"),
		Headers:    `{"Content-Type": "text/plain; charset=utf-8"}`,
		StatusCode: 200,
	}
	return capsule.Success([]byte(capsule.StringifyHTTPResponse(response)))

}
