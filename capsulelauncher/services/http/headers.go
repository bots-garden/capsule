package commons

import (
	"github.com/bots-garden/capsule/capsulelauncher/commons"
	"strings"
)

func GetHeadersMapFromString(headersStr string) map[string]string {
	return commons.CreateMapFromSlice(commons.CreateSliceFromString(headersStr, "|"), ":")
}

func IsJsonContentType(headers map[string]string) bool {
	// ! case of key and value
	// TODO: handle the case issue
	if strings.HasPrefix(headers["Content-Type"], "application/json") {
		return true
	} else {
		return false
	}
}
