package main

import (
	"strconv"
	"time"

	"github.com/bots-garden/capsule/capsule-http/handlers"
)

// ShouldStopAfterDelay -> stop the HTTP server after a given delay
func ShouldStopAfterDelay(flags CapsuleFlags) bool {
		// Set a value for the last call
		if flags.stopAfter == "" {
			return false
		}
		duration, _ := strconv.ParseFloat(flags.stopAfter, 64)
		handlers.SetLastCall(time.Now())
		for {
			time.Sleep(1 * time.Second)
			if time.Since(handlers.GetLastCall()).Seconds() >= duration {
				return true
			}
		}
}

