package middlewares

import (
	"github.com/CELS0/hellogo/services"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer_shelma = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(401)
		}

		token := header[len(Bearer_shelma):]

		if !services.NewJWTService().ValidadeToken(token) {
			c.AbortWithStatus(401)
		}

	}
}
