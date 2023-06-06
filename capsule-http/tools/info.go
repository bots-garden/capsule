// Package tools (info part)
package tools

import (
	_ "embed"
)

//go:embed description.txt
var textVersion []byte

// GetVersion returns the current version
func GetVersion() string {
	return string(textVersion)
}
