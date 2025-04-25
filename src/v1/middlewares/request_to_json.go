package middlewares

import (
	"go-simple-rest/src/v1/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func bindJSONData[T any](c *gin.Context) (T, bool) {
	var data T
	if err := c.BindJSON(&data); err != nil {
		utils.AbortResponse(c, utils.Reponse{StatusCode: http.StatusUnprocessableEntity, Success: false, Message: err.Error(), Data: nil})
		return *new(T), false // Return the zero value and false
	}
	return data, true
}

func RequestToJSON[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {

		data, ok := bindJSONData[T](c)
		if !ok {
			return
		}
		c.Set("body", data)

		c.Next()
	}
}
