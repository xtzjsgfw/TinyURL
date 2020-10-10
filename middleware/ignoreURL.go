package middleware

import (
	"github.com/gin-gonic/gin"
)

func IgnoreURL() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tinyUrl string
		tinyUrl = c.Param("tinyUrl")
		if tinyUrl == "/favicon.ico" {
			c.Abort()
		}
	}
}
