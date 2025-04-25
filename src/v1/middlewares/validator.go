package middlewares

import (
	"go-simple-rest/src/v1/translator"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func isBodyValid[T any](c *gin.Context) (T, bool) {
	payload, ok := c.MustGet("body").(T)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return *new(T), false
	}
	return payload, true
}

func validateRequest[T any](c *gin.Context) map[string]interface{} {
	payload, ok := isBodyValid[T](c)
	if !ok {
		return nil
	}
	// return payload
	validate := validator.New()
	err := validate.Struct(&payload)
	errs := translator.Translate(validate, err)
	return errs
}

func Validator[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		errs := validateRequest[T](c)
		if len(errs) > 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": errs})
			return
		}
		c.Next()
	}
}
