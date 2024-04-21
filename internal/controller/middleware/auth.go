package middleware

import (
	"article-service/internal/controller/response"
	"article-service/lib/auth"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	Admin         = "admin"
	Moderator     = "moderator"
	Student       = "student"
	UsernameClaim = "username"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerHeader := c.Request.Header.Get("Authorization")
		if bearerHeader == "" {
			apiErr := response.CollectError(response.ErrorNoToken)
			c.AbortWithStatusJSON(apiErr.StatusCode, apiErr)
			return
		}

		token := strings.Split(bearerHeader, " ")[1]
		claims, err := auth.GetPayloadAndValidate(token, jwtSecret)
		if err != nil {
			apiErr := response.CollectError(response.ErrorInvalidToken)
			c.AbortWithStatusJSON(apiErr.StatusCode, apiErr)
			return
		}

		if claims.Role == Student {
			apiErr := response.CollectError(response.ErrorAccessDenied)
			c.AbortWithStatusJSON(apiErr.StatusCode, apiErr.Message)
			return
		}
		c.Set(UsernameClaim, claims.Username)

		c.Next()
	}
}
