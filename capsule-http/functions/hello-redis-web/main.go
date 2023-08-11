// Package main
package main

import (
	//"errors"

	"errors"

	"github.com/bots-garden/capsule-module-sdk"
)

func main() {

	capsule.SetHandleHTTP(func(param capsule.HTTPRequest) (capsule.HTTPResponse, error) {

		// get values from Redis
		message, errMsg := capsule.RedisGet("message")
		xHandle, errHandle := capsule.RedisGet("x_handle")
		errs := errors.Join(errMsg, errHandle)

		if errs != nil {
			capsule.Log("😡 " + errs.Error())
		}

		response := capsule.HTTPResponse{
			JSONBody:   `{"message": "` + string(message) + `", "xHandle": "` + string(xHandle) + `"}`,
			Headers:    `{"Content-Type": "application/json; charset=utf-8"}`,
			StatusCode: 200,
		}

		return response, errs

	})
}
