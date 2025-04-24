package errorTransformer

import (
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func TranslateErrorToMap(err error, trans ut.Translator) map[string]interface{} {
	errs := make(map[string]interface{})
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	fmt.Println("validatorErrs", validatorErrs)
	for _, e := range validatorErrs {
		fmt.Printf("Field: %s, Tag: %s, StructField: %s, Namespace: %s\n", e.Field(), e.ActualTag(), e.StructField(), e.Namespace())
		translatedErr := e.Translate(trans)
		errs[e.Field()] = translatedErr
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
