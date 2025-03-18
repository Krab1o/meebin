package api

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

const (
	minMessage       = "%s should be at least %s characters"
	maxMessage       = "%s should be at most %s characters"
	emailMessage     = "Invalid email format"
	requiredMessage  = "%s field is required"
	uppercaseMessage = "Password must contain at least one uppercase letter"
	lowercaseMessage = "Password must contain at least one lowercase letter"
	digitMessage     = "Password must contain at least one digit"
	defaultMessage   = "Invalid value"
)

// TODO: custom map can implement error interface
func ParseValidationErrors(err error) any {
	errorMap := make(map[string]string)
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return "Wrong JSON syntax"
	}
	for _, err := range errs {
		field := err.Field()
		tag := err.Tag()
		param := err.Param()

		switch tag {
		case "min":
			errorMap[field] = fmt.Sprintf(minMessage, field, param)
		case "max":
			errorMap[field] = fmt.Sprintf(maxMessage, field, param)
		case "required":
			errorMap[field] = fmt.Sprintf(requiredMessage, field)
		case "email":
			errorMap[field] = fmt.Sprintf(emailMessage)
		case "uppercase":
			errorMap[field] = fmt.Sprintf(uppercaseMessage)
		case "lowercase":
			errorMap[field] = fmt.Sprintf(lowercaseMessage)
		case "digit":
			errorMap[field] = fmt.Sprintf(digitMessage)
		default:
			errorMap[field] = fmt.Sprintf(defaultMessage)
		}
	}
	return errorMap
}

// func ValidateStruct(validate validator.Validate, data interface{}) map[string]string {
// 	errs := validate.Struct(data)
// 	if errs == nil {
// 		return nil
// 	}

// 	errorMap := make(map[string]string)
// 	for _, err := range errs.(validator.ValidationErrors) {
// 		field := err.Field()
// 		tag := err.Tag()
// 		param := err.Param()

// 		switch tag {
// 		case "min":
// 			errorMap[field] = fmt.Sprintf(minMessage, field, param)
// 		case "max":
// 			errorMap[field] = fmt.Sprintf(maxMessage, field, param)
// 		case "required":
// 			errorMap[field] = fmt.Sprintf(requiredMessage, field)
// 		case "email":
// 			errorMap[field] = fmt.Sprintf(emailMessage)
// 		case "uppercase":
// 			errorMap[field] = fmt.Sprintf(uppercaseMessage)
// 		case "lowercase":
// 			errorMap[field] = fmt.Sprintf(uppercaseMessage)
// 		case "digit":
// 			errorMap[field] = fmt.Sprintf(uppercaseMessage)
// 		default:
// 			errorMap[field] = fmt.Sprintf(defaultMessage)
// 		}
// 	}
// 	return errorMap
// }
