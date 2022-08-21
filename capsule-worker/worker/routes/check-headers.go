package routes

import "github.com/gin-gonic/gin"

func CheckWorkerAdminToken(c *gin.Context, workerAdminToken string) bool {
	if c.GetHeader("CAPSULE_WORKER_ADMIN_TOKEN") == workerAdminToken ||
		c.GetHeader("Capsule_Worker_Admin_Token") == workerAdminToken ||
		c.GetHeader("capsule_worker_admin_token") == workerAdminToken {
		return true
	} else {
		return false
	}
}

// See: https://github.com/gin-gonic/gin/issues/1079
