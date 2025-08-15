package middleware

import (
	"docbook/utils"
	errormodel "docbook/utils/errorModel"
	"docbook/utils/response"

	"github.com/gin-gonic/gin"
)

func RBACAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString(utils.DOCBOOK_ROLE)

		if role != "admin" {
			response.BuildErrorResponse(c, errormodel.ErrNotAllowed)
			c.Abort()
			return
		}

		c.Next()
	}
}

func RBACDoctorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString(utils.DOCBOOK_ROLE)

		if role != "doctor" {
			response.BuildErrorResponse(c, errormodel.ErrNotAllowed)
			c.Abort()
			return
		}

		c.Next()
	}
}
