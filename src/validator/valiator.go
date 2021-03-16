package validator

import playgroundValidator "gopkg.in/go-playground/validator.v9"

type ValidationInterface interface {
	ValidateStruct(s interface{}) error
}

type PlaygroundValidator struct {
	validate *playgroundValidator.Validate
}

func NewPlaygroundValidator() *PlaygroundValidator {
	return &PlaygroundValidator{validate: playgroundValidator.New()}
}

func (v PlaygroundValidator) ValidateStruct(s interface{}) error {
	return v.validate.Struct(s)
}
