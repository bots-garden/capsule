// Package main
package main

import (
	"strconv"

	"github.com/bots-garden/capsule-module-sdk"
	"github.com/valyala/fastjson"
)

func main() {
	capsule.SetHandleHTTP(Handle)
}


// OnStart function
//export OnStart
func OnStart() {
	capsule.Print("🚗 OnStart")
}

// OnStop function
//export OnStop
func OnStop() {
	capsule.Print("🚙 OnStop")
}

// OnHealthCheck function
//export OnHealthCheck
func OnHealthCheck() uint64 {
	capsule.Print("⛑️ OnHealthCheck")

	response := capsule.HTTPResponse{
		JSONBody: `{"message": "OK"}`,
		Headers: `{"Content-Type": "application/json; charset=utf-8"}`,
		StatusCode: 200,
	}

	return capsule.Success([]byte(capsule.StringifyHTTPResponse(response)))
}

// OnMetrics function
//export OnMetrics
func OnMetrics() uint64 {
	capsule.Print("📊 OnMetrics")
	response :=  capsule.HTTPResponse{
		TextBody: "xxxxxxxxxxxx",
		Headers: `{"Content-Type": "text/plain; charset=utf-8"}`,
		StatusCode: 200,
	}
	return capsule.Success([]byte(capsule.StringifyHTTPResponse(response)))

}

// Handle function 
func Handle(param capsule.HTTPRequest) (capsule.HTTPResponse, error) {
	
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
		JSONBody: `{"message": "`+message+`", "things":{"emoji":"🐯"}}`,
		Headers: `{"Content-Type": "application/json; charset=utf-8"}`,
		StatusCode: 200,
	}

	return response, nil
}
