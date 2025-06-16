package validator

import (
	"errors"
	"fmt"
	customerror "gophermarket/internal/error"
	"strings"

	"github.com/go-playground/validator/v10"
)

func New() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func (v *Validator) ValidateStruct(s any) error {
	return v.validator.Struct(s)
}

func (v *Validator) ParseValidationErrors(err error) ([]ValidationError, error) {
	var listErrors []ValidationError

	fmt.Print(err)
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		for _, e := range errs {
			ce := v.CreateError(e)
			listErrors = append(listErrors, ce)
		}

		return listErrors, nil
	}

	return nil, customerror.New(errFailedParseValidationErrors)
}

func (v *Validator) CreateError(e validator.FieldError) ValidationError {
	field := strings.ToLower(e.Field())

	switch e.Tag() {
	case "required":
		return ValidationError{
			Field:   field,
			Message: fmt.Sprintf("%s is required", field),
		}
	case "min":
		return ValidationError{
			Field:   field,
			Message: fmt.Sprintf("%s must be at least %s characters long.", field, e.Param()),
		}
	case "max":
		return ValidationError{
			Field:   field,
			Message: fmt.Sprintf("%s must be at most %s characters long.", field, e.Param()),
		}
	default:
		return ValidationError{
			Field:   field,
			Message: e.Error(),
		}
	}
}
