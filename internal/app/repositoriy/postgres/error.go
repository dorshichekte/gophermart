package repository_postgres

import (
	customerror "gophermarket/internal/error"
)

const (
	errDBFailedInitMigrations  customerror.TextError = "Failed initialize migrations"
	errDBFailedApplyMigrations customerror.TextError = "Failed apply migrations"
)
