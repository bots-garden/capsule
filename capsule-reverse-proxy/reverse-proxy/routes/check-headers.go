package reverse_proxy_memory_routes

import "github.com/gin-gonic/gin"

func CheckReverseProxyAdminToken(c *gin.Context, reverseProxyAdminToken string) bool {
	if c.GetHeader("CAPSULE_REVERSE_PROXY_ADMIN_TOKEN") == reverseProxyAdminToken ||
		c.GetHeader("Capsule_Reverse_Proxy_Admin_Token") == reverseProxyAdminToken ||
		c.GetHeader("capsule_reverse_proxy_admin_token") == reverseProxyAdminToken {
		return true
	} else {
		return false
	}
}

// See: https://github.com/gin-gonic/gin/issues/1079
