package basehandler

import (
	"encoding/json"
	"errors"
	"io"
	"strings"

	customerror "gophermarket/internal/error"
)

func (h *BaseHandler) handleJSONDecodeError(err error) error {
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError

	switch {
	case errors.As(err, &syntaxError) || errors.Is(err, io.ErrUnexpectedEOF):
		return customerror.NewWithData(errRequestBodyContainsBadlyFormedJSON, syntaxError.Offset)

	case errors.As(err, &unmarshalTypeError):
		return customerror.NewWithData(errRequestBodyContainsInvalidValueField, unmarshalTypeError.Field, unmarshalTypeError.Offset)

	case strings.HasPrefix(err.Error(), "json: unknown field "):
		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
		return customerror.NewWithData(errRequestBodyContainsUnknownField, fieldName)

	case errors.Is(err, io.EOF):
		return customerror.New(errRequestBodyMustNotBeEmpty)

	default:
		return err
	}
}
