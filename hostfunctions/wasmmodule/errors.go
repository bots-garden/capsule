// host functions
package hf

import (
	"strconv"
	"strings"
)

// to not display an error code at the end of the error message, code == 0
func CreateStringError(message string, code int) string {
	return "[ERR][" + strconv.Itoa(code) + "]:" + message
	// "[ERR][200]:hello world"
	// message: e.split("]:")[1]
	// code: e.split("]")[1].split("[")[1]
}

func IsStringError(str string) bool {
	return strings.HasPrefix(str, "[ERR]")
}

func GetStringErrorInfo(str string) (string, int) {
	errorMessage := strings.Split(str, "]:")[1]
	errorCode, _ := strconv.Atoi(strings.Split(strings.Split(str, "]")[1], "[")[1])

	return errorMessage, errorCode
}
