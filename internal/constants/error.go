package constants

import (
	customerror "gophermarket/internal/error"
)

const (
	ErrLoggerFailedInitialization customerror.TextError = "Failed initialization logger"
	ErrFailedGettingUserID                              = "Failed getting user ID"

	ErrPathUnknownFolderPath customerror.TextError = "Unable to get directory path"
	ErrDecompressRequestBody customerror.TextError = "Error parsing request body"
)
