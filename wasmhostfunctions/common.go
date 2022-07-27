// host functions
package hf

import (
	"strings"
)

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
