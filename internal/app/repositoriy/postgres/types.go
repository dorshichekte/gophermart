package repositorypostgres

import (
	"database/sql"

	"go.uber.org/zap"

	"gophermarket/internal/app/config/env"
)

type Postgres struct {
	l      *zap.Logger
	config envconfig.Config
	DB     *sql.DB
}
