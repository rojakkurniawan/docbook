package middleware

import (
	"docbook/utils"
	errormodel "docbook/utils/errorModel"
	"docbook/utils/response"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response.BuildErrorResponse(c, errormodel.ErrTokenRequired)
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.BuildErrorResponse(c, errormodel.ErrInvalidAuthorizationHeader)
			c.Abort()
			return
		}

		token := parts[1]
		if token == "" {
			response.BuildErrorResponse(c, errormodel.ErrTokenRequired)
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(token)
		if err != nil {
			fmt.Println("err", err)
			if errors.Is(err, jwt.ErrTokenExpired) {
				response.BuildErrorResponse(c, errormodel.ErrTokenExpired)
			} else if errors.Is(err, jwt.ErrSignatureInvalid) {
				response.BuildErrorResponse(c, errormodel.ErrInvalidToken)
			} else {
				response.BuildErrorResponse(c, errormodel.ErrInvalidToken)
			}
			c.Abort()
			return
		}

		if claims == nil || claims.UserID == 0 {
			response.BuildErrorResponse(c, errormodel.ErrInvalidToken)
			c.Abort()
			return
		}

		c.Set(utils.DOCBOOK_USERID, claims.UserID)
		c.Set(utils.DOCBOOK_ROLE, claims.Role)
		c.Next()
	}
}
