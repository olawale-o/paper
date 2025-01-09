package errorTransformer

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func TranslateErrorToMap(err error, trans ut.Translator) map[string]interface{} {
	var errs map[string]interface{}
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := e.Translate(trans)
		errs = map[string]interface{}{
			e.Field(): translatedErr,
		}
	}
	return errs
}
func TranslateError(err error, trans ut.Translator) (errs []string) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := e.Translate(trans)
		errs = append(errs, translatedErr)
	}
	return errs
}
