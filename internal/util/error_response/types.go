package util

import v "gophermarket/internal/libs/validator"

type ResponseTypeError interface {
	~string | ~[]string | ~[]v.ValidationError | ~
}

type WrapperError[T ResponseTypeError] struct {
	CustomError T `json:"error"`
}
