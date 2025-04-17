package app

import "github.com/go-playground/validator/v10"

func NewValidatorOptions() []validator.Option {
	return []validator.Option{}
}

func NewValidator(options []validator.Option) *validator.Validate {
	return validator.New(options...)
}
