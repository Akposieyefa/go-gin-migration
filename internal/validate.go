package internal

import "github.com/go-playground/validator/v10"

func Validate(request interface{}) error {
	validate := validator.New()
	return validate.Struct(request)
}
