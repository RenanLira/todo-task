package internalerrors

import (
	"errors"

	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	pt_translations "github.com/go-playground/validator/v10/translations/pt_BR"
)

var (
	pt                               = pt_BR.New()
	uni      *ut.UniversalTranslator = ut.New(pt, pt)
	Validate *validator.Validate     = validator.New()
)

func ValidateStruct(s any) error {
	trans, _ := uni.GetTranslator("pt_BR")

	pt_translations.RegisterDefaultTranslations(Validate, trans)

	err := Validate.Struct(s)

	if err != nil {
		errs := err.(validator.ValidationErrors)

		return errors.New(errs[0].Translate(trans))
	}

	return nil

}
