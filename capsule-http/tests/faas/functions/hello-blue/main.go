// Package main
package main

import (
	capsule "github.com/bots-garden/capsule-module-sdk"
)

func main() {
	capsule.SetHandleHTTP(func (param capsule.HTTPRequest) (capsule.HTTPResponse, error) {

		msg := capsule.GetEnv("MESSAGE")
		grt := capsule.GetEnv("GREETING")

		return capsule.HTTPResponse{
			TextBody: "🔵 👋 Hey " + param.Body +" ! ["+msg+" "+grt+"]",
			Headers: `{"Content-Type": "text/plain; charset=utf-8"}`,
			StatusCode: 200,
		}, nil
		
	})
}
