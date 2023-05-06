// Package main
package main

import (
	"strconv"

	capsule "github.com/bots-garden/capsule-module-sdk"
)

func main() {
	capsule.SetHandle(Handle)
}

// Handle function
func Handle(params []byte) ([]byte, error) {

	respJSON, err := capsule.HTTP(capsule.HTTPRequest{
		URI:     "http://localhost:3000/json",
		Method:  "GET",
		Headers: `{"Content-Type": "application/json; charset=utf-8"}`,
	})

	if err != nil {
		capsule.Print(err.Error())
	}

	capsule.Print("📝 JSONBody: " + respJSON.JSONBody)
	capsule.Print("📝 TextBody: " + respJSON.TextBody)
	capsule.Print("🅿️ StatusCode: " + strconv.Itoa(respJSON.StatusCode))

	//capsule.Print("💼 Headers: " + resp.Headers)
	respNoJSON, err := capsule.HTTP(capsule.HTTPRequest{
		URI:     "http://localhost:3000/html",
		Method:  "GET",
		Headers: `{"Content-Type": "text/html; charset=utf-8"}`,
	})
	capsule.Print("📝 JSONBody: " + respNoJSON.JSONBody)
	capsule.Print("📝 TextBody: " + respNoJSON.TextBody)
	capsule.Print("🅿️ StatusCode: " + strconv.Itoa(respNoJSON.StatusCode))
	
	respFromPost, err := capsule.HTTP(capsule.HTTPRequest{
		JSONBody:    `{"msg":"hello"}`,
		URI:     "http://localhost:3000",
		Method:  "POST",
		Headers: `{"Content-Type": "application/json; charset=utf-8"}`,
	})
	capsule.Print("📝 JSONBody: " + respFromPost.JSONBody)
	capsule.Print("📝 TextBody: " + respFromPost.TextBody)
	capsule.Print("🅿️ StatusCode: " + strconv.Itoa(respFromPost.StatusCode))

	return []byte("👋 Hello "), nil

}
