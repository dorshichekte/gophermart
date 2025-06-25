package userusecase

import (
	customerror "gophermarket/internal/error"
)

var (
	ErrLoginAlreadyTaken = customerror.New(loginAlreadyTaken)
	ErrUserNotFound      = customerror.New(userNotFound)
)
