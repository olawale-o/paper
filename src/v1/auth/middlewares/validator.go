package middlewares

import (
	"go-simple-rest/src/v1/auth/model"
	"go-simple-rest/src/v1/translator"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validator() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload, _ := c.MustGet("body").(model.LoginAuth)
		validate := validator.New()
		err := validate.Struct(payload)
		errs := translator.Translate(validate, err)

		if len(errs) > 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": errs})
			return
		}

		c.Set("body", payload)
		c.Next()

	}
}
