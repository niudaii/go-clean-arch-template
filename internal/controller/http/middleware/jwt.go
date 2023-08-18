package middleware

import (
	"go-clean-template/pkg/jwt"
	"go-clean-template/pkg/response"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(jwt.GetConfig().HeaderName)
		if token == "" {
			response.UnAuthWithMessage(jwt.TokenIsEmpty.Error(), c)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			response.UnAuthWithMessage(err.Error(), c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
