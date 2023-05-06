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
