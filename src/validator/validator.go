package validator

import (
	playgroundValidator "gopkg.in/go-playground/validator.v9"
	"regexp"
)

type ValidationInterface interface {
	ValidateStruct(s interface{}) error
}

type PlaygroundValidator struct {
	validate *playgroundValidator.Validate
}

func NewPlaygroundValidator() *PlaygroundValidator {
	v := PlaygroundValidator{validate: playgroundValidator.New()}
	_ = v.validate.RegisterValidation("datetime_fmt", datetimeFormatValidation)
	return &v
}

func (v PlaygroundValidator) ValidateStruct(s interface{}) error {
	return v.validate.Struct(s)
}

func datetimeFormatValidation(field playgroundValidator.FieldLevel) bool {
	ymdRegexStr := "^\\d\\d\\d\\d-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])"
	ymdRegex := regexp.MustCompile(ymdRegexStr)
	return ymdRegex.MatchString(field.Field().String())
}
