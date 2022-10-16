package capsulehttp

import (
	"fmt"
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

func GetReloadTokenFromHeadersRequest(c *fiber.Ctx) string {
	/*
	   Fiber canonicalizes the header name when adding values to the header map.
	   ex: `CAPSULE_RELOAD_TOKEN` becomes `Capsule_reload_token`
	*/
	fmt.Println(c.GetReqHeaders())
	return c.GetReqHeaders()["Capsule_reload_token"]
}
