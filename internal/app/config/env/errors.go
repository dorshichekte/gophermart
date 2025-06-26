package config

import (
	customerror "gophermarket/internal/error"
)

const (
	errEnvMissingVariables customerror.TextError = "Required environment variables not found"
)
