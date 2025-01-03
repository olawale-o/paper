package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set example variable
		c.Set("userId", "12345")

		// before request

		c.Next()

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)

	}
}
