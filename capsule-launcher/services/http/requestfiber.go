package capsulehttp

import (
	"github.com/bots-garden/capsule/commons"
	"github.com/gofiber/fiber/v2"
)

// GetHeadersStringFromHeadersRequest :
func GetHeadersStringFromHeadersRequest(c *fiber.Ctx) string {
	var headersMap = make(map[string]string)
	for key, values := range c.GetReqHeaders() {
		headersMap[key] = values
	}
	headersSlice := commons.CreateSliceFromMap(headersMap)
	headersParameter := commons.CreateStringFromSlice(headersSlice, commons.StrSeparator)

	return headersParameter
}
