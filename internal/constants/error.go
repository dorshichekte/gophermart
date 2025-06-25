package constants

import (
	customerror "gophermarket/internal/error"
)

var (
	ErrFailedGettingUserID = customerror.New(FailedGettingUserID)
	ErrInvalidOrderNumber  = customerror.New(InvalidOrderNumber)
)
