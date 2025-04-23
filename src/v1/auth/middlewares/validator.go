package middlewares

import (
	"go-simple-rest/src/v1/auth/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validator() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload, ok := c.MustGet("body").(model.LoginAuth)
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
			return
		}
		errs := payload.Validate()
		if len(errs) > 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": errs})
			return
		}

		c.Set("body", payload)
		c.Next()
	}
}
