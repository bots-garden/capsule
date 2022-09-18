package commons

import (
	"strconv"
	"strings"
)

var errExit string
var errExitCode int

func GetExitError() string {
	return errExit
}
func SetExitError(errMsg string) {
	errExit = errMsg
}

func GetExitCode() int {
	return errExitCode
}

func SetExitCode(code int) {
	//fmt.Println("📝", code)
	errExitCode = code
}

// CreateStringError :
// to not display an error code at the end of the error message, code == 0
func CreateStringError(message string, code int) string {
	return "[ERR][" + strconv.Itoa(code) + "]:" + message
	// "[ERR][200]:hello world"
	// message: e.split("]:")[1]
	// code: e.split("]")[1].split("[")[1]
}

func IsErrorString(str string) bool {
	return strings.HasPrefix(str, "[ERR]")
}

func GetErrorStringInfo(str string) (string, int) {
	errorMessage := strings.Split(str, "]:")[1]
	errorCode, _ := strconv.Atoi(strings.Split(strings.Split(str, "]")[1], "[")[1])
	return errorMessage, errorCode
}
