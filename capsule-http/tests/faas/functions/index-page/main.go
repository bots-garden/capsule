// Package main => serving an html resource
package main

import (
	_ "embed"
	capsule "github.com/bots-garden/capsule-module-sdk"
)

var (
	//go:embed index.html
	html []byte
)

func main() {

	capsule.SetHandleHTTP(func (param capsule.HTTPRequest) (capsule.HTTPResponse, error) {
		
		return capsule.HTTPResponse{
			TextBody: string(html),
			Headers: `{
				"Content-Type": "text/html; charset=utf-8",
				"Cache-Control": "no-cache",
				"X-Powered-By": "capsule-module-sdk"
			}`,
			StatusCode: 200,
		}, nil
	})
}
//TODO: make something easier for the headers (see old version of Capsule)