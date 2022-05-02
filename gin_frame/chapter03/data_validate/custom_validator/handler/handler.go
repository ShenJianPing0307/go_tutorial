package handler

import "github.com/go-playground/validator"

var CustomValid validator.Func = func(fl validator.FieldLevel) bool {

	data := fl.Field().Interface().(string)

	if len(data) > 3 {
		return false
	} else {
		return true
	}
}
