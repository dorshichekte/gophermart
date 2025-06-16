package constants

import (
	customerror "gophermarket/internal/error"
)

const (
	ErrFailedGettingUserID   customerror.TextError = "Failed getting user ID"
	ErrPathUnknownFolderPath customerror.TextError = "Unable to get directory path"
)
