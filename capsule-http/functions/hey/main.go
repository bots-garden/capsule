// Package main
package main

import (
	capsule "github.com/bots-garden/capsule-module-sdk"
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
	
	return capsule.HTTPResponse{
		TextBody: "👋 Hey " + param.Body +" !",
		Headers: `{"Content-Type": "text/plain; charset=utf-8"}`,
		StatusCode: 200,
	}, nil
}
