package capsulehttp

import (
	"github.com/gofiber/fiber/v2"
)

func GetReloadTokenFromHeadersRequest(c *fiber.Ctx) string {
	/*
	   Fiber canonicalizes the header name when adding values to the header map.
	   ex: `CAPSULE_RELOAD_TOKEN` becomes `Capsule_reload_token`
	*/
	return c.GetReqHeaders()["Capsule_reload_token"]
}
