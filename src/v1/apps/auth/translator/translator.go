package translator

import (
	errorTransformer "auth/error-transformer"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func Translate(validate *validator.Validate, err error) map[string]interface{} {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)
	errs := errorTransformer.TranslateErrorToMap(err, trans)
	return errs
}
