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
	
	return capsule.HTTPResponse{
		Body: "<h1>👋 Hello World! 🌍</h1>",
		Headers: `{"Content-Type": "text/html; charset=utf-8"}`,
		StatusCode: 200,
	}, nil
}
