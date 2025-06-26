package auth

import (
	customerror "gophermarket/internal/error"
)

var (
	errExpiredToken            = customerror.New(expiredToken)
	errInvalidToken            = customerror.New(invalidToken)
	errInitializationToken     = customerror.New(initializationToken)
	errEmptyUserID             = customerror.New(emptyUserID)
	errUnexpectedSigningMethod = customerror.New(unexpectedSigningMethod)
)
