package flatjson

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"
)

func tokenize(value string) string {

	if strings.HasPrefix(value, "\"") {
		return "STRING"
	} else {
		if strings.HasPrefix(value, "true") {
			return "BOOL"
		} else {
			return "NUMBER" // I should test if the string contains a number
		}
	}
}

func trimFirstRune(str string) string {
	_, i := utf8.DecodeRuneInString(str)
	return str[i:]
}

func trimLastRune(str string) string {
	for len(str) > 0 {
		_, size := utf8.DecodeLastRuneInString(str)
		return str[:len(str)-size]
	}
	return str
}

func trimFirstAndLastRune(str string) string {
	return trimFirstRune(trimLastRune(str))
}

//TODO: isJsonStr
//TODO: StrToSlice
//TODO: isJsonArrayStr
//TODO: isFieldExist
//TODO: flatten

//QUESTION: should I return an error too?

// StrToMap : get a flat json string (no nested obj) and return a map
func StrToMap(flatJsonStr string) map[string]interface{} {

	firstStep := strings.Split(flatJsonStr, "{")
	var secondStep []string
	for _, item := range firstStep {
		secondStep = append(secondStep, strings.Split(item, "}")[0])
	}
	dataStr := secondStep[1]
	//hf.Log("=> " + dataStr)
	thirdStep := strings.Split(dataStr, ",")

	jsonMap := make(map[string]interface{})
	for _, item := range thirdStep {
		keyVal := strings.Split(item, ":")
		// trim quotes ""
		key := trimFirstAndLastRune(keyVal[0])
		value := keyVal[1]

		switch what := tokenize(value); what {
		case "STRING":
			// trim quotes ""
			jsonMap[key] = trimFirstAndLastRune(value)

		case "BOOL":
			if value == "true" {
				jsonMap[key] = true
			} else {
				jsonMap[key] = false
			}

		case "NUMBER":
			if strings.Contains(value, ".") {
				floatVal, err := strconv.ParseFloat(value, 8)
				if err != nil {
					jsonMap[key] = err
				} else {
					jsonMap[key] = floatVal
				}
			} else {
				intVal, err := strconv.Atoi(value)
				if err != nil {
					jsonMap[key] = err
				} else {
					jsonMap[key] = intVal
				}
			}

		default:
			jsonMap[key] = errors.New("flatJsonStrToMap: unknown type")
		}

		//hf.Log("-> " + item + " " + key + " = " + keyVal[1])
	}
	return jsonMap
}

// MapToStr get a flat json map and return a json string
func MapToStr(jsonMap map[string]interface{}) string {
	items := []string{}
	for key, value := range jsonMap {

		switch value.(type) {
		case float64, float32, uint, int:
			item := `"` + key + `":` + value.(string)
			items = append(items, item)
		case string:
			item := `"` + key + `":"` + value.(string) + `"`
			items = append(items, item)

		case bool:
			item := `"` + key + `":` + value.(string)
			items = append(items, item)
		default:
			item := `"` + key + `":undefined`
			items = append(items, item)
		}
	}
	return "{" + strings.Join(items, ",") + "}"
}

/* the package is too heavy
case types.Nil:
	item := `"` + key + `":null`
	items = append(items, item)
*/
