package capsulehttp

import (
	"encoding/json"
	"github.com/bots-garden/capsule/capsulelauncher/commons"
	"github.com/labstack/echo/v4"
)

// GetJsonStringFromPayloadRequest :
func GetJsonStringFromPayloadRequest(c echo.Context) (string, error) {
	jsonMap := make(map[string]interface{})
	if err := c.Bind(&jsonMap); err != nil {
		return "", err
	}
	// Convert map to json string
	jsonBytes, err := json.Marshal(jsonMap)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// GetHeadersStringFromHeadersRequest :
func GetHeadersStringFromHeadersRequest(c echo.Context) string {
	var headersMap = make(map[string]string)
	for key, values := range c.Request().Header {
		headersMap[key] = values[0]
	}
	headersSlice := commons.CreateSliceFromMap(headersMap)
	headersParameter := commons.CreateStringFromSlice(headersSlice, "|")

	return headersParameter
}
