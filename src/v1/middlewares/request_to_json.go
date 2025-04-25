package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindJSONData[T any](c *gin.Context) (T, bool) {
	var data T
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return *new(T), false // Return the zero value and false
	}
	return data, true
}

func RequestToJSON[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {

		data, ok := BindJSONData[T](c)
		if !ok {
			return
		}
		c.Set("body", data)

		c.Next()
	}
}
