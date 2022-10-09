package capsulehttp

import (
    "encoding/json"
    "github.com/bots-garden/capsule/commons"
    "github.com/gofiber/fiber/v2"
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

func GetBodyAndHeaders(bytes []byte, c *fiber.Ctx) (body string, headers map[string]string) {
    response := strings.Split(string(bytes), "[HEADERS]")
    body = response[0]
    respHeadersStr := response[1]

    headers = GetHeadersMapFromString(respHeadersStr)

    //add headers to echo context response
    for key, value := range headers {
        c.Append(key, value)
    }

    return body, headers
}

func SendErrorMessage(str string, headers map[string]string, c *fiber.Ctx) error {
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
        c.Status(500)
        return c.JSON(jsonMap)
    } else {
        c.Status(500)
        return c.SendString(returnValue)
    }

}

func SendBodyMessage(str string, headers map[string]string, c *fiber.Ctx) error {
    // check content type
    body := GetBodyString(str)

    switch contentType := GetContentType(headers); contentType {
    case "text/plain":
        c.Status(http.StatusOK)
        return c.SendString(body)
    case "text/html":
        c.Status(http.StatusOK)
        c.Append("Content-Type", headers["Content-Type"])
        return c.Send([]byte(body))

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
                    jsonMap = make(map[string]interface{})
                    jsonMap["error"] = "JSON string bad format"
                    jsonMap["from"] = "host"
                    c.Status(500)
                    return c.JSON(jsonMap)
                } else {
                    c.Status(http.StatusOK)
                    return c.JSON(jsonMapArray)
                }

            } else {
                var jsonMap map[string]interface{}

                err := json.Unmarshal([]byte(jsonString), &jsonMap)
                if err != nil {
                    jsonMap = make(map[string]interface{})
                    jsonMap["error"] = "JSON string bad format"
                    jsonMap["from"] = "host"
                    c.Status(500)
                    return c.JSON(jsonMap)
                } else {
                    c.Status(http.StatusOK)
                    return c.JSON(jsonMap)
                }
            }
        } else {
            c.Status(http.StatusOK)
            return c.SendString(GetBodyString(str))
        }

    default:
        message := "🔴 something bad is happening..."
        c.Status(500)
        return c.SendString(message)
    }
}

func SendJsonMessage(str string, headers map[string]string, c *fiber.Ctx) error {

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
                jsonMap = make(map[string]interface{})
                jsonMap["error"] = "JSON string bad format"
                jsonMap["from"] = "host"
                c.Status(500)
                return c.JSON(jsonMap)
            } else {
                c.Status(http.StatusOK)
                return c.JSON(jsonMapArray)
            }

        } else {

            var jsonMap map[string]interface{}

            err := json.Unmarshal([]byte(jsonString), &jsonMap)
            if err != nil {
                jsonMap = make(map[string]interface{})
                jsonMap["error"] = "JSON string bad format"
                jsonMap["from"] = "host"
                c.Status(500)
                return c.JSON(jsonMap)
            } else {
                c.Status(http.StatusOK)
                return c.JSON(jsonMap)
            }
        }

    } else {
        c.Status(http.StatusOK)
        return c.SendString(GetBodyString(str))
    }
}
