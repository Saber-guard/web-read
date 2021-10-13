package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header["Origin"][0]
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Cookie, Accept")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, OPTIONS")
		c.Header("Access-Control-Allow-Credentials", "true")
		method := c.Request.Method
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
