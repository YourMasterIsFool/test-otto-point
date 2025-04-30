package validation

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate   *validator.Validate
	Translator ut.Translator
)

// Helper function to convert camelCase to snake_case
func toSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func InitValidator() {
	eng := en.New()
	uni := ut.New(eng, eng)

	trans, _ := uni.GetTranslator("en")

	Validate = validator.New()

	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	_ = enTranslations.RegisterDefaultTranslations(Validate, trans)

	// custom requred
	Validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is dibtuhkan", false)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		fieldName := toSnakeCase(fe.Field())
		t, _ := ut.T("required", fieldName)
		return t
	})

	Validate.RegisterTranslation("numeric", trans, func(ut ut.Translator) error {
		return ut.Add("numeric", "{0} harus angka", false)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		fieldName := toSnakeCase(fe.Field())
		t, _ := ut.T("numeric", fieldName)
		return t
	})

	Translator = trans
}

func FormatValidationErrors(err error) map[string]string {
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	validationErrors := err.(validator.ValidationErrors)

	for _, e := range validationErrors {
		// Convert field name to snake_case for the error key
		fieldName := toSnakeCase(e.Field())
		errors[fieldName] = e.Translate(Translator)
	}

	return errors
}
