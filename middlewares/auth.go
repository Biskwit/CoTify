package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if Authorization header is present
		if c.Request.Header.Get("Authorization") == "" {
			c.AbortWithStatus(401)
			return
		}

		if !strings.Contains(c.Request.Header.Get("Authorization"), "Bearer ") {
			c.AbortWithStatus(401)
			return
		}

		sk := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")[1]

		if sk == "" {
			c.AbortWithStatus(401)
			return
		}

		c.Set("apiKey", sk)

		c.Next()
	}
}
