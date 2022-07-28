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


func CreateSliceFromString(str string, separator string) []string {
    return strings.Split(str, separator)
}

func CreateMapFromSlice(strSlice []string, separator string) map[string]string {
    strMap := make(map[string]string)
    for _, item := range strSlice {
        res := strings.Split(item, separator)
        strMap[res[0]] = res[1]
    }
    return strMap
}
