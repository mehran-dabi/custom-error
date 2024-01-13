package customerror

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var errorMessages = map[string]func(string) string{
	"required": func(param string) string {
		return "is required"
	},
	"min": func(param string) string {
		return fmt.Sprintf("must be greater than %s", param)
	},
}

func ErrorValidator(err error, input any) map[string]string {
	errorMap := make(map[string]string)
	var valErrors validator.ValidationErrors
	ok := errors.As(err, &valErrors)
	if !ok {
		return errorMap
	}
	for _, fieldError := range valErrors {
		fieldName := fieldError.Field()
		field, found := reflect.TypeOf(input).FieldByName(fieldName)
		if !found {
			errorMap[fieldName] = fmt.Sprintf("Unknown field: %s", fieldName)
			continue
		}
		jsonName := field.Tag.Get("json")
		if jsonName == "" {
			jsonName = strings.ToLower(fieldName)
		}

		var message string
		tag := fieldError.Tag()
		if errorMessageFunc, ok := errorMessages[tag]; ok {
			message = errorMessageFunc(fieldError.Param())
		} else {
			message = fmt.Sprintf("validation failed on '%s' condition", fieldError.Tag())
		}

		errorMap[jsonName] = message
	}

	return errorMap
}
