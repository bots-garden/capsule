package capsulehttp

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

//TODO refactor
