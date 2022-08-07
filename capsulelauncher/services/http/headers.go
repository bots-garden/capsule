package capsulehttp

import (
	"github.com/bots-garden/capsule/capsulelauncher/commons"
	"strings"
)

func GetHeadersMapFromString(headersStr string) map[string]string {
	return commons.CreateMapFromSlice(commons.CreateSliceFromString(headersStr, "|"), ":")
}

//TODO: add other content types

func IsJsonContentType(headers map[string]string) bool {
	// ! case of key and value
	// TODO: handle the case issue
	if strings.HasPrefix(headers["Content-Type"], "application/json") {
		return true
	} else {
		return false
	}
}

// IsHtmlContentType : Content-Type: text/html; charset=UTF-8
func IsHtmlContentType(headers map[string]string) bool {
	// ! case of key and value
	// TODO: handle the case issue
	if strings.HasPrefix(headers["Content-Type"], "text/html") {
		return true
	} else {
		return false
	}
}

// IsTxtContentType : Content-Type: text/plain; charset=UTF-8
func IsTxtContentType(headers map[string]string) bool {
	// ! case of key and value
	// TODO: handle the case issue
	if strings.HasPrefix(headers["Content-Type"], "text/plain") {
		return true
	} else {
		return false
	}
}

func GetContentType(headers map[string]string) string {
	if IsTxtContentType(headers) {
		return "text/plain"
	}
	if IsJsonContentType(headers) {
		return "application/json"
	}
	if IsHtmlContentType(headers) {
		return "text/html"
	} else {
		return "text/plain"
	}
}
