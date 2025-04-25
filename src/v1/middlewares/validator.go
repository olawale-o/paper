package middlewares

import (
	"go-simple-rest/src/v1/translator"
	"go-simple-rest/src/v1/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func isBodyValid[T any](c *gin.Context) (T, bool) {
	payload, ok := c.MustGet("body").(T)
	if !ok {
		utils.AbortResponse(c, utils.Reponse{StatusCode: http.StatusBadRequest, Success: false, Message: "Validation error", Data: nil})
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
			utils.AbortResponse(c, utils.Reponse{StatusCode: http.StatusBadRequest, Success: false, Message: errs, Data: nil})
			return
		}
		c.Next()
	}
}
