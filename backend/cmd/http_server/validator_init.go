package main

import (
	"log"
	"unicode"

	"github.com/go-playground/validator"
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

func validatorInit() *validator.Validate {
	v := validator.New()
	err := v.RegisterValidation("uppercase", hasUppercase)
	if err != nil {
		log.Fatal("Failed to setup validator")
	}
	err = v.RegisterValidation("lowercase", hasLowercase)
	if err != nil {
		log.Fatal("Failed to setup validator")
	}
	err = v.RegisterValidation("digit", hasDigit)
	if err != nil {
		log.Fatal("Failed to setup validator")
	}
	return v
}
