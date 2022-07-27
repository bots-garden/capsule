// host functions
package hf

import (
	"strconv"
	"strings"
)


func IsStringError(str string) bool {
	return strings.HasPrefix(str, "[ERR]")
}

func GetStringErrorInfo(str string) (string, int) {
	errorMessage := strings.Split(str, "]:")[1]
	errorCode, _ := strconv.Atoi(strings.Split(strings.Split(str, "]")[1], "[")[1])

	return errorMessage, errorCode
}

func CreateSliceFromMap(strMap map[string]string) []string {
	var strSlice []string
	for field, value := range strMap {
		strSlice = append(strSlice, field+":"+value)
	}
	return strSlice
}

func CreateStringFromSlice(strSlice []string, separator string) string {
	return strings.Join(strSlice[:], separator)
}

// "[ERR][200]:hello world"
// message: e.split("]:")[1]
// code: e.split("]")[1].split("[")[1]
