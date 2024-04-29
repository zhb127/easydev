package validate

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"golang.org/x/text/language"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var t ut.Translator
var v = validator.New()

func init() {
	en := en.New()
	ut := ut.New(en)

	t, _ = ut.GetTranslator(language.English.String())

	enTranslations.RegisterDefaultTranslations(v, t)
}

func Struct(s interface{}) error {
	if err := v.Struct(s); err != nil {
		return formatValidationError(err)
	}
	return nil
}

func formatValidationError(err error) error {
	vErrs := err.(validator.ValidationErrors)
	vErrTrans := vErrs.Translate(t)

	errMsgs := make([]string, 0)
	for k, v := range vErrTrans {
		errMsgs = append(errMsgs, fmt.Sprintf("%s, %s", k, v))
	}

	return errors.New(strings.Join(errMsgs, "\n"))
}
