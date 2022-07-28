package commons

import (
	"strings"
)

func IsBodyString(str string) bool {
	return strings.HasPrefix(str, "[BODY]")
}

func GetBodyString(str string) string {
	return strings.Split(str, "[BODY]")[1]
}
