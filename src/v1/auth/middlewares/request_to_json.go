package middlewares

import (
	"go-simple-rest/src/v1/auth/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func bindData(c *gin.Context, data any) bool {
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return false
	}
	return true
}

func RequestToJSON(requestType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch requestType {
		case "login":
			var data model.LoginAuth
			if !bindData(c, &data) {
				return
			}
			c.Set("body", data)
		case "register":
			var data model.RegisterAuth
			if !bindData(c, &data) {
				return
			}
			c.Set("body", data)
		}
		c.Next()
	}
}
