// Package main
package main

import (
	capsule "github.com/bots-garden/capsule-module-sdk"
)

func main() {
	capsule.SetHandleHTTP(func (param capsule.HTTPRequest) (capsule.HTTPResponse, error) {

		return capsule.HTTPResponse{
			TextBody: "🟢 👋 Hey " + param.Body +" !",
			Headers: `{"Content-Type": "text/plain; charset=utf-8"}`,
			StatusCode: 200,
		}, nil
		
	})
}
