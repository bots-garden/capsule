package capsulehttp

import (
	"encoding/json"
	"github.com/bots-garden/capsule/commons"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func IsBodyString(str string) bool {
	return strings.HasPrefix(str, "[BODY]")
}

func GetBodyString(str string) string {
	return strings.Split(str, "[BODY]")[1]
}

func IsJsonArray(str string) bool {
	return strings.HasPrefix(str, "[") && strings.HasSuffix(str, "]")
}

func GetBodyAndHeaders(bytes []byte, c *gin.Context) (body string, headers map[string]string) {
	response := strings.Split(string(bytes), "[HEADERS]")
	body = response[0]
	respHeadersStr := response[1]

	headers = GetHeadersMapFromString(respHeadersStr)
	//add headers to echo context response
	for key, value := range headers {
		c.Writer.Header().Add(key, value)
		//c.Response().Header().Add(key, value)
	}
	return body, headers
}

func SendErrorMessage(str string, headers map[string]string, c *gin.Context) {
	var returnValue string
	errorMessage, errorCode := commons.GetErrorStringInfo(str)
	if errorCode == 0 {
		returnValue = errorMessage
	} else {
		returnValue = errorMessage + " (" + strconv.Itoa(errorCode) + ")"
	}
	// check content type
	if IsJsonContentType(headers) {
		jsonMap := make(map[string]interface{})
		jsonMap["error"] = returnValue
		jsonMap["from"] = "host"
		c.JSON(500, jsonMap)
	} else {
		c.String(500, returnValue)
	}

}

func SendBodyMessage(str string, headers map[string]string, c *gin.Context) {
	// check content type
	body := GetBodyString(str)

	switch contentType := GetContentType(headers); contentType {
	case "text/plain":
		c.String(http.StatusOK, body)
	case "text/html":
		//c.HTML(http.StatusOK, body, nil)
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(body))

	case "application/json":
		// TODO: to be refactored (see the POST version)
		if IsJsonContentType(headers) {
			// an arbitrary json string

			jsonString := GetBodyString(str)

			// check if it's an array

			if IsJsonArray(jsonString) {
				var jsonMap map[string]interface{}
				var jsonMapArray []map[string]interface{}
				err := json.Unmarshal([]byte(jsonString), &jsonMapArray)

				if err != nil {
					//fmt.Println(err.Error())
					jsonMap = make(map[string]interface{})
					jsonMap["error"] = "JSON string bad format"
					jsonMap["from"] = "host"
					c.JSON(500, jsonMap)
				} else {
					//return c.JSON(http.StatusOK, jsonString)
					c.JSON(http.StatusOK, jsonMapArray)
				}

			} else {
				var jsonMap map[string]interface{}

				err := json.Unmarshal([]byte(jsonString), &jsonMap)
				if err != nil {
					//fmt.Println(err.Error())
					jsonMap = make(map[string]interface{})
					jsonMap["error"] = "JSON string bad format"
					jsonMap["from"] = "host"
					c.JSON(500, jsonMap)
				} else {
					//return c.JSON(http.StatusOK, jsonString)
					c.JSON(http.StatusOK, jsonMap)
				}
			}
		} else {
			c.String(http.StatusOK, GetBodyString(str))
		}

	default:
		message := "🔴 something bad is happening..."
		c.String(500, message)
	}
}

func SendJsonMessage(str string, headers map[string]string, c *gin.Context) {

	// check content type
	if IsJsonContentType(headers) {
		// an arbitrary json string
		jsonString := GetBodyString(str)
		// check if it's an array

		if IsJsonArray(jsonString) {
			var jsonMap map[string]interface{}
			var jsonMapArray []map[string]interface{}
			err := json.Unmarshal([]byte(jsonString), &jsonMapArray)

			if err != nil {
				//fmt.Println(err.Error())
				jsonMap = make(map[string]interface{})
				jsonMap["error"] = "JSON string bad format"
				jsonMap["from"] = "host"
				c.JSON(500, jsonMap)
			} else {
				c.JSON(http.StatusOK, jsonMapArray)
			}

		} else {

			var jsonMap map[string]interface{}

			err := json.Unmarshal([]byte(jsonString), &jsonMap)
			if err != nil {
				//fmt.Println(err.Error())
				jsonMap = make(map[string]interface{})
				jsonMap["error"] = "JSON string bad format"
				jsonMap["from"] = "host"
				c.JSON(500, jsonMap)
			} else {
				//return c.JSON(http.StatusOK, jsonString)
				c.JSON(http.StatusOK, jsonMap)
			}
		}

	} else {
		c.String(http.StatusOK, GetBodyString(str))
	}
}
