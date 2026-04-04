package data

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Errors []FieldError `json:"errors"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func FormatValidationErrors(err error, obj any) ErrorResponse {
	var errors []FieldError

	ve, ok := err.(validator.ValidationErrors)
	if !ok {
		return ErrorResponse{Errors: errors}
	}

	for _, fe := range ve {
		errors = append(errors, FieldError{
			Field:   getJSONFieldName(fe, obj),
			Message: getErrorMessage(fe),
		})
	}

	return ErrorResponse{Errors: errors}
}

func getJSONFieldName(fe validator.FieldError, obj any) string {
	t := reflect.TypeOf(obj)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	field, found := t.FieldByName(fe.StructField())
	if !found {
		return strings.ToLower(fe.Field())
	}

	jsonTag := field.Tag.Get("json")
	if jsonTag == "" {
		return strings.ToLower(fe.Field())
	}

	name := strings.Split(jsonTag, ",")[0]
	if name == "" {
		return strings.ToLower(fe.Field())
	}

	return name
}

func getErrorMessage(fe validator.FieldError) string {
	field := strings.ToLower(fe.Field())

	switch fe.Tag() {
	case "required":
		return field + " is required"
	case "email":
		return "invalid email format"
	case "min":
		return field + " must be at least " + fe.Param() + " characters"
	case "max":
		return field + " must be at most " + fe.Param() + " characters"
	case "len":
		return field + " must be exactly " + fe.Param() + " characters"
	default:
		return field + " is invalid"
	}
}
