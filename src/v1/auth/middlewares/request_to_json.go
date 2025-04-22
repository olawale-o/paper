package middlewares

import (
	"go-simple-rest/src/v1/auth/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequestToJSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.LoginAuth
		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		c.Set("body", user)
		c.Next()
	}
}
