package main

import (
	"errors"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func hasUppercase(fl validator.FieldLevel) bool {
	for _, char := range fl.Field().String() {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

// Check if password has at least one lowercase letter
func hasLowercase(fl validator.FieldLevel) bool {
	for _, char := range fl.Field().String() {
		if unicode.IsLower(char) {
			return true
		}
	}
	return false
}

// Check if password has at least one digit
func hasDigit(fl validator.FieldLevel) bool {
	for _, char := range fl.Field().String() {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func validatorInit(s *gin.Engine) (*validator.Validate, error) {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := v.RegisterValidation("uppercase", hasUppercase)
		if err != nil {
			return nil, err
		}
		err = v.RegisterValidation("lowercase", hasLowercase)
		if err != nil {
			return nil, err
		}
		err = v.RegisterValidation("digit", hasDigit)
		if err != nil {
			return nil, err
		}

		return v, nil
	}
	return nil, errors.New("Failed to get gin validator")
}
